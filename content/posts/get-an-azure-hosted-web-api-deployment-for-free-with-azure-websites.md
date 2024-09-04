---
author: gripdev
category:
  - how-to
date: "2013-03-19T17:07:02+00:00"
guid: http://gripdev.wordpress.com/?p=282
title: Get an Azure hosted Web API deployment for free with Azure Websites
url: /2013/03/19/get-an-azure-hosted-web-api-deployment-for-free-with-azure-websites/

---
Hi all,

I was looking at doing a rough draft (think POC) for an app and wanted to use WebAPI for the back end. Azure hosting came to mind so I had a play creating a cloud service. Unfortunately my azure bill, mainly from the beeb live tile's using up bandwidth, is already more than enough.

This was when I realised that you can host MVC 4 projects in Azure Websites and the shared instances allow 10 free sites. This means I can quickly setup a simple Web API, that isn't going to see huge traffic or require gigantic resources, for free and use this to build out the app!

\*Quick note – Azure Websites are in preview so the usual caveats apply\*

Another cool feature that I'll cover in another post is that you can set it up to publish from TFS by integrating with tfs.visualstudio.com to get a completed continuous integration solution.

Simple Guide:

Detailed guide - http://www.windowsazure.com/en-us/develop/net/tutorials/get-started/

1. Setup a new shared website through the azure portal.

![](/wp-content/uploads/2013/03/031913_1706_getanazureh1.png)

1. Download the publishing profile and save it to you disk

![](/wp-content/uploads/2013/03/031913_1706_getanazureh2.png)

1. Create your WebAPI Project in Visual Studio

![](/wp-content/uploads/2013/03/031913_1706_getanazureh3.png)

1. Add in some code, this is the default controller you'll get with the project.
   ![](/wp-content/uploads/2013/03/031913_1706_getanazureh4.png)
1. Right click the project and select publish then import the publishing profile you downloaded.

![](/wp-content/uploads/2013/03/031913_1706_getanazureh5.png)

1. Click publish and your site it live, for free and hosting a nice Web API.
1. Just to prove it the values controller returns as expected.

![](/wp-content/uploads/2013/03/031913_1706_getanazureh6.png)

More info - http://www.windowsazure.com/en-us/home/scenarios/web-sites/
