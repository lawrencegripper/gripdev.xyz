---
author: gripdev
category:
  - how-to
date: "2014-12-10T13:09:14+00:00"
guid: http://gripdev.wordpress.com/?p=590
title: Chocolatey, PowerShell DSC and Azure IAAS - Automating dev box creation
url: /2014/12/10/chocolatey-powershell-dsc-and-azure-iaas-automating-dev-box-creation/

---
Hi,

So back in 2013 I [wrote a post](http://wp.me/p1He68-5u) on automating the creation of a Virtual Machine, and installation of all my loved bits of software, using Chocolaty and Remote Powershell. Well things have moved on in the Azure IAAS world and we have a nice new way to automate installation and configuration of VM's.

PowerShell DSC (desired state configuration) gives you a nice declarative way to define the setup of your machine. It also allows you to write custom modules to extend its functionality, which is just what I've done for Chocolatey. [It's here on GitHub.](https://github.com/PowerShellOrg/cChoco)

So now you can add in all the nice apps you want on the box is one simple DSC config, along with windows features etc.

Below is an example of using the cChoco module to install Git on a new Azure VM. I've broken down the script into sections with a full version at the bottom.

Brief description of Sections:

1. This is an example of a simple DSC config that installs IIS on the machine where it runs, for reference.
1. Pulls down the custom cChoco module from Github, this needs to be copied onto your box so it can be used later.
1. Creates a DSC config using the module which installs git from Chocolatey. You can add more "cChocoPackageInstaller" nodes, changing the Name parameter to be the package you'd like to install or add "WindowsFeature" nodes to install IIS and other bits.
1. Takes this DSC config and pushing it up to Azure, ready to apply to my nice new VM.
1. Create a VM adding in the DSC extension and providing the DSC Config name and file from step 4

Finally at the bottom of the Gist is a full script that will combine all these steps and do some extra checks and goodness.

Happy Automating!

@lawrencegripper

PS. You can check on the install progress through the portal, screenshot below the Gist.

{{< gist lawrencegripper 770c78b12bc752c5c307 >}}

Review Progress in Preview Portal:

[![DSCStatus](/wp-content/uploads/2014/12/dscstatus.png?w=300)](/wp-content/uploads/2014/12/dscstatus.png)
