---
author: gripdev
category:
  - servicefabric
date: "2016-01-28T18:17:25+00:00"
guid: https://gripdev.wordpress.com/?p=839
tag:
  - servicefabric
title: PubSub in Service Fabric with Redis
url: /2016/01/28/pubsub-in-service-fabric-with-redis/

---
I’ve been working on a project that uses Fabric and I’m hosting Redis inside the cluster as a simple cache system.

One of things that isn’t baked into Fabric is a pub/sub model for communicating between services  about events that are occuring.

As I’ve got the Redis instance up and running in the cluster I decided to take a look at using the Pub/Sub capabilities in Redis to make this happen. N.B Redis isn’t a guarenteed delivery so use where appropriate, there are lots of discussions around it’s pub/sub model and when/where to use etc.

Turns out it’s nice and easy to get working, I’m a big fan of using RX to make nice reactive programs operating on streams of events and there is already a nice sample combineing Redis and RX in C# [here](https://github.com/KidFashion/redis-pubsub-rx/blob/master/src/Redis.PubSub.Reactive/Observable.cs).

In not too long I had just what I wanted and through it might be useful to others so I’ve put together a sample. My sample [is here](https://github.com/lawrencegripper/RedisOnSerivceFabric-Example/tree/master/ExamplePubSub) with one “EventPublisher” service pushing out events and an “EventSubscriber” listening to events.

Both services write out what they’re up to as ETW messages so you can view in the diagnostic window.

[![image](/wp-content/uploads/2016/01/image_thumb8.png)](/wp-content/uploads/2016/01/image8.png)

Lawrence Gripper

@lawrencegripper
