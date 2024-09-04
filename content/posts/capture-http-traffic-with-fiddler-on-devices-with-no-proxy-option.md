---
author: gripdev
category:
  - how-to
date: "2014-07-05T16:02:03+00:00"
guid: http://gripdev.wordpress.com/?p=559
title: Capture HTTP traffic with Fiddler on Devices with no Proxy option
url: /2014/07/05/capture-http-traffic-with-fiddler-on-devices-with-no-proxy-option/

---
Hi,

This is how I ended up writing a custom DNS server to redirect network traffic.. [code is here](https://github.com/lawrencegripper/FiddlerDnsForwarder "code is here").

I recently got a smart TV and I wanted to see what it was up too, having heard all the stories of them leaking personal information left right and center.

Fiddler is my go to tool for any HTTP inspection, so I went about look for a proxy setting in on the TV ... there isn't one.

To get around this I started pondering and came up with a plan to write my own DNS server in C# which would respond to the TV with the IP address of my machine. On that machine I'd configure Fiddler to run on port 80 and allow remote clients. The upshot of this would be that all traffic from the TV would hit the Fiddler proxy, I would inspect it, then it would get forwarded on to the the internet as normal.

Having a search around I found this nuget package, ARSoft.Tools.Net, from ARSoft which gives me the basics of a DNS server. I then hacked together a bit of code to setup the DNS server to respond as I'd like and bound it up to the correct IPs.

The net result was I could now look at \*some\* of the traffic. The downside is that a lot of apps would break as the HTTPS connections are being forwarded but my machine, and fiddler, but the machine isn't listening or forwarding port 443.

I had a couple of attempts at getting around this, you can see TCPForwarderSlim and PortForwardingWrapper, thanks to Bruno Garcia, in the code. This would setup a straight forward on 443 when a DNS request would come through but it wasn't too reliable. I also playing with using fiddlerCore to listen on a number of ports at the same time but this lead to certificate errors in the apps on the TV as the certs it generated where self signed.

Either way it almost worked, screenshots below, I've put the code up on GitHub for anyone who wants to have a play! https://github.com/lawrencegripper/FiddlerDnsForwarder

Go a rough idea of some of the stuff my TV gets up to...

For example: If you want the TV version of BBC IPlayer set you're UserAgent string to the below and hit http://bbc.co.uk/iplayer

Mozilla/5.0 (DirectFB; Linux armv7l) AppleWebKit/534.26+ (KHTML, like Gecko) Version/5.0 Safari/534.26+ LG Browser/5.00.00(+mouse+SCREEN+TUNER; LGE; 42LS570T-ZB; 04.54.03; 0x00000001;); LG NetCast.TV-2012

[![fiddlersetup](/wp-content/uploads/2014/07/fiddlersetup.png?w=300)](/wp-content/uploads/2014/07/fiddlersetup.png)[![redirecting](/wp-content/uploads/2014/07/redirecting.png?w=300)](/wp-content/uploads/2014/07/redirecting.png)
