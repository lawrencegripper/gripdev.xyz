---
author: gripdev
category:
  - uncategorized
date: "2019-11-07T00:30:11+00:00"
guid: http://blog.gripdev.xyz/?p=1234
title: Using Azure DevOps to speed up Docker builds
url: /2019/11/07/using-azure-devops-to-speed-up-docker-builds/

---
\[Braindump - warning\]

So I've been playing with `devcontainers` for Visual Studio Code, they're awesome... go play with them. They let you use a `Dockerfile` to describe all the tooling needed for devs to get started with your project.

One of the side effects is that you have a nice `Dockerfile` which you can then also use it for your build server meaning that you never have an inconsistency between your local setup and your CI server.

In this example I build a `golang` project and use Azure DevOps and use caching to minimize the amount of time for each build.

https://gist.github.com/lawrencegripper/5979977b895d345fe688c5c26e99748f
