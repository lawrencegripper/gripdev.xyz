---
author: gripdev
category:
  - '#arm'
  - '#terraform'
  - coding
date: "2018-06-21T10:49:25+00:00"
guid: http://blog.gripdev.xyz/?p=1149
tag:
  - '#arm'
  - '#terraform'
title: 'From ARM to Terraform: Why I''m becoming a Terraform convert'
url: /2018/06/21/from-arm-to-terraform-why-im-becoming-a-terraform-convert/

---
Let me give you some background: I’ve been a heavy user of the [Azure Resource Manager (ARM) template](https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-group-overview) s for some time now. We’ve had some good times together. We’ve also had some bad times and boy have we seriously fallen out at times!

I’ve used ARM extensively to automate deployment, configuration and update of the infrastructure in Azure. I’ve used it on everything from single-person projects to large scale dual-region deployments of services with critical live traffic.

The takeaway for me has been: ARM is solid, reliable and does what I need - however, it isn’t always the easiest to use and it forces a wall between infrastructure deployment and application configuration which I’ve always found to be painful.

For example, when working with Azure Storage through ARM you can create a storage account, but you can’t create ‘queues’ or upload ‘blobs’ into the account. Another example - you can create an Azure Kubernetes cluster with AKS but you can’t create ‘secrets’ or ‘deployments’ in the cluster.

Shortcomings like these can be painful at times, and difficult to work around.

Typically, an application running in AKS might use a managed Redis Cache instance or an Azure Storage queue. Ideally, you want to use ARM to…

1. Create your Redis cache
1. Create your storage queue
1. Create your AKS cluster
1. Create secretes in AKS cluster with Redis and Storage access keys
1. Deploy your container based application.

Once ARM has finished, you want your application set and and ready to go. You don’t want the container to have to check and create queues if they don’t exist. You don’t want it to somehow retrieve the Redis access keys.

But, because you can’t set these things up with ARM you’re unable to deploy a working application in one hit. You have to run ARM to create Redis, AKS and storage then after that you have to run a script which injects the access keys into your AKS cluster so your app can use them. That’s not ideal. (Some interesting integration work between ARM and Terraform is underway which provides an interesting way to merge both but I'm not going into that now. [Read more here.](https://azure.microsoft.com/en-us/blog/introducing-the-azure-terraform-resource-provider/))

My other issue with ARM is much more personal and many may disagree, but I find it very hard to author. VSCode has introduced some great ARM tooling which certainly helps make things better, I still find it dense and hard to parse mentally.

{{< figure align=alignnone width=1994 src="/wp-content/uploads/2018/06/armvshcl.png" alt="" >}}

### So, what’s the solution?

Recently, I’ve [been working with Terraform](https://www.terraform.io/) and I think it helps with a lot of these issues. First up, in my opinion, it’s nice and simple to read. Here is an example creating a storage account and uploading a file into it:

{{< figure align=alignnone width=938 src="/wp-content/uploads/2018/06/storageandupload.png" alt="" >}}

Terraform is run on your client, which means you can also do more; things that aren’t possible with ARM, such as uploading ‘blobs’ or configuring ‘queues’.

But that’s not really the main thing that has me converted. Terraform is made up of ‘Providers’ which each know how to talk to a set of API’s. Think of these a bit like ‘Resource Providers’ in ARM. Terraform has [an Azure provider](https://www.terraform.io/docs/providers/azurerm/index.html) but it also has a [Kuberenetes provider](https://www.terraform.io/docs/providers/kubernetes/index.html), [a CloudFlare provider](https://www.terraform.io/docs/providers/cloudflare/index.html), [a RabbitMQ provider](https://www.terraform.io/docs/providers/rabbitmq/index.html)… you get the point.

This is where things really light up. No modern solution of any serious scale is solely confined to one stack/cloud… with Terraform you can reach into all of these by including, or even writing your own, providers. The result is that you have one toolchain and one interface to work with. You can bridge between infrastructure and runtime, taking access keys and injecting them into your k8s cluster as a secret, for example.

You can see an end-to-end example of this [with an AKS cluster in my repo here.](https://github.com/lawrencegripper/azure-aks-terraform)

### “Well that all sounds great Lawrence but what’s the catch?”

Good question. There are, as they’re always are, some downsides to Terraform. Firstly, not all of ARM APIs are provided. There is a nice workaround to this which lets you embed ARM within Terraform which goes a long way to mitigate this issue. You can see an example [of that here where I call Azure Batch from Terraform](https://github.com/lawrencegripper/ion/blob/f52f242dc368cb27f09423aac503eb522e6c3c79/deployment/azurebatch/azurebatch.tf). Some of the docs and examples of Azure are out-of-date or need some tweaking too.  But, it’s an OSS effort so you can (and I have) add functionality for resources that are missing and fix up docs as you go.

Secondly, like any abstraction there are leaks occasionally and where you gain those abstractions it will, ultimately, bite or nip or maybe it won’t, but this depends on what you’re doing and how you’re doing it. For me this hasn’t been a problem… _yet_.

#### What else is in this space?

 [https://pulumi.io/](https://pulumi.io/) also just launched which has an interesting take on the deployment/configuration problem but I haven't yet had a chance to play with it in detail.

Big thanks to [@martinpeck](https://twitter.com/martinpeck) for helping with this post!
