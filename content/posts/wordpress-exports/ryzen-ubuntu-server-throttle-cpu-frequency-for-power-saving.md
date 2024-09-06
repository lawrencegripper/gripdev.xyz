---
author: gripdev
category:
  - coding
date: "2023-01-13T09:46:46+00:00"
guid: https://blog.gripdev.xyz/?p=1598
title: 'Ryzen Ubuntu Server: Throttle CPU Frequency for  power saving'
url: /2023/01/13/ryzen-ubuntu-server-throttle-cpu-frequency-for-power-saving/

---
This is a super quick one, I have a Ryzen server running a bunch of VMs. I noticed it's running quite hot and pulling a fair bit of power.

As none of the VMs running are particularly performance sensitive I wanted to force the CPU to use a more conservative power setting.

First up, how do I see what frequencies the CPU is currently running at? For that we'll use `cpufreq-info`, this will tell us what frequency each core is at:

```
cpufreq-info | grep current
  current policy: frequency should be within 2.20 GHz and 3.60 GHz.
  current CPU frequency is 4.23 GHz.
  current policy: frequency should be within 2.20 GHz and 3.60 GHz.
  current CPU frequency is 2.70 GHz.
.... more
```

Now how do I see what power draw this is causing, well that's more complicated. In my case I have an APC UPS connected to the system to keep it up during power outages. This has a tool called `apcaccess` which gives me information about the UPS's load. Knowing the size of the UPS you can back track from the load % to a rough watt's usage.

In our case what I want to do though is use this to prove that changing the CPU has worked. Before making any changes it outputs 19% load.

```
apcaccess | grep LOADLOADPCT  : 19.0 Percent
```

To throttle the CPU we can use a `govenor` lets see what governors we have available

```
cpufreq-info | grep governors
  available cpufreq governors: conservative, ondemand, userspace, powersave, performance, schedutil
  available cpufreq governors: conservative, ondemand, userspace, powersave, performance, schedutil
  available cpufreq governors: conservative, ondemand, userspace, powersave, performance, schedutil
  available cpufreq governors: conservative, ondemand, userspace, powersave, performance, schedutil
```

Cool, well `powersave` looks like a good one to try out, lets give that a go by running `sudo cpupower frequency-set --governor powersave` and then looking at the load and frequencies again.

```
lawrencegripper@libvirt:~$ sudo cpupower frequency-set --governor powersave
Setting cpu: 0
Setting cpu: 1
Setting cpu: 2
....
Setting cpu: 15
lawrencegripper@libvirt:~$ cpufreq-info | grep current
  current policy: frequency should be within 2.20 GHz and 2.20 GHz.
  current CPU frequency is 2.20 GHz.
  current policy: frequency should be within 2.20 GHz and 3.60 GHz.
  current CPU frequency is 2.20 GHz.
  current policy: frequency should be within 2.20 GHz and 3.60 GHz.
  current CPU frequency is 2.20 GHz.
....
lawrencegripper@libvirt:~$ apcaccess | grep LOAD
LOADPCT  : 15.0 Percent
```

That did the trick, `apcaccess` is reporting a drop to 15% load and CPU frequency is down to 2.2GHz.

I've got a smart meter for my home and can also see the drop in usage roughly reflected there too.

\[Edit:\] I also found this isn't persisted between reboots. The following shows how to persist the change https://askubuntu.com/questions/410860/how-to-permanently-set-cpu-power-management-to-the-powersave-governor

Done. Server is now being more eco-friendly.
