---
author: gripdev
category:
  - how-to
date: "2018-09-05T19:31:17+00:00"
guid: http://blog.gripdev.xyz/?p=1163
title: 'How to: Check your DNS entry maps to your Public IP in Bash'
url: /2018/09/05/how-to-check-your-dns-entry-maps-to-your-public-ip-in-bash/

---
I wrote this today as I wanted to ensure that a service waiting for its DNS name to be updated with the correct IP address (its Public IP) before starting.

This little script uses Curl with Akamai's 'whatsismyip.akamai.com' endpoint to get the Public IP and then NSLookup to get the IP returned by the DNS server for the domain. It keeps trying for a while until they match or exits if they don't match after 250 seconds.

_WARNING:_ In my case it turned out that outbound traffic didn't route through the same IP as inbound so the script always failed. This may happen to you too if you're using this in K8s.

_WARNING:_ The AWK logic extracting the IP from the NSLookup is brittle is expects result on line 5. This works on Alpine but may need tweaking, likely are better approaches here.

Run "dnscheck.sh mydns.name.here"

https://gist.github.com/lawrencegripper/58eaa1dd8ed858557b3165382e9306bb
