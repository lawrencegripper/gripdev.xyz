---
author: gripdev
category:
  - uncategorized
date: "2016-04-08T08:18:17+00:00"
guid: https://gripdev.wordpress.com/?p=849
title: Backup Azure Table Storage - Quick PowerShell Script
url: /2016/04/08/backup-azure-table-storage-quick-powershell-script/

---
Hi,

So I recently had to make some changes and wanted to ensure I could roll back, in the event of any issues.

Azure Storage is great at replication and availability but it doesn't offer point in time restore, should you accidentally make some changes and want to roll back. So I wrote a quick little powershell script which uses AzCopy and the Table Storage Rest API to find all the tables in an account and then pull down a JSON file of their contents, tagged with the current time.

It uses a SAS token to limit the permissions the script has, only needing list & read, to minimize the risk if the script has an issue.

{{< gist lawrencegripper ed4419bbbdb89534970ab6bba8265a81 >}}

[![TableBackupPic](/wp-content/uploads/2016/04/tablebackuppic.png)](/wp-content/uploads/2016/04/tablebackuppic.png)

The only requirement for the script is [AzCopy](https://azure.microsoft.com/en-gb/documentation/articles/storage-use-azcopy/), an MS utility which it uses to pull down the content of each table.

(Here is a quick overview of [SASTokens in Azure Storage](https://azure.microsoft.com/en-gb/documentation/articles/storage-dotnet-shared-access-signature-part-1/), the [Microsoft Storage Explore](http://storageexplorer.com/) r offers a nice right click option to generate them - make sure you give the SASToken Permissions to access the Table Service, Read and List)

Hope it's useful!
