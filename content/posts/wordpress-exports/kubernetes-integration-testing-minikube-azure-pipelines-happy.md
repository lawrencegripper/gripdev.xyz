---
author: gripdev
category:
  - automation
  - azurepipelines
  - ci
  - devops
  - golang
  - minikube
date: "2018-10-19T09:35:32+00:00"
guid: http://blog.gripdev.xyz/?p=1166
summary: |-
  Update: With the release of KIND (Kubernetes in Docker) I've now moved to using this over minikube as it's quicker and simpler.

  I recently did some work on a fairly simple controller to run inside Kubernetes. It connects to the K8s API and watches for changes to `ingress` objects in the cluster.

  I had a nice cluster spun up for testing which I could tweak and poke then observe the results. This was nice BUT I wanted to translate it into something that ran as part of my CI process to make it more repeatable. Having not played much with the new [Azure Pipelines I decided to try and get this working using one.](https://azure.microsoft.com/en-us/services/devops/pipelines/)

  Here was the goal:

  - Build the source for the controller
  - Spin up a Kuberentes cluster
  - Deploy test resources (Ingress and Services) into the cluster
  - Connect the controller code to the cluster and run it's tests

  The obvious choice was to look at creating the clusters inside a cloud provider and using it for testing **but** I wanted each PR/Branch to be validated independently in a separate cluster, ideally in parallel, so things get complicated and expensive if we go down that route.

  Instead I worked with [MiniKube which has a 'no vm mode'](https://github.com/kubernetes/minikube#linux-continuous-integration-without-vm-support), this spins up a whole cluster using just docker containers. The theory was, if the CI supports running docker containers it should support MiniKube clusters...

  **TLDR:** Yes this is possible with MiniKube and Azure Pipelines or Travis CI - Skip to the end to see how.
tag:
  - automation
  - azurepipelines
  - ci
  - devops
  - golang
  - minikube
  - testing
  - travis
  - ubuntu
  - yaml
title: 'Kubernetes Integration Testing: MiniKube + Azure Pipelines = Happy'
url: /2018/10/19/kubernetes-integration-testing-minikube-azure-pipelines-happy/

---
Update: With the release of KIND (Kubernetes in Docker) I've now moved to using this over minikube as it's quicker and simpler.

I recently did some work on a fairly simple controller to run inside Kubernetes. It connects to the K8s API and watches for changes to `ingress` objects in the cluster.

I had a nice cluster spun up for testing which I could tweak and poke then observe the results. This was nice BUT I wanted to translate it into something that ran as part of my CI process to make it more repeatable. Having not played much with the new [Azure Pipelines I decided to try and get this working using one.](https://azure.microsoft.com/en-us/services/devops/pipelines/)

Here was the goal:

- Build the source for the controller
- Spin up a Kuberentes cluster
- Deploy test resources (Ingress and Services) into the cluster
- Connect the controller code to the cluster and run it's tests

The obvious choice was to look at creating the clusters inside a cloud provider and using it for testing **but** I wanted each PR/Branch to be validated independently in a separate cluster, ideally in parallel, so things get complicated and expensive if we go down that route.

Instead I worked with [MiniKube which has a 'no vm mode'](https://github.com/kubernetes/minikube#linux-continuous-integration-without-vm-support), this spins up a whole cluster using just docker containers. The theory was, if the CI supports running docker containers it should support MiniKube clusters...

**TLDR:** Yes this is possible with MiniKube and Azure Pipelines or Travis CI - Skip to the end to see how.

Azure Pipelines offer 'Ubuntu 16.04' as a base for builds so I set out building a script that would work against that. Luckily there is some prior work [by the travis team which got me started.](https://blog.travis-ci.com/2017-10-26-running-kubernetes-on-travis-ci-with-minikubehttps://blog.travis-ci.com/2017-10-26-running-kubernetes-on-travis-ci-with-minikube)

I reworked their `.travis.yaml` into a script file which could be used against, in theory, any Ubuntu 16 image. There are a few notable tweaks that I had to do here from the original Travis example:

1. For some reason the file permissions for the '.kube' and '.minikube' folders misbehaved in Azure Pipelines so this is fixed up on like #18 and #19
1. I pinned the version numbers of both 'kubectl' and 'minikube' on #12 and #14 to prevent the script breaking as changes are made to either tool. (Previously these took 'latest')
1. As miniKube clusters are run locally they don't understand how to deal with a 'Service' with 'Type=LoadBalancer'. On line #36 I include a workaround for this by [elsonrodriguez](https://github.com/elsonrodriguez/minikube-lb-patch). This means I can test the same YAML I'll be using in my real clusters, rather than having a separate YAML for MiniKube and Production.

https://gist.github.com/lawrencegripper/1ff51414ec77a907d4f0ec647e846d59

The second script loops through all the YAML files in the 'testyaml' folder and deploys them to the newly created cluster with 'kubectl'. To test the controller against different setups [I create several namespaces, one for each test case, and deploy test resources](https://github.com/lawrencegripper/azurefrontdooringress/tree/3a2c308eeb972dc63081ce7ba874cc4cb5c98cc1/scripts/testyaml) into the namespace. My tests then pick different namespaces and assert the behavior is correct. You c [an see this in the following code](https://github.com/lawrencegripper/azurefrontdooringress/blob/3a2c308eeb972dc63081ce7ba874cc4cb5c98cc1/controller/controller_test.go), the 'name' parameter is the namespace which each of the tests will run against.

https://gist.github.com/lawrencegripper/deff8dd8844fc397fe4b5eab5f91b46a

Last but not least I need to run these scripts inside Azure Pipelines. I'm a big fan of 'Configuration as Code' so I used the YAML definition files rather than the UI editor. Quite a bit of this file is Golang specific build configuration, if your not using Go then all you'll need is the `pool.vmImage` definition on #1-2 and to invoke the script we created earlier, I'm doing this on line #25 with 'bash -f ./scripts/startminikube\_ci.sh' then starting my integration tests with `make integration` (you can replace the `make` call with your tests).

https://gist.github.com/lawrencegripper/deff8dd8844fc397fe4b5eab5f91b46a

To see how this all works together have a look at the [repository hosting my controller code here.](https://github.com/lawrencegripper/azurefrontdooringress/tree/b157b248400d7180951a985f15245ac46877506c)

If you want to build a docker image and then deploy it inside the [MiniKube cluster there appears to be a way to do this too](https://stackoverflow.com/a/48999680). I haven't tried it but it would remove the need for a build to push the image to a repository before testing. If the test passes then the build can push and tag the image for others to use.

I'm pretty happy with the results and enjoyed working with Azure Pipelines for the first time.

\[gallery ids="1174,1172,1171" type="rectangular"\]As a point of interest I also got the same setup working with TravisCI (which has been my go-to CI for OSS projects in the past) to compare the two. Apart from a slightly different amount of time to start MiniKube (AzurePipelines: 3.38mins vs Travis 5.10mins) there was very little difference between them. The only one I noticed is that [the Travis.yaml file](https://github.com/lawrencegripper/azurefrontdooringress/blob/b157b248400d7180951a985f15245ac46877506c/.travis.yml) may be a little more readable but this is pretty subjective. One this that I'm not making use of yet in Azure Pipelines, but does set them apart, is the ' [Release Management](https://docs.microsoft.com/en-us/azure/devops/pipelines/release/what-is-release-management?view=vsts)' you can tag on after a CI build.
