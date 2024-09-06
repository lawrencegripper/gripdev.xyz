---
author: gripdev
category:
  - devcontainers
  - kind
  - vscode
date: "2020-02-19T21:04:45+00:00"
guid: http://blog.gripdev.xyz/?p=1263
summary: |-
  I found myself last week looking at a bit of code in K8s which I thought I could make better, so I set about trying to understand how to clone, change and test it.

  Luckily K8s has some good docs [,](https://github.com/kubernetes/community/tree/master/contributors/devel#readme) [trust these over me as they're a great gui](https://github.com/kubernetes/community/tree/master/contributors/devel#readme) [de.](https://github.com/kubernetes/community/tree/master/contributors/devel#readme) This blog is more of a brain dump of how I got on trying with Devcontainers and VSCode.This is my first try at this so I've likely got lots of things wrong.

  Roughly I knew what I needed as I'd heard about:

  1. Bazel for the Kubernetes build
  2. [Kind](https://github.com/kubernetes-sigs/kind) to run a cluster locally
  3. [Kubetest](https://github.com/kubernetes/test-infra/tree/master/kubetest) for running end2end tests

  As the K8s build and testing cycle can use up quite a bit of machine power I didn't want be doing all this on my laptop and ideally I wanted to capture all the setup in a nice repeatable way.

  Enter [DevContainers for VSCode](https://code.visualstudio.com/docs/remote/containers), I've [got them setup on my laptop to actually run on a meaty server for me](https://blog.gripdev.xyz/2019/09/17/vscode-container-dev-a-remote-docker-daemon-brain-dump/) and I can use them to capture all the setup requirements for building K8s.
tag:
  - devcontainers
  - kind
  - vscode
title: How to build Kubernetes from source and test in Kind with VSCode &amp; devcontainers
url: /2020/02/19/how-to-build-kubernetes-from-source-and-test-in-kind-with-vscode-devcontainers/

---
I found myself last week looking at a bit of code in K8s which I thought I could make better, so I set about trying to understand how to clone, change and test it.

Luckily K8s has some good docs [,](https://github.com/kubernetes/community/tree/master/contributors/devel#readme) [trust these over me as they're a great gui](https://github.com/kubernetes/community/tree/master/contributors/devel#readme) [de.](https://github.com/kubernetes/community/tree/master/contributors/devel#readme) This blog is more of a brain dump of how I got on trying with Devcontainers and VSCode.This is my first try at this so I've likely got lots of things wrong.

Roughly I knew what I needed as I'd heard about:

1. Bazel for the Kubernetes build
1. [Kind](https://github.com/kubernetes-sigs/kind) to run a cluster locally
1. [Kubetest](https://github.com/kubernetes/test-infra/tree/master/kubetest) for running end2end tests

As the K8s build and testing cycle can use up quite a bit of machine power I didn't want be doing all this on my laptop and ideally I wanted to capture all the setup in a nice repeatable way.

Enter [DevContainers for VSCode](https://code.visualstudio.com/docs/remote/containers), I've [got them setup on my laptop to actually run on a meaty server for me](/2019/09/17/vscode-container-dev-a-remote-docker-daemon-brain-dump/) and I can use them to capture all the setup requirements for building K8s.

After a lot of wrangling I got a `Dockerfile` and `devcontainer.json` file together which mount and install the required bits (warning: There may be lots wrong with this, I'm pretty new to the world of building K8s from source).

Now all I need to do is clone K8s, drop these two files into a `.devcontainer` folder and open the folder in VSCode. Once open you get prompted to open the folder in the devcontainer and away you go.

## What do you do now?

- Make the change you want to make
- Build an image for kind with your changes:

```
kind build node-image --type bazel \
 -v 5 --image kindest/node:lawrence
```

- Start a cluster based off that image:

```
kind create cluster --name lg \
--image kindest/node:lawrence
```

- You can run the e2e tests too (still getting to grips with this)

```
kubetest --up --test --down \
--deployment=kind \
--kind-node-image=kindest/node:lawrence \
--provider=skeleton
```

The best part? All of this runs nicely in VSCode like its running locally but behind the scenes it's on a remote server. The left window is my laptop and the right is the server doing a build.

![](/wp-content/uploads/2020/02/devcontainer.jpg?w=1024)

## Let me try it

Here is a [Gist of my devcontainer files for you to play with. If you spot something silly I'm doing please let me know!](https://gist.github.com/lawrencegripper/0027a76b90539a9534a87c136c35c484)

https://gist.github.com/lawrencegripper/0027a76b90539a9534a87c136c35c484

Current devcontainer setup (WIP)
