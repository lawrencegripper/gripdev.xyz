---
author: gripdev
category:
  - reactiveextensions
  - windows-phone
date: "2015-05-13T17:00:33+00:00"
guid: https://gripdev.wordpress.com/?p=647
summary: |-
  Hi,

  So I've recently been working on programming for the Microsoft Band. In particular looking at streaming sensor data back from the band in realtime to process on my Windows Phone phone and, depending on the data, push up to a cloud service.

  Out of the box the Band SDK will give you a set of SensorManagers to which you can hook up .NET EventHandlers and then do what you will with the output.

  However, working with streams of data in .NET using EventHandlers is a pain and there is a much nicer technology for dealing with streams -> [Reactive Extensions](http://reactivex.io/). (It's a pretty bit topic, if you haven't heard of it I'd strongly recommend reading up and having a play.)

  So I set about writing a little wrapper to take the input from the Band's SensorManagers and create an IObservable stream of events.

  This lets you do awesome stuff, like doing linq queries over the realtime data stream or time based operations, among other things.
tag:
  - reactiveextensions
title: Streaming Sensor data from the Microsoft Band using Reactive Extensions and C#
url: /2015/05/13/streaming-sensor-data-from-the-microsoft-band-using-reactive-extensions-and-c/

---
Hi,

So I've recently been working on programming for the Microsoft Band. In particular looking at streaming sensor data back from the band in realtime to process on my Windows Phone phone and, depending on the data, push up to a cloud service.

Out of the box the Band SDK will give you a set of SensorManagers to which you can hook up .NET EventHandlers and then do what you will with the output.

However, working with streams of data in .NET using EventHandlers is a pain and there is a much nicer technology for dealing with streams -> [Reactive Extensions](http://reactivex.io/). (It's a pretty bit topic, if you haven't heard of it I'd strongly recommend reading up and having a play.)

So I set about writing a little wrapper to take the input from the Band's SensorManagers and create an IObservable stream of events.

This lets you do awesome stuff, like doing linq queries over the realtime data stream or time based operations, among other things.

Here is a little example of me using the BandReativeExtensionsWrapper to output my average heart rate every 10 seconds as a stream of data.

Below the example is the code for the wrapper, it uses the FromEvent method to build up an IObservable from the EventHandlers. As code creates a subscription to a feed it kicks off the connection to the band and when the subscription is disposed it cleans up and closes the subscription. At the moment it has Accelerometer and HeartRate but as you can see you can easily add the other Sensors to the wrapper as needed using the generic GetSensorStream method. Feel free to have a play!

\[gist https://gist.github.com/lawrencegripper/3f319725189c0c8345c1\]

N.B - The "hrSubscription" variable is an IDisposable, make sure you keep a reference to it as if disposed the stream with stop and disconnect from the SensorManager on the band.

I'll go into a bit more detail on what this is doing under the covers in another post, got some interesting stuff to do with detecting shake/tapping of the band as well.
