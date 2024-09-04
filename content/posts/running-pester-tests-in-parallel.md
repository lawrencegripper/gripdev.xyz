---
author: gripdev
category:
  - uncategorized
date: "2020-11-12T17:06:32+00:00"
guid: http://blog.gripdev.xyz/?p=1370
title: Running Pester Tests in Parallel
url: /2020/11/12/running-pester-tests-in-parallel/

---
I'm working on a project which uses [pester tests](https://pester.dev/) to validate our deployment and system health.

We've accumulated a lot of tests. Nearly all of these sit and wait on things to happen in the deployment and we're running them sequentially which takes around 20mins. Each of our tests is in it's own file.

What I wanted to do was, as these are IO bound tests waiting on things to trip in the environment or polling http endpoints, run them all in parallel. One file per thead or something similar.

This issue [discusses this topic](https://github.com/pester/Pester/issues/613). I took a lead from the `Invoke-Parallel` snipped and started playing, unfortunately the output from the tests was mangled and overlapping.

Then I realised I could use [Powershell Jobs](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_jobs?view=powershell-7.1) to schedule the work and poll for the job to be completed then receive each jobs to have to output displayed nicely.

So now the output looks good:

[![](/wp-content/uploads/2020/11/image.png)](/wp-content/uploads/2020/11/image.png)

**Note: We're using Pester v4 you'll have to do a little bit of fiddling to port to Pester 5**  

https://gist.github.com/lawrencegripper/3428970a5be6e1c5e62a13b22e639cd9

Here is the full repro: https://github.com/lawrencegripper/hack-parallelpester
