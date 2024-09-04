---
author: gripdev
category:
  - '#terraform'
  - azure
  - azurefunctions
  - devops
date: "2019-07-16T14:56:50+00:00"
guid: http://blog.gripdev.xyz/?p=1195
tag:
  - '#terraform'
  - azure
  - azurefunctions
  - devops
title: 'Terraform: Get Azure Function key'
url: /2019/07/16/terraform-get-azure-function-key/

---
**Update 12/11/2020: This is now supported directly in the Azure [Terraform Provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/function_app_host_keys)** see here.

**Updated 09/03/2020: This new method in the Azure provider has intermittent issues. [I have another workaround here which avoids ARM templates as an alternative.](/2021/03/09/azure-functions-get-key-from-terraform-without-internalservererror/)**

So you've deployed your function and you want to get pass the secure `url` another component in your deployment so it can use it...

[Well currently there isn't an output item on the `azurerm_function_app`](https://github.com/terraform-providers/terraform-provider-azurerm/issues/699) resource in Terraform (I'm hoping to fix that up if I get some time) so how do you do it?

Here is a my quick and dirty fix using the `azure_template_deployment` resource in Terraform.

We create an empty release and then use the `listkeys` function to pull back the keys for the function. We only want the function key so we index into the object with `functionKeys.default` (you can get the `master` key too if you want).

Then we output this from the Terraform so it can be used elsewhere. You can now go ahead and pass this into your other component.

https://gist.github.com/lawrencegripper/48cb08fe1c7952f1caf2a18828c6a357
