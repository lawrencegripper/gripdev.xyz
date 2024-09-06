---
author: gripdev
category:
  - docker
  - healthcheck
  - quick-post
date: "2019-02-11T16:48:32+00:00"
guid: http://blog.gripdev.xyz/?p=1182
tag:
  - docker
  - healthcheck
title: Docker and Healthchecks outside of Kubernetes
url: /2019/02/11/docker-and-healthchecks-outside-of-kubernetes/

---
So I've been working with a containerized solution recently which runs outside of `Kuberenetes` using an Azure VMSS to scale out. I won't dive into the reasons why we went down this route but one really interesting thing came of out of it.

### How do you automatically healthcheck a container outside of Kubernetes?

Well it turns out `docker` has this covered in newer versions. You can [specify a `HEALTHCHECK` inside the docker file to monitor the containers state](https://docs.docker.com/engine/reference/builder/#healthcheck)

### How do you ensure it restarts when unhealthy?

Well here you have a couple of options but both rely on using `--restart=always` when starting the container:

1. You \`healthcheck\` command runs inside the container so you can have it kill the root process of the container causing the container to restart - Example: https://github.com/opencb/opencga/pull/1121/files
1. You can use \`AutoHeal\` container which monitors the docker deamon via it's socket and handles and containers which report unhealthy https://hub.docker.com/r/willfarrell/autoheal/

> Note: I'm trying a new format for shorter slightly rougher blog posts covering specific topics quickly. They'll appear under `Quick-post` tags. Please excuse typos and grammar issues!
