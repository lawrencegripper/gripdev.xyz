---
author: gripdev
category:
  - security
date: "2017-07-17T11:45:06+00:00"
guid: https://gripdev.wordpress.com/?p=936
tag:
  - security
title: 'Kubernetes: Prevent Container from accessing Cluster API'
url: /2017/07/17/kubernetes-prevent-container-from-accessing-cluster-api/

---
I've recently been playing with Kubernetes as way to efficiently host my microservices. It's great and I'm really enjoying it.

One thing that did take me by surprise is that Containers, [by default, have credentials mounted which allow them to talk to the cluster API.](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/) This is incredibly useful for pulling information about other services, secrets or enabling self-orchestrating/scaling services, however, for the services that don't need this it, it presents on opportunity for an attacker to escalate their privileges.

We tend to think about the containers (roughly) providing limited access to resources, information and preventing access to other containers but with this token the container is able to reach beyond the container. The mitigation is that, first, the attacker would have to compromise the application running in the container, once they had done this they could extract and use the token to control or pull information from your cluster.

As always, defense in depth is a good policy. We aim to prevent our container from being compromised but we ALSO aim to limit the damage if it is.

So here is how to prevent that service account being mounted into your containers which don't require it:

```
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  serviceAccountName: build-robot
  automountServiceAccountToken: false
  ...

```

[RBAC also presents an opportunity](https://kubernetes.io/docs/admin/authorization/rbac/) to have finer grained control over the level of access and is almost certainly something to consider too along with the use of namespaces.

## How could an attacker get access?

Here is a really simple example of a badly written nodejs app which allows the users to pass in a version, this is used to find a file and the contents of the file is then returned to the user.

The k8s token is mounted at /var/run/secrets/kubernetes.io/serviceaccount

Using the '../' syntax we can traverse out of the apps folder and down to the root then we can request the token, ca and namespace files from the service account.

https://gist.github.com/lawrencegripper/2fe9caf84863c70b5f83bc81ff94393a

This is simple example to show the files exist and can be leaked. Other vulnerabilities with remote code execution would allow the attacker to make requests to the cluster API using this token.

## Summary

Kubernetes has a great architecture which lets containers talk to the cluster API but in some scenarios, when this access isn't required, you're likely to want to turn this off to provide [defense in depth](https://en.wikipedia.org/wiki/Defense_in_depth_%28computing%29) by the [principal of least privileged.](https://en.wikipedia.org/wiki/Principle_of_least_privilege)
