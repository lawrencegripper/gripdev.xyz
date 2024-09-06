---
author: gripdev
category:
  - uncategorized
date: "2020-05-11T13:48:27+00:00"
guid: http://blog.gripdev.xyz/?p=1308
title: Cleanup in Bash Scripts
url: /2020/05/11/cleanup-in-bash-scripts/

---
\[Brain dump so I don't forget this one\]

So you want your bash script to exit on an error **but** you'd like it to clean some stuff up before it closes after the error occurs.

No problem a `TRAP` can do this for you (read detailed docs for caveats).

In a very simple form it looks like this:

https://gist.github.com/lawrencegripper/9e778601b2a21d7891e46cf0e1765f46

Using Trap to fire cleanup on exit

Learn more here: [https://www.linuxjournal.com/content/bash-trap-command](https://www.linuxjournal.com/content/bash-trap-command)
