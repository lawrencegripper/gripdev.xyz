---
author: gripdev
category:
  - uncategorized
date: "2021-02-10T12:27:20+00:00"
guid: http://blog.gripdev.xyz/?p=1382
title: Add Diagnostic Settings to Azure Storage account via Terraform
url: /2021/02/10/add-diagnostic-settings-to-azure-storage-account-via-terraform/

---
So you want to add a diagnostic setting to your Azure storage account via Terraform and you pass the storage account ID to `target_resource_id` only to get the following error:

> Status=400 Code="BadRequest" Message="Category 'StorageWrite' is not supported."

Here is the fix, the diagnostic target resource actually needs to be a sub-resource of the storage account, the ID for that is constructed as:

> <StorageAccountId>/blobServices/default/

Using this you can create a terraform file similar to the below and it will create the diagnostic setting for you on the blob account.

For a more [detailed explanation the issue here goes into more detail.](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8275#issuecomment-755222989)

https://gist.github.com/lawrencegripper/b05872a6d3d72641c7a276c8c1357d53
