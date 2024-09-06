---
author: gripdev
category:
  - opa
date: "2020-01-13T20:20:40+00:00"
guid: http://blog.gripdev.xyz/?p=1241
summary: |-
  First up, quick refresher - what is a mutating admission controller?

  Well it's a nice feature in Kubernetes which lets you intercept objects when they're created and make changes to them before they are deployed into the cluster.

  Cool right? All those fiddly bits of YAML or hard to enforce company policies around network access, image stores you can and can't use, they can all be enforced and FIXED automagically! (Like all magic caution is advised, choose wisely - queue Monty python gif)

  ![giphy](https://blog.gripdev.xyz/wp-content/uploads/2020/01/giphy.gif)

  So what's the catch? Well without Open Policy Agent (OPA) you had to build out a web api to do the magic of changing the object then build/push an image and go through maintaining the solution. While you can write them quite easily now with solutions like KubeBuilder, [or if you really love node I build one using that too](https://github.com/lawrencegripper/MutatingAdmissionsController), I wanted to see if OPA made things easier.

  So say you want something more dynamic, flexible and a little easier to look after?

  [This is where Open Policy Agent comes in, they have a DSL language specially designed to build out and enforce complex policies.](https://www.openpolicyagent.org/docs/latest/kubernetes-introduction/)

  Today I've been having a play with it to work out if I could build a controller which would set a certain `nodeSelector` on pods based on which `namespace` they are deployed in.

  [I'll go over this very broadly I highly recommend looking at the docs in detail before diving in](https://www.openpolicyagent.org/docs/latest/), I lost quite a bit of time to not reading things properly before starting.

  I won't lie, getting used to the DSL ( `rego`) was painful for me, mainly because I came at it thinking it was going to be really like Golang. It does look quite like it but that's where the similarity ends, it's more functional/pattern matching and better suited to tersely making decisions based on data.

  To counter the learning curve of `rego` I have to say, as I've raised issues and contributions the maintainers have been super responsive and helpful (even when I've made some silly mistakes) and the docs are great with runnable samples to get started.

  Lets talk more about what I built out.
tag:
  - opa
  - openpolicyagent
  - rego
title: Mutating Admissions Controllers with Open Policy Agent and Rego
url: /2020/01/13/mutating-admissions-controllers-with-open-policy-agent-and-rego/

---
First up, quick refresher - what is a mutating admission controller?

Well it's a nice feature in Kubernetes which lets you intercept objects when they're created and make changes to them before they are deployed into the cluster.

Cool right? All those fiddly bits of YAML or hard to enforce company policies around network access, image stores you can and can't use, they can all be enforced and FIXED automagically! (Like all magic caution is advised, choose wisely - queue Monty python gif)

![giphy](/wp-content/uploads/2020/01/giphy.gif)

So what's the catch? Well without Open Policy Agent (OPA) you had to build out a web api to do the magic of changing the object then build/push an image and go through maintaining the solution. While you can write them quite easily now with solutions like KubeBuilder, [or if you really love node I build one using that too](https://github.com/lawrencegripper/MutatingAdmissionsController), I wanted to see if OPA made things easier.

So say you want something more dynamic, flexible and a little easier to look after?

[This is where Open Policy Agent comes in, they have a DSL language specially designed to build out and enforce complex policies.](https://www.openpolicyagent.org/docs/latest/kubernetes-introduction/)

Today I've been having a play with it to work out if I could build a controller which would set a certain `nodeSelector` on pods based on which `namespace` they are deployed in.

[I'll go over this very broadly I highly recommend looking at the docs in detail before diving in](https://www.openpolicyagent.org/docs/latest/), I lost quite a bit of time to not reading things properly before starting.

I won't lie, getting used to the DSL ( `rego`) was painful for me, mainly because I came at it thinking it was going to be really like Golang. It does look quite like it but that's where the similarity ends, it's more functional/pattern matching and better suited to tersely making decisions based on data.

To counter the learning curve of `rego` I have to say, as I've raised issues and contributions the maintainers have been super responsive and helpful (even when I've made some silly mistakes) and the docs are great with runnable samples to get started.

Lets talk more about what I built out.

 **Health warning:** I'm new to this and still learning, I may get some of this wrong. Over the next few days I'll be doing more testing and loop back to fix things up.

First up you need to process the input from the request `input` and output a `main` object which will be the response sent to the K8s API.

The first `response` on line 10 is the default which is returned if nothing else takes over.

The second `response` is a mix between and object definition and a set of `rules` lines 18->25 (think assertions). If the `rules` all match then the assignment takes place and `response` becomes the `output` object on line 49. After the assertions there is some plumbing which builds up the `response` with the JSON patches K8s is looking, these set the `nodeselector`

https://gist.github.com/lawrencegripper/b2603df6267334e6752d25d2eb2eb228

So how do these rules work then? Well they're like little functions that check certain things.

In `isPod` we do a simple equality check on the `kind` of the request.

In `hasNodeSelector` we check if the pod already has a node selector by first checking if the field exists, if it doesn't the next check doesn't happen, then checking how many items are in it.

`getPoolForNamespace` is a special case it takes an input `namespace` then loops through the array defined as `namespaceToAgentPool` and sees if any pools match the namespace of the pod. The magic happens here with the `namespaceToAgentPool[_]` which means, roughly, check the rules against all of the items in that array.

https://gist.github.com/lawrencegripper/7b796d5eb691e9bd1e8e23cb93ba8155

So how do you know this stuff works? Well it's got a nice testing framework you can use to check things are working how you expect. I found the `trace` command super useful when running `opa test *.rego -v --explain full` as it would prove the value of the items passed to it along with other information about the execution.

https://gist.github.com/lawrencegripper/3d1a2822dd57d000237d0b5153977c7f

So all together now this looks like: https://github.com/open-policy-agent/contrib/pull/93

Hopefully this is useful. For more advanced stuff there is a library of shared helpers that can be pulled into `rego` which is well worth a look: https://github.com/open-policy-agent/library/tree/master/kubernetes/mutating-admission nearly everything I've done here is largely based on simplifying those funcs and adding more comments.
