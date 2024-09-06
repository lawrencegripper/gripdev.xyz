---
author: gripdev
category:
  - how-to
date: "2016-01-28T17:20:25+00:00"
guid: https://gripdev.wordpress.com/?p=841
title: Running Redis in a Service Fabric cluster
url: /2016/01/28/running-redis-in-a-service-fabric-cluster/

---
**\[Update 2021**\] This approach is now very out-dated.

I’m working on a project where we’re in need of a cache inside Service Fabric. Now the cache doesn’t need to be durable just nice, quick and easy to use.

So I came up with the idea of running a Redis cache inside the cluster as a stateless service fabric service. The problem with this is that when the cluster re-balances the services between nodes I will lose the inmemory data in Redis (all of it). What I can do is set the movement cost set to 'high', to discourage the cluster from reallocating the service between nodes regularly and causing unnecessary loss of the cache. Great read on how and why the Resource Manager allocations services between nodes [here](http://blogs.msdn.com/b/azureservicefabric/archive/2015/12/15/service-fabric-under-the-hood-the-cluster-resource-manager-part-1.aspx), if you'd like to know more.

It’s worth noting that, if I wanted a durable cache, I could look at using Redis Clustering/Redis as a Service in Azure to prevent this occurring but for what we need it’s overkill.

Either way, I got this up and running and have shared the code here with a ReadMe going into more detail for others that are interested in having a play. Hope it's useful!

https://github.com/lawrencegripper/RedisOnSerivceFabric-Example

Lawrence Gripper

@lawrencegripper
