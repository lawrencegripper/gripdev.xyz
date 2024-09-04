---
author: gripdev
category:
  - how-to
date: "2017-01-27T15:30:41+00:00"
guid: https://gripdev.wordpress.com/?p=875
title: Fixing ASPNET Production Issues by adding custom data to App Insights logs
url: /2017/01/27/fixing-aspnet-production-issues-by-adding-custom-data-to-app-insights-logs/

---
Debugging issues in production is hard. Things vary - seemingly identical requests to the same URL could be made but one succeed and the other fail.

Without information about the context and configuration it's hard to isolate issue, here is a quick way to get more of that context.

The key is having the data to understand the cause of the failures. One of the things we do to help that is creating our own ITelemetryInitializer for application insights.

In this we track which environment the request was made in, the cloud instance handling the request, code version, UserId and lots more.

This means, if an issue occurs, we have a wealth of data to understand and debug the issue. It's embedded in every telemetry event tracked.

Here is an example of a Telemetry Initializer which adds the UserID and Azure WebApps Instance Id to tracked events, if a user is signed in. You'll find, once you start using this you'll add more information over time.

{{< gist lawrencegripper a59040d043dfafc483a8896c662f4f06 >}}

To get it setup register it, like this, in your Startup.cs:

TelemetryConfiguration.Active.TelemetryInitializers.Add(newAppInsightsTelemetryInitializer());

Go do it now, seriously you'll need it later!
