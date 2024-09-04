---
author: gripdev
category:
  - c#
  - xaml
date: "2014-03-20T20:15:01+00:00"
guid: http://gripdev.wordpress.com/?p=493
tag:
  - c#
  - xaml
title: Windows Store App - Infinite Scrolling list with Xaml
url: /2014/03/20/windows-store-app-infinite-scrolling-list-with-xaml/

---
I recently came across a great interface and behavior in windows 8.1. If you implement ISupportIncrementalLoading on your bound property then bind up to a gridview you'll get infinite scrolling for free.

Below is an example from my HypeMix app where I pull in more tracks and add more to that list as the user scrolls, hopefully useful to others.

\[gist https://gist.github.com/lawrencegripper/9396993\]
