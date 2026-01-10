---
title: "Using WebRTC and Ghostty to get a terminal when GitHub Actions fail"
date: 2025-05-06T07:02:00+00:00
author: "Lawrence Gripper"
tags: ["programming", "github", "actions", "terminal", "debugging"]
categories: ["programming"]
url: /2026/01/10/actions-terminal-on-failure-for-debugging/
draft: true
---

**Spoiler:** I made a free and open-source way to get an interactive web terminal to your GitHub Action when it fails. Try it out here: https://actions-term.gripdev.xyz/

I think we've all been there, **you're build fails in Actions, but the script works fine locally.** You now settle into a slow loop of:

1. Push speculative change
2. See if it worked

It was in the middle of one of these when I started thinking about how to make it better. 

A Terminal would be great, that's obvious, but how to make it happen? How could I make it a free, open to anyone, without costing me lots of money?

Operating a service that forward traffic between a customer and the Actions VM would cost money. 

What about a Peer-to-Peer connection? I'd recently been reading about how [Tailscale](https://tailscale.com/blog/how-tailscale-works), [iroh](https://github.com/n0-computer/iroh) and [WebRTC](https://webrtc.org/) use [UDP Hole Punching to create Peer-to-Peer (P2P) connections](https://tailscale.com/blog/how-nat-traversal-works) between nodes without relaying traffic. 

Could I use P2P and funnel a terminal session over it? Well the Actions VM is on the internet and allows UDP outbound, so it should work!

A simple bit of scripting proved it did 🥳 With WebRTC, if the two nodes exchange information about their connectivity ([ICE Candidates](https://webrtc.org/getting-started/peer-connections#ice_candidates)) then I could form a connection.

## Security and Identities

The next problem is, **how do you prove each end of the P2P connection is who they say they are?**

It's importantly. I want to ensure that `lawrencegripper` can only access terminals for Actions triggered by `lawrencegripper`. 

The browser side is relatively easy, we can use OAuth to login via GitHub and get a verified username ✅

On the Actions VM [we have OIDC](https://docs.github.com/en/actions/concepts/security/openid-connect), commonly used to auth from Actions to cloud providers. 

Anyone can use it though, it gives us the ability to issue a signed OIDC token which lets us confirm:

1. The repo it's running on
2. The user account that triggered it
3. We are the intended audience for the token

You request this token via a REST request in the action:

```typescript
    const requestURL = process.env.ACTIONS_ID_TOKEN_REQUEST_URL;
    const requestToken = process.env.ACTIONS_ID_TOKEN_REQUEST_TOKEN;
    const SERVER_URL = 'https://actions-term.gripdev.xyz';
    const url = new URL(requestURL);
    url.searchParams.set('audience', SERVER_URL);

    const resp = await fetch(url.toString(), {
        headers: {
        Authorization: `Bearer ${requestToken}`,
        Accept: 'application/json',
        },
    });
```

> [Complete code](https://github.com/lawrencegripper/actions-term-on-fail/blob/21c8350bc33a4bf4451473eabecc9d7b2eedc716/client/src/index.ts#L35-L70)

The when the Action calls the server it can include this token and we validate it:

```golang
    const githubOIDCIssuer = "https://token.actions.githubusercontent.com"
    const githubJWKSURL = "https://token.actions.githubusercontent.com/.well-known/jwks"
    // Fetch JWKS
	keySet, err := jwkCache.Get(ctx, githubJWKSURL)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to fetch JWKS: %w", err)
	}

	// Parse and validate token with clock skew tolerance
	parseOpts := []jwtx.ParseOption{
		jwtx.WithKeySet(keySet),
		jwtx.WithIssuer(githubOIDCIssuer),
		jwtx.WithValidate(true),
		jwtx.WithAcceptableSkew(2 * time.Minute),
		jwtx.WithAudience(oidcExpectedAudience),
	}
	token, err := jwtx.Parse([]byte(tokenStr), parseOpts...)
	if err != nil {
		return "", "", "", fmt.Errorf("token validation failed: %w", err)
	}
```

> [Complete code](https://github.com/lawrencegripper/actions-term-on-fail/blob/main/server/main.go#L173-L221)

## Connecting the Peers (ie. [Signaling Server](https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Connectivity#signaling)) 

At this point we know:
1. We can create a connection between two peers (Actions VM and the Users Browser) with WebRTC
2. We have a way to validate the identity of both ends of the connection

What's left is the server to introduce the two peers.

This server doesn't need to handle the data that goes between the two peers, it's only doing introductions. 

As such we can store all we need like this in `go`. 

The browser and the VM both create a [Server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events) connection to the signaling server, providing their OAuth credentials or OIDC to prove their identity. 

```golang
	runIdToSessions            = make(map[string]*Session) // runId -> session
	runIdToSessionsMu          sync.RWMutex
	runIdRunnerSseClient       = make(map[string]*SSEClient) // runId -> SSE client (runner)
	runIdRunnerSseClientsMu    sync.RWMutex
	actorToBrowserSseClients   = make(map[string][]*SSEClient) // actor -> list of browser SSE clients
	actorToBrowserSseClientsMu sync.RWMutex
```

Then we send the Actions VM connectivity details to the browser and the Browser's connectivity details to the Actions VM. 

At this point they establish the Peer-to-Peer connection 🥳

For bonus points, when a new Actions VM connects I can see if there is a brower session from that user available and send them a notification. 

```golang
    runIdRunnerSseClientsMu.Lock()
	runIdRunnerSseClient[runId] = client
	log.Printf("SSE: Runner connected for actor %s (total clients: %d)", actor, len(runIdRunnerSseClient))
	runIdRunnerSseClientsMu.Unlock()

	// Notify browser subscribers about new session
	sess, ok := runIdToSessions[runId]
	if ok {
		notifyNewSession(sess)
	}
```

> [Full Code](https://github.com/lawrencegripper/actions-term-on-fail/blob/255c79feee7d2cbb854144409a93bdd3a03fcdb4/server/main.go#L262-L271)

## Displaying the Terminal

Ok, we're close now. We have the peers established and signaling server to show it. 

What about creating a terminal and streaming the data?

WebRTC has a `datachannel` which you push arbitrary data through. 

On the Actions VM side we create a `pty.Shell` and stream that data over our `datachannel (dc)`

```javascript
    shell = pty.spawn(SHELL, [], {
        name: 'xterm-256color',
        cols,
        rows,
        cwd: process.env.GITHUB_WORKSPACE || process.cwd(),
        env: process.env as Record<string, string>,
    });

    console.log(`PTY started with dimensions ${cols}x${rows}, PID:`, shell.pid);

    shell.onData((shellData) => {
        dc.sendMessage(shellData);
    });
```