---
author: gripdev
category:
  - uncategorized
date: "2021-06-07T08:26:55+00:00"
guid: https://blog.gripdev.xyz/?p=1497
title: TrueNAS storage controller pass-through with Windows Hyper-V (DDA)
url: /2021/06/07/truenas-storage-controller-pass-through-with-windows-hyper-v-dda/

---
Hyper-v on Server 2019 supports Discrete Device Assignment (DDA) which allows PCI-E devices to be assigned directly to underlying VMs. This through me off as my searches for Device Pass Through didn't return any results!

Typically this is used with Graphics cards and all the docs talk extensively about doing just that. What I wanted to do was pass through an LSI SAS controller to my TrueNAS VM.

Here are my learnings:

1. Enable SR-IOV and I/O MMU [Guide here](http://(https://docs.microsoft.com/en-us/windows-server/virtualization/hyper-v/plan/plan-for-deploying-devices-using-discrete-device-assignment#system-requirements)
1. Start by downloading and running the [Machine Profile Script](https://docs.microsoft.com/en-us/windows-server/virtualization/hyper-v/plan/plan-for-deploying-devices-using-discrete-device-assignment#machine-profile-script). This is going to tell you if you have a machine setup that can support pass-through. If things are good you'll see something like this (but with LSI adapter name not my adapter - my LSI is already setup so it doesn't show here). Make a note of the \`PCIROOT\` portion we'll need that later.

[![](/wp-content/uploads/2021/06/image-1.png)](/wp-content/uploads/2021/06/image-1.png)

1. Use steps 1/2 in a tight loop to make sure your all setup right. My BIOS settings weren't clear, so I did a couple of loops here trying different settings with the Chipset, PCI-E and other bits.
1. Find and disable the LSI Adapter in Device Manager. The easiest way I found to do this is to find a hard drive you know is attacked to the device then switch the device manager view to "by connection" and the hard drive you have selected will now show under the LSI Adapter. Right-click the adapter and click disable (note at this point you'll lose access to the drives). Reboot.
1. Run the following script replacing `$instancePath` with the `PCIROOT` line from the Machine Profile script and `truenas` with your VMs name.

```
$vm = Get-VM -Name truenascore
```

```
$locationPath = "PCIROOT(0)#PCI(0102)#PCI(0000)#PCI(0200)#PCI(0000)"
```

```
Dismount-VmHostAssignableDevice -LocationPath $locationPath -Force -Verbose
Add-VMAssignableDevice -VM $vm -LocationPath $locationPath –Verbose
```

Boot the VM and your done.

Things to note, I tried to pass through the inbuilt AMD storage controller with `-force` even though the Machine Profile script said it wouldn't work. It did kind of work, showing one of the disks but it also made the machine very unstable rebooting the host when the VM was shut down so best to listen to the output of the script and only try to pass through devices that show up green!

I've run now for a couple of days with the LSI adapter passed through and loaded about 2TB onto a RAIDZ2 pool of 5x3TB disks and so far everything is working well.
