---
author: gripdev
category:
  - uncategorized
date: "2023-02-21T14:49:40+00:00"
draft: "false"
guid: https://blog.gripdev.xyz/?p=1615
title: 'Tailscale + Synology: Connect to machines from another Subnet Router'
url: /

---
Quick note to hopefully help others.

If you try and run `tailscale up --accept-routes` from the synology nas it will say

```
user@nas:~$ tailscale up --accept-routes
--accept-routes is not supported on Synology; see https://github.com/tailscale/tailscale/issues/1995
```

This means the NAS can't connect to other subnets published by another subnet router on your tailnet.

To fix this you first need to enable outbound connections via tailscale, [follow this guide](https://tailscale.com/kb/1131/synology/#enabling-synology-outbound-connections).

Once done you should see `tailscale0` as a device

```
nas:~$ ip addr show
1: lo: mtu 65536 qdisc noqueue state UNKNOWN
link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
inet 127.0.0.1/8 scope host lo
valid_lft forever preferred_lft forever
inet6 ::1/128 scope host
valid_lft forever preferred_lft forever
2: eth0: mtu 1500 qdisc mq state UP qlen 1024
link/ether 00:11:32:5c:ad:a2 brd ff:ff:ff:ff:ff:ff
inet 10.99.1.5/24 brd 10.99.1.255 scope global eth0
valid_lft forever preferred_lft forever
inet6 fe80::211:32ff:fe5c:ada2/64 scope link
valid_lft forever preferred_lft forever
3: sit0: mtu 1480 qdisc noop state DOWN
link/sit 0.0.0.0 brd 0.0.0.0
53: tailscale0: mtu 1280 qdisc pfifo_fast state UNKNOWN qlen 500
link/none
inet 100.84.105.62/32 scope global tailscale0
valid_lft forever preferred_lft forever
inet6 fd7a:115c:a1e0:ab12:4843:cd96:6254:693e/128 scope global
valid_lft forever preferred_lft forever
```

With that setup you can now create a static route to the host on your subet via the subnet router and traffic will go through. Where `100.3.3.3` is the tailscale address of the subnet router and `10.0.1.14` is the address you want to access on it's subnet.

nas:~$ sudo ip route add 10.0.1.14 via 100.3.3.3 dev tailscale0  
Password:  
gripper@nas:~$ sudo ping 10.0.1.14  
PING 10.0.1.14 (10.0.1.14) 56(84) bytes of data.  
64 bytes from 10.0.1.14: icmp\_seq=1 ttl=63 time=593 ms  
64 bytes from 10.0.1.14: icmp\_seq=2 ttl=63 time=54.6 ms  
64 bytes from 10.0.1.14: icmp\_seq=3 ttl=63 time=81.0 ms
