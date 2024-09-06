---
author: gripdev
category:
  - how-to
date: "2015-01-08T11:53:17+00:00"
guid: http://gripdev.wordpress.com/?p=638
title: Gracefully Drain Session from IIS Before Restarting/Stopping
url: /2015/01/08/gracefully-drain-session-from-iis-before-restartingstopping/

---
I needed to do this recently when working with a team to automated releases of several websites.

The team have a fairly common setup. They've got more than one web server behind a load balancer, the load balancer calls a keep alive page to see if the servers are behaving and if they fail to respond stops sending them traffic until they're healthy again.

What I wanted to do was rename the keep alive page, so the box stops getting new traffic, then wait for the existing sessions to complete and finally start updating it.

The second bit is key here, you don't want customers using the site to get dropped mid-upload or form post.

I set about seeing what could be done with powershell and put together this little sample. It uses the IIS CmdLets to see how many sessions are still active and waits for them to complete before stopping IIS.

Thought it was a neat little sample worth sharing.

**EDIT**: I realized that the current script only looked at inflight http requests, I've put together a more complex script to interrogate the Active ASP Net sessions on a given IIS Website. See second file in the gist below, hopefully useful.

https://gist.github.com/lawrencegripper/a1561be7eeac433262fd
