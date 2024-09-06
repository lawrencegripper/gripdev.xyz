---
author: gripdev
category:
  - uncategorized
date: "2013-12-02T17:44:59+00:00"
guid: http://gripdev.wordpress.com/?p=312
title: Using LinqPad to query Azure Diagnostics (Best WAD viewing client for Devs imo)
url: /2013/12/02/using-linqpad-with-azure-diagnostics-best-wad-viewing-client-for-devs-imo/

---
Hi All,

Really quick and simple post. In a normal windows azure cloud service you get all your lovely eventlogs put into table storage for you to use when diagnosing issues.

Unfortunately that can be a bit daunting as there isn't, or I didn't think there was, a handy client to view them all in.

With the AzureTableStorage driver for LinqPad I can quickly write complex and simple queries to integrate all the WAD logs and I get lovely intellisense in a nice familiar interface.

```
from items in WADWindowsEventLogsTable
 where items.Timestamp > DateTime.Now.AddMinutes(-10)
 where items.Role == "ControllerRole"
 select new { time = items.Timestamp, desc = items.Description}
```

And I'm straight to the root of the problem, or near enough.

[![linqpadAzureWAD](/wp-content/uploads/2013/12/linqpadazurewad.png?w=300)](/wp-content/uploads/2013/12/linqpadazurewad.png)
