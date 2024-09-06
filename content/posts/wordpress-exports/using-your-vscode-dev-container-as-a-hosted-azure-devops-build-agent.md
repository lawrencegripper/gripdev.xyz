---
author: gripdev
category:
  - uncategorized
date: "2021-03-15T11:12:45+00:00"
guid: https://blog.gripdev.xyz/?p=1428
title: Using your VSCode dev container as a hosted Azure DevOps build agent
url: /2021/03/15/using-your-vscode-dev-container-as-a-hosted-azure-devops-build-agent/

---
[Devcontainers are awesome for keeping tooling consistent over the team](https://code.visualstudio.com/docs/remote/create-dev-container), so what about when you need to run your build?

There is some great work already done [talking about how to use these as part of a normal pipeline](https://dev.to/eliises/dockerizing-devops-39hk) ( [shout out to Eliise!](https://dev.to/@eliises)), what about if you need your build agent to be inside a virtual network in Azure?

The standard approach [would be to create a VM, setup tools and join that as an Agent to Azure Devops](https://docs.microsoft.com/en-us/azure/devops/pipelines/agents/v2-linux?view=azure-devops).

As we've already got a definition of the tooling we need, our devcontainer, can we reuse that to simplify things?

Turns out we can, using an Azure Container Repository, Azure Container Instance and a few tweaks to our devcontainer we can spin up an agent for Devops based on the devcontainer and start using it.

To do this we need to:

1. Add the AzureDevops Agent script to your devcontainer
1. Build the image and push up to your [Azure Container Repository following this guide](https://docs.microsoft.com/en-gb/azure/container-registry/container-registry-get-started-docker-cli)
1. Use Terraform to deploy the built container into an [Azure Container Instance](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_group)

The snippets below assume you already have your agent built and pushed up to your Azure Container Repository with the name `your_repo_name_here.azurecr.io/devcontainer:buildagent`.

It shows the `.Dockerfil` e for the devcontainer the `bash` script to start a devcontainer ( [slight edit from doc here](https://docs.microsoft.com/en-us/azure/devops/pipelines/agents/docker?view=azure-devops#linux)) and the terraform to deploy it into a VNET.

You'll have to do some tweaks, best to treat this as a starting point. [See this doc for more detailed docs on how this work.](https://docs.microsoft.com/en-us/azure/devops/pipelines/agents/docker?view=azure-devops#linux)

https://gist.github.com/lawrencegripper/2be62a42a56e27622b4e81e9210e99bd
