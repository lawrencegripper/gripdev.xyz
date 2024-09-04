---
author: gripdev
category:
  - how-to
date: "2014-08-15T12:00:14+00:00"
guid: http://gripdev.wordpress.com/?p=579
title: Integration testing MongoDB with C# - More Fluent Syntax
url: /2014/08/15/integration-testing-mongodb-c-sharp-more-fluent-syntax/

---
Hi,

This is similar to my post around Azure Storage Integration testing just with a mongo twist. I've been working with mongo and didn't really like the debugging workflow: Do something, start mongovue and manually check the item changed in the way you thought it would.

I've written a set of extensions to allow a nice set of declarative fluent assertions about the state of the MongoDB after an action has been performed.

\[sourcecode language="csharp"\]
 //assert
 mongoServer.AssertDatabaseExists(databaseName)
 .AssertCollectionExists<ExampleType>()
 .AssertCollectionItemCount(1)
 .AssertCollectionHasItemWithProperty<ExampleType>(item, x => x.MyProperty == 1);
\[/sourcecode\]

Alongside this I've used an awesome [nuget package](https://www.nuget.org/packages/MongoDB.Embedded/) from the guys at [FireFunnel](https://github.com/funnelfire/MongoDB.Embedded), which lets me spin up and work with an embedded Mongodb instance. I've written some simple methods to spin up the embedded instance and clean it in-between tests like so:

\[sourcecode language="csharp"\]
 \[TestClass\]
 public class ExampleTestClass
 {
 private static EmbeddedMongoDbServer mongoEmbedded;
 private static MongoClient mongoClient;
 private static MongoServer mongoServer;

 \[ClassInitialize\]
 public static void StartMongoEmbedded(TestContext cont)
 {
 mongoEmbedded = new EmbeddedMongoDbServer();
 mongoClient = mongoEmbedded.Client;
 mongoServer = mongoClient.GetServer();
 }

 \[ClassCleanup\]
 public static void ShutdownMongo()
 {
 mongoEmbedded.Dispose();
 }

 \[TestInitialize\]
 public void CleanMongo()
 {
 var databases = mongoServer.GetDatabaseNames();
 foreach (var databaseName in databases)
 {
 mongoServer.DropDatabase(databaseName);
 }
 }

...
\[/sourcecode\]

Then I've also tweaked the extensions methods to give some nice actionable feedback when a test fails. For example when you create a query this is captured as an expression that is shown in the test output so you know exactly what query ran and can go about fixing it.

The source is up on Github, [here](https://github.com/lawrencegripper/FluentMongoIntegrationTesting). Let me know how you get on!

[![fluentMongo](/wp-content/uploads/2014/08/fluentmongo.png?w=300)](/wp-content/uploads/2014/08/fluentmongo.png)[![FluentMongoOutput](/wp-content/uploads/2014/08/fluentmongooutput.png?w=300)](/wp-content/uploads/2014/08/fluentmongooutput.png)
