---
author: gripdev
category:
  - how-to
date: "2014-05-09T16:57:46+00:00"
guid: http://gripdev.wordpress.com/?p=536
title: Configure Azure Cache From Cloud Config in PAAS
url: /2014/05/09/configure-azure-cache-from-cloud-config-in-paas/

---
I’ve recently been working with the Azure cache  client which, by default, is configured by storing the details in your app.config or web.config.

Now this is great if you’re in the traditional world but for Azure Platform as a service solutions Cloud Configuration is key. It allows you to change settings easily across multiple machines without having to jump into web.configs or redeploy etc.

I wrote the little method below to help me to do this as in my testing I was jumping between a couple of different cache's regularly.

I’ve recently been working with the Azure cache  client which, by default, is configured by storing the details in your app.config or web.config.

Now this is great if you’re in the traditional world but for Azure Platform as a service solutions Cloud Configuration is key. It allows you to change settings easily across multiple machines without having to jump into web.configs or redeploy etc.

I wrote the little method below to help me to do this as in my testing I was jumping between a couple of different cache's regularly.

\[gist https://gist.github.com/2beafe11eef1e024e493/\]

You can then use the CloudConfigurationManager to retreive the info needed then call into the method to start using the cache:

```
CloudConfigurationManager.GetSetting("CacheDiscoryUrl")
```

```
CloudConfigurationManager.GetSetting("Token")
```

You may want to alter to use something other than the "GetDefaultCache" method but hopefully a good starting point.

\[This method was tested in a limited way, double check it makes sense for you before going nuts and putting it into prod!\]

You can then use the CloudConfigurationManager to retreive the info needed then call into the method to start using the cache:

```
CloudConfigurationManager.GetSetting("CacheDiscoryUrl")
```

```
CloudConfigurationManager.GetSetting("Token")
```

You may want to alter to use something other than the "GetDefaultCache" method but hopefully a good starting point.

\[This method was tested in a limited way, double check it makes sense for you before going nuts and putting it into prod!\]
