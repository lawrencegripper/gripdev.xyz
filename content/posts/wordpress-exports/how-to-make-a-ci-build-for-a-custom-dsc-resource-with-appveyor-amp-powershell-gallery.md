---
author: gripdev
category:
  - how-to
date: "2016-01-21T16:43:21+00:00"
guid: https://gripdev.wordpress.com/?p=813
summary: |-
  A while back I created a custom DSC resource which enabled you to manage packages installed with [Chocolatey](https://chocolatey.org/packages) (awesome package manager for Windows, like apt-get).

  You can do some cool stuff with it, like build out a [dev box in azure](https://gripdev.wordpress.com/2014/12/10/chocolatey-powershell-dsc-and-azure-iaas-automating-dev-box-creation/) with one script, use [for deployment](https://azure.microsoft.com/en-us/documentation/articles/automation-dsc-cd-chocolatey/) or just to setup your machine after a rebuild.

  So it turns out people have been doing just that, it’s had over [1,200 downloads from the PowerShell Gallery](https://www.powershellgallery.com/packages/cChoco/)!

  [![image](http://blog.gripdev.xyz/wp-content/uploads/2016/01/image_thumb.png)](http://blog.gripdev.xyz/wp-content/uploads/2016/01/image.png)

  Not only that but it’s been improved by the [community with pull requests](https://github.com/PowerShellOrg/cChoco/pulls?utf8=%E2%9C%93&q=) adding functionality and improving the resource.

  [![image](http://blog.gripdev.xyz/wp-content/uploads/2016/01/image_thumb1.png)](http://blog.gripdev.xyz/wp-content/uploads/2016/01/image1.png)

  I wasn’t really expecting this to happen so didn’t have a clear idea of how to test, merge and publish these contributions. In the past, [on other projects like powergist](https://github.com/lawrencegripper/PowerGist), [I’ve used Appveyor to handle](https://ci.appveyor.com/) the CI as it’s free for OSS projects and it’s great to use. I thought I’d give it a spin here and see if I could get it to play ball with the PowerShell Resouce.
title: How to make a CI Build for a custom DSC Resource with Appveyor &amp; PowerShell Gallery
url: /2016/01/21/how-to-make-a-ci-build-for-a-custom-dsc-resource-with-appveyor-powershell-gallery/

---
A while back I created a custom DSC resource which enabled you to manage packages installed with [Chocolatey](https://chocolatey.org/packages) (awesome package manager for Windows, like apt-get).

You can do some cool stuff with it, like build out a [dev box in azure](https://gripdev.wordpress.com/2014/12/10/chocolatey-powershell-dsc-and-azure-iaas-automating-dev-box-creation/) with one script, use [for deployment](https://azure.microsoft.com/en-us/documentation/articles/automation-dsc-cd-chocolatey/) or just to setup your machine after a rebuild.

So it turns out people have been doing just that, it’s had over [1,200 downloads from the PowerShell Gallery](https://www.powershellgallery.com/packages/cChoco/)!

[![image](/wp-content/uploads/2016/01/image_thumb.png)](/wp-content/uploads/2016/01/image.png)

Not only that but it’s been improved by the [community with pull requests](https://github.com/PowerShellOrg/cChoco/pulls?utf8=%E2%9C%93&q=) adding functionality and improving the resource.

[![image](/wp-content/uploads/2016/01/image_thumb1.png)](/wp-content/uploads/2016/01/image1.png)

I wasn’t really expecting this to happen so didn’t have a clear idea of how to test, merge and publish these contributions. In the past, [on other projects like powergist](https://github.com/lawrencegripper/PowerGist), [I’ve used Appveyor to handle](https://ci.appveyor.com/) the CI as it’s free for OSS projects and it’s great to use. I thought I’d give it a spin here and see if I could get it to play ball with the PowerShell Resouce.

### What I wanted from my DSC CI Build

1. Do some simple tests to make sure the resources are valid.
1. Increment the manifest file version number.
1. Publish the updated resources to the Powershell Gallery.
1. Checkin the updated manifest file to Git, so we can track the releases back to the source.

Before we get started it’s worth covering some basics of appveyor. Firstly the service lets you store config information and access it via environment variables. For secure variables you can encypt these as well. Secondly it lets you checkin your build definition and have it sit alongside you’re code as an ‘appveyor.yaml’ file. You can [see mine here](https://github.com/PowerShellOrg/cChoco/blob/master/appveyor.yml).

My yaml file defines but can do lots more too:

- The OS build I’d like, we’ve using WMF5 but you might want VS2013 etc.
- Any install scripts to run before building, here we make sure nuget is setup and ready.
- The environment variables, secure and simple.
- Finally the script to run for the build.

With that done we move onto where the meat of the work is going on - the build script. Throughout the description below I've linked to the lines/sections in question, click the links to see the accompanying code.

A the [start of the script](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L52) I pull in the environment variables for my modules name, the folder it’s been cloned too, the nuget key used to publish PSGallery and the Build number. These are all then used later in the script, separating these out makes it easy to chance them or use the script for another DSC resource in the future.

### 1\. Testing the Resource

There is a great package for designing and testing custom resources called ‘xDSCResourceDesigner’ and [a great blog](http://blogs.technet.com/b/privatecloud/archive/2014/05/09/powershell-dsc-blog-series-part-3-testing-dsc-resources.aspx) here on how to use this to run tests.

I adapted the script and added it to my CI Build script, you can see the [chunk here](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L82).

It loops through the resources, tests them and, should any fail, exits with an error code so a broken package isn't pushed to the PSGallery.

### 2\. Increment the version number from the manifest

To do this I read in the [file here](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L121) and invoke is as an expression, it’s a PowerShell hashtable so that's the quickest way to get at it’s content. Once we have it we can then create a new version number, add it in and write the updated file back out to disk.

N.B I couldn’t get the ‘Update-ModuleManifest’ cmdlet to work in appveyor, if I had it would have simplified this process greatly as I could drop the ‘ConvertTo-PSON’ function etc.

### 3\. Publishing to PS Gallery

This is a simple one liner [here](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L154), which takes in the Nuget key and pushes the package up to the gallery.

### 4\. Pushing back to GIT

This is a bit more complex and there is a [great guide here](https://www.appveyor.com/docs/how-to/git-push) by the appveyor guys.

Appveyor, by default, only checks out the commit that triggered the build, not a branch so early on in the [script I checkout the master branch](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L67) to ensure we can push later.

Then [I’ve created an](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L157) oauth token to authorize the build agent with the repo and stored it in Appveyor, as a secure variable. This is then retreived as an environment variable and used to to push the updated file push back to the repo.

N.B Because of the way the git commands write out to stderr I use the [start-process](https://github.com/PowerShellOrg/cChoco/blob/master/AppveyorCIScript.ps1#L166) command to prevent these from registering as build failures. Also, depending on how you work with git and appveyor, you might not want to checkout the master branch and only build a specific commit, however, this simple approach works well for me at the moment.

### Done

I now have a nice build that checks the DSC resource, publishes an update version to PSGallery and commits an updated version number in Github. You can see the build [for yourself here and](https://ci.appveyor.com/project/LawrenceGripper/cchoco) all the code involved, including the yaml and [build script are here.](https://github.com/PowerShellOrg/cChoco)

Ultimately it means I won't be a bottleneck for the project, hopefully allowing more contributions more often and  making for a healthy OSS project.

[![image](/wp-content/uploads/2016/01/image_thumb2.png)](/wp-content/uploads/2016/01/image2.png)
