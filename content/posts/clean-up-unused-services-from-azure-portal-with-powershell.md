---
author: gripdev
category:
  - how-to
date: "2014-03-06T20:30:46+00:00"
guid: https://gripdev.wordpress.com/?p=480
title: Clean up unused Services from Azure Portal with PowerShell
url: /2014/03/06/clean-up-unused-services-from-azure-portal-ghost-services/

---
So I've recently been using lots of VM's and I end up getting through them fairly quickly – deleting and creating them regularly.

What I ended up with was loads of ghost services in the portal that were, once upon a time, associated with a VM in the distant past.

Given PowerShell is awesome, I set about automating the removal of these using the azure cmdlts. The trick is that I don't want to remove services or VM's that are currently "Stopped Deallocated" only those ghost services.

The script has three stages:

1. Find all VM's and their associated services.
1. Find all the services without active deployments.
1. Identify if any of these empty services have VM's associated with them.
1. Perform remove action on these services.

{{< gist lawrencegripper 7a284e635a964b127118 >}}
