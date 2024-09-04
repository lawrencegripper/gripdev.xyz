---
author: gripdev
category:
  - quick-post
date: "2019-02-19T13:18:42+00:00"
guid: http://blog.gripdev.xyz/?p=1186
title: Debugging Cloud-Init on Ubuntu (in Azure or anywhere)
url: /2019/02/19/debugging-cloud-init-on-ubuntu-in-azure-or-anywhere/

---
I've recently been working with [`cloud-init`](https://cloud-init.io/) in Azure to setup Ubuntu machines and for the most part I've really like it as it solves lots of problems and fits my use case BUT debugging it has been a pain so I thought I'd write up some notes here for others.

### Did it work?

After deployment SSH onto the node and run this: `cloud-init status -w`  <\- wait for it to finish

`cloud-init status --long` <\- get the results

### Did the stuff I expect get onto the node?

So things didn't go your way, it looks like it hasn't behaved or your just curious. Well this lets you see the contents of the ​ `cloud-init` as it landed on your box.

`sudo cat /var/lib/cloud/instance/user-data.txt.i`

### What exactly failed? I want the verbose logs

So the script didn't run or something else failed and you've got some logging that writes to console in there, never fear this will show you what happened.

`sudo cat /var/log/cloud-init-output.log`

### I'd like to know early if things are broken, can I validate these things on my dev machine?

Yup, this will pick up **some** errors but be warned the validation is limited.

1. Get the tooling (use docker run -it  ubuntu if not on ubuntu): \`sudo apt install cloud-init\`
1. Run the validation: `cloud-init devel schema --config-file your-cloud-init.txt`
1. Profit

### Note

I’m trying a new format for shorter slightly rougher blog posts covering specific topics quickly. They’ll appear under `Quick-post` tags. Please excuse typos and grammar issues!
