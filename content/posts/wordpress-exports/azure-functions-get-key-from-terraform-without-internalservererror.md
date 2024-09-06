---
author: gripdev
category:
  - azurefunctions
  - terraform
date: "2021-03-09T09:55:27+00:00"
guid: https://blog.gripdev.xyz/?p=1393
tag:
  - azurefunctions
  - terraform
title: Azure Functions Get Key from Terraform without InternalServerError
url: /2021/03/09/azure-functions-get-key-from-terraform-without-internalservererror/

---
So you're trying to use the Terraform `azurerm_function_app_host_keys` resource to get the keys from an Azure function after deployment. Sadly, as of 03/2021, this can fail intermittently 😢 (See [issue 1](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9854) [and 2](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9869)).+

**\[Edit: Hopefully this [issue is resolved by this PR once released](/2021/03/09/azure-functions-get-key-from-terraform-without-internalservererror/) so worth reviewing once the change is released\]**

These errors can look something like these below:

_Error making Read request on AzureRM Function App Hostkeys "\*\*\*": web.AppsClient#ListHostKeys: Failure responding to request: StatusCode=400 -- Original Error: autorest/azure: Service returned an error. Status=400 Code="BadRequest" Message="Encountered an error (ServiceUnavailable) from host runtime"_

_Error: Error making Read request on AzureRM Function App Hostkeys "somefunx": web.AppsClient#ListHostKeys: Failure responding to request: StatusCode=400_

You can [work around this by using my previous workaround with ARM templates but it's a bit clunky](/2019/07/16/terraform-get-azure-function-key/) so I was looking at another way to do it.

There is an [AWESOME project by Scott Winkler called Shell Provider,](https://github.com/scottwinkler/terraform-provider-shell) it lets you write a custom Terraform provider using scripts. You can implement data types and full resources with CRUD support.

Looking into the errors returned by the azurerm\_function\_app\_host\_keys resource they're intermittent and look like they're related to a timing issue. Did you know the `curl` command support retrying out of the box?

https://twitter.com/lawrencegripper/status/1369008407127224322

So using the Shell provider we can create a simple script to make the REST request to the Azure API and use `curls` inbuilt retry support to have the request retried with an exponential back-off until it succeeds or 5mins is up!

**Warning: This script uses --retry-all-errors which is only available in** **v7.71 and above**. The version shipped with the distro your using might not be up-to-date user `curl --version` to check.

Here is a rough example of what you end up with:

https://gist.github.com/lawrencegripper/2d68c369ba48667583df1538c4276026
