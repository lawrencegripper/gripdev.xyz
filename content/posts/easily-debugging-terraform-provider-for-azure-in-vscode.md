---
author: gripdev
category:
  - how-to
date: "2019-09-12T15:49:31+00:00"
guid: http://blog.gripdev.xyz/?p=1212
title: Easily Debugging Terraform Provider for Azure in VSCode
url: /2019/09/12/easily-debugging-terraform-provider-for-azure-in-vscode/

---
So you're making a change to the provider to add a feature, it's going great and your ready to test it out.... but then you realize things get a bit ropey... ideally you want a visual debugger to step through the code.

Well here is how to set that up in VSCode.

First make sure you have VSCode setup for golang debugging (delve configured etc). Then it's easy, say you want to debug a new provider you've written and it has a test called `TestAccDataSourceAzureRMFunction_basic` in the file `data_source_function_test.go` then you can setup your `launch.json` file in VSCode to look like:

(Makes sure you replace the details in the `private.env` with your service principal and subscription details)

https://gist.github.com/lawrencegripper/94921efed78d6f1514b95b28ae073be6
