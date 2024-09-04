---
author: gripdev
category:
  - api-management
  - azure
  - powershell
date: "2015-07-10T16:08:56+00:00"
guid: https://gripdev.wordpress.com/?p=746
tag:
  - api-management
  - azure
  - powershell
title: 'Azure API Management: PowerShell Reporting &amp; Cache Hit Ratio'
url: /2015/07/10/azure-api-management-powershell-reporting-cache-hit-ratio/

---
I was recently working with a customer who needed to view the cache hit ratio of an Azure API Management instance. Currently the UI dashboard doesn't report this information, however there is a lovely API which you can call.

As we needed something up and running fast PowerShell seemed like the perfect way to quickly retrieve the JSON and convert it into an object then format it nicely. The script also stores the results to disk so we have a copy to hand.

[![Output](/wp-content/uploads/2015/06/output.png?w=500)](/wp-content/uploads/2015/06/output.png)

Here is the resulting little script.

\[gist https://gist.github.com/lawrencegripper/c92c245f20050381b7dd\]

To use this on your instance replace the URL and SharedSigniture with your own. You can get these from the administration portal for APIM under the security tab.

For full details on what you can do with the rest API head here for the docs. https://msdn.microsoft.com/en-us/library/azure/dn776326.aspx
