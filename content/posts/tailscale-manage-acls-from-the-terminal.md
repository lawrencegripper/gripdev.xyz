---
author: gripdev
category:
  - aws
  - azure
  - devops
  - gitops
  - kubernetes
  - uncategorized
date: "2024-03-22T10:46:29+00:00"
guid: https://blog.gripdev.xyz/?p=1725
tag:
  - aws
  - azure
  - devops
  - gitops
  - kubernetes
title: 'Tailscale: Manage ACLs from the Terminal'
url: /2024/03/22/tailscale-manage-acls-from-the-terminal/

---
Tailscale supports using [a GitOps flow](https://tailscale.com/kb/1204/gitops-acls) to manager ACLs for you're tailnet.

This involves configuring a [GitHub Action](https://tailscale.com/kb/1306/gitops-acls-github) then you commit the ACLs and the Action runs.

Pretty cool right? I can see this being a great flow for larger Tailnets managed by multiple people.

For me tho, with my homelab Tailnet, this was a bit heavy weight. What I wanted was:

1. ACLs tracked in my HomeNet git repo which has all my homelab config
1. Ability to easily edit the file and apply without going to the Tailscale admin page and copy/paste dance

Knowing that the GitHub action was doing effectively this I started there.

When digging through the source code of the GitHub Action I found this Go CLI util [`gitops-pusher`](https://pkg.go.dev/tailscale.com/cmd/gitops-pusher). With a bit of playing this did exactly what I wanted.

This little script uses `gitops-pusher` to update my ACLS (note: I use 1Password injection to pull secrets, [talk about this more here](/2024/02/26/homelab-using-1password-cli-to-handle-secrets-in-kubernetes-compose-yaml/))

```
#!/bin/bash

if ! command -v gitops-pusher &> /dev/null; then
    go install tailscale.com/cmd/gitops-pusher@gitops-1.58.2
fi

export TS_TAILNET={{ op://Personal/homenet/tailscale/tailnet-name }}
export TS_OAUTH_ID={{ op://Personal/homenet/tailscale/acl-oauth-client-id }}
export TS_OAUTH_SECRET={{ op://Personal/homenet/tailscale/acl-oauth-secret }}

gitops-pusher --policy-file ./tailscale/acls.hujson $1
```

I can then call this from my `make` file like so:

```
tailscale-test-acls:
	cat ./tailscale/update-acls.sh | op inject | bash -s -- testtailscale-push-acls:
	cat ./tailscale/update-acls.sh | op inject | bash -s -- apply
```

That's it, all done!

[![](/wp-content/uploads/2024/03/image-6.png)](/wp-content/uploads/2024/03/image-6.png)

Bonus: Tailscale ACL files are `hujson` format, allowing comments and tailing commas. You can setup VSCode to recognize this as valid to avoid it showing errors when editing the files, [instructions here](https://github.com/tailscale/hujson?tab=readme-ov-file#visual-studio-code-association).
