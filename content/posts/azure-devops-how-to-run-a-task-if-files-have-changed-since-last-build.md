---
author: gripdev
category:
  - uncategorized
date: "2020-11-11T13:01:18+00:00"
guid: http://blog.gripdev.xyz/?p=1360
title: 'Azure Devops: How to run a Task if files have changed since last build'
url: /2020/11/11/azure-devops-how-to-run-a-task-if-files-have-changed-since-last-build/

---
Shout out to [the awesome work here from Alex Yates!](http://workingwithdevs.com/azure-devops-services-api-powershell-hosted-build-agents/) This post builds on that work and updates a few bits.

What is the the aim? I have a file called `IMAGETAG.txt` which contains a simple version `v1.0.1`. It is used to build and push a Docker container as part of the build. If the file is changed in a commit, I want to build and push the docker image.

Now normally you could use the [Path filtering stuff build into the Azure Devops Triggers](https://docs.microsoft.com/en-us/azure/devops/pipelines/repos/azure-repos-git?view=azure-devops&tabs=yaml#paths) but in this case I have lots of other tasks which I DO want to run and I don't want multiple builds.

So how does it work? Well first up we're going to create a script based off Alex's work and update the params, API versions used and update the way it matches files.

The result is something we can call like this,  

```
      changesSinceLastBuild.ps1`
            -outputVariableName 'ML_IMAGE_TAG_CHANGED'`
            -fileMatchExpression 'containers/IMAGETAG.txt'`
            -branch 'refs/heads/main'`
            -buildDefinitionId '20'

```

This goes and gets the latest successful build from `main` for the specified build definition (in case you have multiple builds) and compares the changes done between then and the current `HEAD` commit.

In then uses powershell `-match` to find out if the file `containers/IMAGETAG.txt` was changed.

If it was it sets the Azure Devops Build variable `ML_IMAGE_TAG_CHANGED` to `true`.

We can then use this as a condition on another task in the Job.

So all together it looks like this, in my case the tasks invoke `psake` targets for `ciml-docker`.

https://gist.github.com/lawrencegripper/6904804849fe609daa9d1d736aa405e8
