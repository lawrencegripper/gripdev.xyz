---
author: gripdev
category:
  - how-to
date: "2014-07-22T12:23:22+00:00"
guid: http://gripdev.wordpress.com/?p=567
title: Integration testing azure storage - Fluent syntax
url: /2014/07/22/integration-testing-azure-storage-fluent-syntax/

---
Hi,

**\[Update 03/21: Azureite local emulator is a good one to look at in this space too. [More details here.](_wp_link_placeholder)\]**

I recently set about writing a solution that's heavily reliant on Azure Blob storage. I found my debugging cycle wasn't nice, I'd spin up the code then spend ages in Azure Storage Explorer to work out what had happened. I also new I'd want some integration tests for the future.

So I did some research set about writing a quick set of helpers to allow me to write clean, quick and simple tests that worked in VS Test explorer.

The result is [FluentAzureBlobTesting](https://github.com/lawrencegripper/FluentAzureBlobTesting "FluentAzureBlobTesting"), this allows me to write lovely declarative statements like this:

```
[sourcecode language="csharp"]
           blobClient
                .AssertContainerExists(expectedContainerName)
                .AssertBlobExists(expectedBlobName)
                .AssertBlobDataIs(expectedBlobData)
                .AssertBlobContainsMetaData(expectedMetaDataKey, expectedMetaDataValue)[/sourcecode]
```

It also handles the starting, stopping and clearing of the storage emulator so all you have to do is click "Run All" and it handles the rest. \*Dependency on Azure 2.3 SDK

```
[sourcecode language="csharp"]
        private static CloudStorageAccount account;
        private static CloudBlobClient blobClient;

        [ClassInitialize]
        public static void StartAndCleanStorage(TestContext cont)
        {
            account = CloudStorageAccount.DevelopmentStorageAccount;
            blobClient = account.CreateCloudBlobClient();
            blobClient.StartEmulator();
        }

        [ClassCleanup]
        public static void ShutdownStorage()
        {
            blobClient.StopEmulator();
        }

        [TestInitialize]
        public void CleanAndRestartStorage()
        {
            blobClient.ClearBlobItemsFromEmulator();
        }
[/sourcecode]
```

Should the test fail the extensions report back the reason for failures and log trace output on success.

[![FailedTest](/wp-content/uploads/2014/07/failedtest.png)](/wp-content/uploads/2014/07/failedtest.png)[![PassingTestWithTrace](/wp-content/uploads/2014/07/passingtestwithtrace.png?w=300)](/wp-content/uploads/2014/07/passingtestwithtrace.png)

Now my debugging cycle is nice a quick and I'm writing a good set of unit tests as I go.

The source is available on github here, hopefully useful to you!

https://github.com/lawrencegripper/FluentAzureBlobTesting.

Thanks to Rory for the starting point here: http://www.neovolve.com/post/2012/01/12/Integration-testing-with-Azure-development-storage.aspx
