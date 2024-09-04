---
author: gripdev
category:
  - how-to
date: "2014-03-01T09:30:09+00:00"
guid: http://gripdev.wordpress.com/?p=473
title: Start and Stop Azure VM’s (and more) in Parallel from Powershell
url: /2014/03/01/start-and-stop-azure-vms-and-more-in-parallel-from-powershell/

---
**Edit 2017:** With the change to AzureRM this approach need some tweaking, [see this new GIST for how to](https://gist.github.com/lawrencegripper/9876ca03e82d4af26a0017a0ab45346e) allow AzureRM to execute commands in Powershell jobs. The script itself is used to remove deployment history for a resource group but can easily be edited to work on VMs.

So I've got a demo that I've put together that uses a LOAD of Azure VM's to demonstrate scripting against them.

To make sure I don't get charged for them while I'm not using them I've got the following powershell to start and stop them:

Get-AzureVM \| Stop-AzureVM –Force

Get-AzureVM \| Start-AzureVM

I notice that this took a longer than I was expecting and so I started to do some digging. It turn out the PowerShell pipeline is executing these API requests in a synchronous way. Firing off the first request, waiting for the response and then, when completed, moving on to the next. This wouldn't be a problem if I didn't have loads of VM's!

So I wrote the following, it uses PowerShell Jobs (think Task<T> if you're a C# person), to kick these off as background jobs that run in parallel. (nb. You have to pass the variables in as arguments to the script block)

{{< gist lawrencegripper 9876ca03e82d4af26a0017a0ab45346e >}}

Now we get this lovely output. These all run as individual jobs and complete asynchronously, we then wait for all of them (think WaitAll() for C#). Now I can start all of my machines in the demo really quickly!

This is really useful in other scenarios too, think of any time you're waiting on azure management API and the actions could be done simultaneously.

Id     Name            PSJobTypeName   State         HasMoreData     Location             Command

--     ----            -------------   -----         -----------     --------             -------

22     Job22           BackgroundJob   Running       True            localhost            ...

24     Job24           BackgroundJob   Running       True            localhost            ...

26     Job26           BackgroundJob   Running       True            localhost            ...

28     Job28           BackgroundJob   Running       True            localhost            ...

30     Job30           BackgroundJob   Running       True            localhost            ...

RunspaceId           : 88465362-c8ce-45de-866d-9ced85cc3e40

OperationDescription : Start-AzureVM

OperationId          : 60712d20-311b-6cd5-b0df-9a750f024d0c

OperationStatus      : Succeeded

RunspaceId           : 3a9e93cf-7d16-4a2c-aeb7-4493e8ec8abb

OperationDescription : Start-AzureVM

OperationId          : fb494c0c-3728-62d4-b288-b8d391ad32d7

OperationStatus      : Succeeded

RunspaceId           : c1144e74-ba7c-4aa3-bcbb-aa9f2bd72fdf

OperationDescription : Start-AzureVM

OperationId          : 194254da-be36-6c7d-a147-67a980816c8e

OperationStatus      : Succeeded

RunspaceId           : 44af74a2-fc11-4e9b-8158-d9d84dd49d9e

OperationDescription : Start-AzureVM

OperationId          : 86059d5d-22ee-6a55-871d-28bcf52241a2

OperationStatus      : Succeeded

RunspaceId           : a44b51e9-11cc-4d31-875e-5e0fd16b935a

OperationDescription : Start-AzureVM

OperationId          : d4e2fcce-b270-6d6e-ad0c-8322e2c4df23

OperationStatus      : Succeeded
