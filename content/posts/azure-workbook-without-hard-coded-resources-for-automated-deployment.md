---
author: gripdev
category:
  - applicationinsights
  - azureresourcegraph
date: "2021-04-14T11:05:24+00:00"
guid: https://blog.gripdev.xyz/?p=1455
tag:
  - applicationinsights
  - azureresourcegraph
  - workbooks
title: 'Azure: Workbook without hard coded resources for automated deployment'
url: /2021/04/14/azure-workbook-without-hard-coded-resources-for-automated-deployment/

---
**Brain dump post, excuse typos and writing getting this out of my head while I still remember it.**

In this post I'm going to go through how I used query based parameters to setup an Application Insights workbook so it **does not** have hard coded resource ID's in it's definition.

This means it's much easier to use for automated deployments where these ID's aren't known **and** as new resources are deployed the existing workbook picks these up automatically without requiring manual changes.

In this scenario we're deploying a Resource Group with ARM/Terraform. Each group has it's own Application Insights deployed. In the group are some App Service plans, Cosmos DBs and Service Bus Namespaces. We want the workbook to deploy into the Application insights instance in the group and graph the resources for App Service Plans, Cosmos and ServiceBus.

_I'm new to workbooks so be aware there may be a simpler way to do this that I've not yet found!_

First up what does it look like if you don't use this approach and just use the GUI to add metrics to your workbook.

Along the top we use the "Resource Type" and other drop downs to select the resource we want to graph then we add our "CPU" metric.

[![](/wp-content/uploads/2021/04/image.png)](/wp-content/uploads/2021/04/image.png)

If you click into the "Advanced editor" you'll see the following, notice that the "resourceIds" field now has a hard coded reference to the resource we selected.

[![](/wp-content/uploads/2021/04/image-1.png)](/wp-content/uploads/2021/04/image-1.png)

This means if we exported this JSON and deployed it using ARM or Terraform the workbook wouldn't work. We'd want it to graph the metrics for the resource deployed alongside it not the resource that is hard coded.

**So how do we fix this?**

Well we can use workbook parameters.

Parameters can be simple strings or more complex queries and resource selectors.

The first step is to find out the resource group we're deployed into, this can be done by creating a parameter which finds the "Owned Resources".

"Owned Resources" for this Application Insights instance is itself and the query returns it's full Azure ID like: `/subscriptions/YOURSUB/resourceGroups/rg-processing/providers/Microsoft.Insights/components/app-insights`

We're going to use this to extract the current resource group's name.

[![](/wp-content/uploads/2021/04/image-3.png)](/wp-content/uploads/2021/04/image-3.png)

Next we use the an Azure Resource Graph query `where id == "{OwningAppInsights}" | project split(id, "/", 4)[0]` to pull the Resource Group name out of this ID.

The query finds the application insights instance then pulls out the Resource Group it's deployed into (this is what the split is doing on the Azure ID). We add this as a new parameter called "ResourceGroup" notice this param can depend on the previous param "OwningAppInsights" we just created.

[![](/wp-content/uploads/2021/04/image-4.png)](/wp-content/uploads/2021/04/image-4.png)

Now we can create our last workbook parameter, one which selects the App Service Plans in the resource group. This uses the output of the "ResourceGroup" parameter above to query for all the Plans in the group by filter the "type" of the resources in the group.

[![](/wp-content/uploads/2021/04/image-5.png)](/wp-content/uploads/2021/04/image-5.png)

To find out which type you should use in the above query run the following Azure Graph Query and review the results (note turn "formatted results" off to see the original values not the cleaned up ones) `where resourceGroup == "processing-myrg" | project type, name`

So the query `where resourceGroup == "{ResourceGroup}" and type == "microsoft.web/serverfarms"` is returning all the resources that are `servicefarms`... this is internal azure speak for App Service Plans.

We've ticked the box to "Allow multiple selections" and we've ticked "Hide parameter in reading mode" as we don't want users of the workbook to change this manually.

Then we can use this parameter when setting up our metric graph like so, we can select the "ResourceApplicationPlans" parameter from the drop down and the graph now uses our auto-populated set of App Service plans.

[![](/wp-content/uploads/2021/04/image-6.png)](/wp-content/uploads/2021/04/image-6.png)

**Now we're they're** the code/json of the workbook no longer contains any hard coded references to ID's

[![](/wp-content/uploads/2021/04/image-7.png?w=661)](/wp-content/uploads/2021/04/image-7.png)

You can see the "ResourceIds" is now set by out "ResourceApplicationPlans" parameter which is dynamically generated and selects all the App Plans that are deployed in the resource group the workbook is deployed in.

We can now automate the deployment of the workbook without templating the json!

**Bonus** if you add a new App Plan the workbook will pick it up and start graphing it. You can use the same approach to add parameters detecting other resource types like cosmos and graph those too.

[![](/wp-content/uploads/2021/04/image-8.png?w=879)](/wp-content/uploads/2021/04/image-8.png)
