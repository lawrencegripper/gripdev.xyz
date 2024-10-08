---
author: lawrencegripper
category:
  - programming
  - docker
  - docker-compose
  - advanced-docker
date: "2024-10-08T07:02:00+00:00"
title: "Safer Docker Builds: Using Secrets to temporarily mount sensitive info"
url: /2024/10/08/advanced-docker-secrets
---

So you have a docker build which requires access to an authenticated resource?

Docker Secrets can help do this cleanly!

## Temporarily provide build-time secrets

There are times when you want to provide a secret, or file with secrets in it, for use in your docker build.

Now if you `COPY secert.yaml /etc/secrets` that's going to end up in your public image 🫢😨 Also, secret are commonly outside the `context` of the docker build, say it's `$HOME/.netrc` making this a pain.

You should try out `secrets` in docker, they solve this problem.

You can mount a secret for a single line of a `Dockerfile`, use it, and then it's gone.

Take an example: `go mod` is pulling code from a private repo

We can do the following 👇

```Dockerfile
COPY go.mod go.sum ./
RUN --mount=type=secret,id=netrc,target=/root/.netrc \
	go mod download
```

The `--mount=type=secret,id=netrc,target=/root/.netrc` tells docker to mount the secret temporarily during
this `RUN` command and place the content at `/root/.netrc`.

We need to wire this up during builds.

In our `docker-compose.yaml` file we can then let docker know where to get the file by using a top level 
`secrets` entry and adding an entry under `build` for the container. 

```yaml
secrets:
  netrc:
    file: $HOME/.netrc
services:
  windmill-host:
    build:
      context: .
      dockerfile: ./services/some-machine/Dockerfile
      secrets:
        - netrc
    container_name: some-machine
```

If you're using straight docker you can still use this by adding the secret to the command line. 

```
docker build --secret id=netrc,src=$HOME/.netrc .
```

## Docs

Here are the Docker docks which go into more detail:
- [Build Secrets Docker CLI](https://docs.docker.com/build/building/secrets/)
- [Build Secrets Docker Compose](https://docs.docker.com/compose/how-tos/use-secrets/)

{{< catlisttitle category="advanced-docker" title="Mini Series" desc="This is part of a set of posts on useful Advanced Docker techniques" >}}