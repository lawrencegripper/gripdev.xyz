---
author: gripdev
category:
  - paperboy
date: "2014-06-05T15:30:24+00:00"
guid: http://gripdev.wordpress.com/?p=497
title: How I unwittingly became part of a DDOS attack
url: /2014/06/05/how-i-unwittingly-became-part-of-a-ddos-attack/

---
I recently stood up a azure VM to test out some work I've been doing around creating my own DNS server.

Feeling bold and overconfident I created a machine with a long random name, punched open the firewall for dns and set my custom dns running. Having played with the code till late in the evening I left the box running and went to bed. My machines existence was only known to me, so I just left it there with unfinished code answering dns queries.

The next evening, when I set about to continue my work, I noticed that the server was actually already really busy. I was receiving LOADS of requests.

Looking at the apparent source IP it was from a Russian hosting company. However, as DNS is UDP, spoofing the source IP is possible and therefore it can be used it to amplify DDOS attacks. You can read more on this attach type here, http://blog.cloudflare.com/deep-inside-a-dns-amplification-ddos-attack

So it turns out that overnight my machine had apparently been found by a bot crawling IPs and conscripted into a DNS DDOS attack.

Lesson learned, security through obscurity doesn't cut it on the internet!
