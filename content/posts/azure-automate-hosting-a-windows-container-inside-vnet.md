---
author: gripdev
category:
  - containers
date: "2021-03-17T16:11:32+00:00"
guid: https://blog.gripdev.xyz/?p=1439
tag:
  - containers
  - docker
  - dsc
  - powershell
  - windows-containers
title: 'Azure: Automate hosting a Windows Container inside VNET'
url: /2021/03/17/azure-automate-hosting-a-windows-container-inside-vnet/

---
The quickest and easiest way to start running a Windows Container in Azure is using Azure Container Instances (ACI).

The problem is [that they currently (as of 03/21) don't support running Windows Containers inside a VNET.](https://docs.microsoft.com/en-us/azure/container-instances/container-instances-region-availability#windows-container-groups) This blog is about how I worked around this limitation by automating the deployment and management of a Windows Containers with PowerShell DSC and Terraform.

As we needed VNET integration for the sensitive data handled on the project, I set out to build an ACI like experience on a VM and have that connected to a VNET in Azure.

**Warning: [Please ensure you fully understand how PowerShell DSC works and review the code in full as there is complexity in this approach](https://docs.microsoft.com/en-us/powershell/scripting/dsc/reference/resources/windows/scriptResource?view=powershell-7.1&viewFallbackFrom=powershell-7.).**

First up it must:

- Handle restarting the container if things go wrong
- Give an easy way to retrieve the container logs from the commandline
- Connect reliably to the VNET
- Support updating easily (ie. When I push a new image tag the container is restarted running the new version **or** when new environment variables are applied handle restarting the container to pick these up)
- Support authentication to an Azure Container Repository
- Be runnable as part of a Terraform Deployment

Seems like a pretty long list right? At this point I reached out to a friend, [Marcus Robison,](https://twitter.com/techdiction) who'd done more Windows Admin than me in his time. He suggested looking at Powershell DSC.

So what is Powershell DSC, what does it give us?

- Desired State Configuration for the VM. "I want a VM that looks like x" and it makes that happen. Much like Terraform or a K8s operator. It queries the current state and takes actions that move the current state closer to the desired state
- Integration with Azure VMs. There is a nice extension in Azure which allows you to submit a DSC config and Azure manages starting it on the VM for you.
- Handling of sensitive variables securely, with the Azure extension variables are encrypted

What does this all look like when you have it finished?

## `dsc_config.ps1`

This script is responsible for configuring the machine, logging into ACR and ensuring the container is running. This runs periodically and handles things like restarting the [container if a new deployment has been made with update environment variables](https://gist.github.com/lawrencegripper/c06239f37ace287ce44e4bf36dd6ee2f#file-dsc_config-ps1-L223-L238).

[Each Script (think resource in terraform](https://docs.microsoft.com/en-us/powershell/scripting/dsc/reference/resources/windows/scriptResource?view=powershell-7.1&viewFallbackFrom=powershell-7.)) has a **Get, Set and Test** method. **Test** checks the current state of things, if they're not how they're meant to be **Set** is responsible for getting them configured correctly and lastly **Get** returns an identifier for the item.

## `module.tf`

This is the terraform responsible for creating the VM and pushing up the DSC script for it to run.

It takes the `dsc_config.ps1` and creates a zip file, this zip is then passed to the PowerShell DSC extension for the Azure VM which is responsible for applying the configuration to the VM.

As well as this the module also takes the environment variables you want set for your container. These are provided as a map and converted to a base64 encoded .env file. The DSC config on the VM decodes them and provides the `.env` file to the `docker run` command used to start the container.

\*Worth nothing env.tpl is used in the process of creating the env file.

## `usage.tf`

This is an example of using the terraform module from `module.tf` to create a VM which runs a container image on a VNET with a set of environment variables.

## `getlogs.ps1`

Once deployed this little script demonstrates how you can get the logs out from the container running in the VM. It requires the Azure CLI to be installed and you to provide the VM's Azure ID.

You can also hook this up to the outputs of your Terraform to automate it further.

## All together now!

https://gist.github.com/lawrencegripper/c06239f37ace287ce44e4bf36dd6ee2f
