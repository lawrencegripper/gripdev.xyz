---
author: gripdev
date: "2020-03-30T20:08:16+00:00"
guid: http://blog.gripdev.xyz/?p=1293
tag:
  - '#terraform'
  - azure
  - cluster
  - databricks
title: 'Azure Databricks and Terraform: Create a Cluster and PAT Token'
url: /2020/03/30/azure-databricks-and-terraform-create-a-cluster-and-pat-token/

---
**\[Update Feb 2021**\] There is now a Terraform Provider for Databricks, it’s a better route - https://registry.terraform.io/providers/databrickslabs/databricks/latest/docs

My starting point for a recent bit of work was to try and reliably and simply deploy and manage Databricks clusters in Azure. Terraform was already in use so I set about trying to see how I could use that to also manage Databricks.

I had a look around and after trying the Terraform REST provider and a third party Datbricks provider (didn't have much luck with either) found a Terraform Shell provider. This turned out to be exactly what I needed.

If you haven't written a Terraform provider here's a crash course. You basically just define a method for `create`, `read`, `update` and `delete` and the parameters they take. Then Terraform does the rest.

The Shell provider ( [https://github.com/scottwinkler/terraform-provider-she](https://github.com/scottwinkler/terraform-provider-she) ll) lets you do this by passing in scripts (bash, powershell, any executable that can take stdin and output stdout). In this case I wrote some `powershell` to wrap the `databricks-cli`.

It's _better_ (or different) to `localexec` with `nullresources` as you can store information in the Terraform State and detect drift. If a `read` returns different information than the current information in the state then `update` will be called, for example.

So I took the work of [Alexandre and wrapped it into this provider](https://cloudarchitected.com/2020/01/provisioning-azure-databricks-and-pat-tokens-with-terraform/) and using the Shell provider have a simple, no frills Databricks provider for Terraform which makes calls to Databricks via the `databricks-cli`.

This is currently a simple hack and hasn't undergone any significant testing: [https://github.com/lawrencegripper/hack-databricksterraform](https://github.com/lawrencegripper/hack-databricksterraform). The flow is as follows:

![](/wp-content/uploads/2020/03/image.png?w=399)

Hopefully this might be useful to others as a starting point for others.
