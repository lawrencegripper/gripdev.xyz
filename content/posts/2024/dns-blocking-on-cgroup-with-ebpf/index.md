---
author: lawrencegripper
category:
  - programming
  - ebpf
  - cgroups
date: "2024-11-09T07:02:00+00:00"
title: "Preventing Cicular Dependencies in Deployment Systems with cGroups and eBPF"
url: /2024/10/08/dns-blocking-on-cgroup-ebpf
draft: true
---

I work at GitHub on the internal deployment tooling. 

The site is a distributed system, made up of services which together form the functionality of the site and our deployment tooling ships all of these. 

This blog is going to look at the issues of circular dependencies and a recent experiment that I've run to try and prevent them. 

## What is a deploy-time circular dependency in a distributed system?

Let's take a simplified but real example.

You have 2 components, a frontend and a database cluster. 

The frontend is hosting source code, docker images and build systems. 

The database cluster is run on machines and during deployment scripts are executed which pull new configuration, code and containers which together make the database cluster function. 

Have you spotted it yet? If you deploy a broken version to the database cluster it results in the frontend being down. To fix the database cluster you need to deploy a change but the deploy script for the cluster needs to pull assets from the frontend. 

```mermaid
sequenceDiagram
    participant Deploy as Deploy Scripts
    participant Frontend as Frontend API
    participant DB as Database Cluster
    
    Note over Deploy,DB: Normal Deployment Flow
    Deploy->>Frontend: GET /deployment-assets
    Frontend->>DB: Query database
    DB->>Frontend: Database response
    Frontend->>Deploy: Return deployment assets
    Deploy->>DB: Apply updates
    
    Note over Deploy,DB: Failed Deployment Scenario
    Deploy->>DB: Deploy broken update
    DB--xFrontend: Database failure
    Frontend--xDeploy: Cannot serve deployment assets
    Note over Deploy,DB: Circular dependency trap!<br/>Cannot fix DB because<br/>Frontend API is unreachable
```

In the cause of incidents it is responsible for rolling back a bad change too.

