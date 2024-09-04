---
author: gripdev
category:
  - ci
  - kind
  - kubernetes
date: "2020-01-15T12:54:49+00:00"
guid: http://blog.gripdev.xyz/?p=1250
tag:
  - ci
  - kind
title: Use KIND (Kubernetes in Docker) in CI/CD reliably
url: /2020/01/15/use-kind-kubernetes-in-docker-in-ci-cd-reliably/

---
I've been working with OPA recently and using KIND to test things out. This works really nicely but when I started using the same approach in CI I saw some errors.

Digging into things you can see that the nodes of the KIND cluster aren't "READY" when the CLI finishes up so you need a bit of extra bash foo to make the process wait on the READY status.

This monster line of bash does the trick:

> `JSONPATH='{range .items[*]}{@.metadata.name}:{range @.status.conditions[*]}{@.type}={@.status};{end}{end}'; until kubectl get nodes -o jsonpath="$JSONPATH" 2>&1 | grep -q "Ready=True"; do sleep 5; echo "--------> waiting for cluster node to be available"; done`

In this example I'm also deploying a K8s operator this needs be to up and running before I can run the integration tests , a similar bit of bash ensures that's true too:

> `JSONPATH='{range .items[*]}{@.metadata.name}:{range @.status.conditions[*]}{@.type}={@.status};{end}{end}'; until kubectl -n opa -lapp=opa get pods -o jsonpath="$JSONPATH" 2>&1 | grep -q "Ready=True"; do sleep 5;echo "--------> waiting for operator to be available"; kubectl get pods -n opa; done`

If you put those all together I can have a nice `make` file which:

1. Deploys a KIND clsuter
1. Wait for it to work
1. Deploys Open Policy Agent
1. Waits for it to be running
1. Runs my python integration tests

All by running `make kind-integration` :)

Full file:

https://gist.github.com/lawrencegripper/c7b2d6324afcdaf2ebe2f9a2be56be7a
