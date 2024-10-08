---
author: lawrencegripper
category:
  - programming
  - docker
  - docker-compose
  - advanced-docker
date: "2024-10-08T07:02:00+00:00"
title: "Faster Docker Builds using Inline Caching"
url: /2024/10/08/advanced-docker-inline-build-cache
---

So you're building a container where it has packages, like `npm ci` or `go mod download`, 
and would you like to make it quicker? For example:

```Dockerfile
COPY go.mod go.sum ./
RUN go mod download
```

Normally when you edit `go.mod` or `go.sum` the `go mod download` will have to redownload
all the packages as its cache was busted by the change in the files.

Well - we can use inline bind mounts to avoid that.

These appear as `--mount` statements in the Dockerfile as part of `RUN` commands. 

They instruct Docker to keep a cache, such as downloaded files, between different runs of `docker build`.

To use them add `--mount=type=cache` to your `RUN` command and specify a target.

```Dockerfile
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod  \
	go mod download
```

Or for Node/NPM

```Dockerfile
RUN --mount=type=cache,target=/root/.npm npm install
```

## Docs 

Here are the Docker docks which go into more detail:
- [🚢 Docker Build: Using cache mounts](https://docs.docker.com/build/cache/optimize/#use-cache-mounts)

{{< catlisttitle category="advanced-docker" title="Mini Series" desc="This is part of a set of posts on useful Advanced Docker techniques" >}}