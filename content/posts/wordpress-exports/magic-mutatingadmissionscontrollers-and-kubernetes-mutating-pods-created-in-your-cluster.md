---
author: gripdev
category:
  - '#kubernetes-#mutatingadmissionscontroller-#webhook'
  - kubernetes
date: "2018-08-16T10:52:07+00:00"
guid: http://blog.gripdev.xyz/?p=1155
tag:
  - '#kubernetes-#mutatingadmissionscontroller-#webhook'
title: 'Magic, MutatingAdmissionsControllers and Kubernetes: Mutating pods created in your cluster'
url: /2018/08/16/magic-mutatingadmissionscontrollers-and-kubernetes-mutating-pods-created-in-your-cluster/

---
I recently wanted to use a Mutating Admissions Controller in Kubernetes to alter pods submitted to the cluster - here is a quick summary of how to do it.

In this case we wanted to change the image pull location, just as a quick example (I'm not sure this is a great idea in a real system as it introduces a single point of failure for pod creation but the sample code should be useful to others).

So how do they work? Well it's super simple, you register a webhook in K8s which is called when a certain action occurs and you create a receiver which accepts that webhook and responds with a JSONPatch containing any changes you want to make.

Lets try it out, first up you'll need ngrok, this creates a public endpoint for a port on your machine with an https cert and everything. We'll use this for testing.

Lets start our webhook receiver locally.

1. . `ngrok http 3000`
1. `​git clone https://github.com/lawrencegripper/MutatingAdmissionsController` and CD into the dir
1. `npm install && npm watch-server`

Well you register a webhook in kubernetes which is called when certain things happen, in this case we register one to be called when a pod is created:

```
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: local-repository-controller-webhook
webhooks:
- clientConfig:
    # ngrok public cabundle
    caBundle: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlGcXpDQ0JKT2dBd0lCQWdJUURITW1IVnRRYmdwZGRocWhDUllpbVRBTkJna3Foa2lHOXcwQkFRc0ZBREJlDQpNUXN3Q1FZRFZRUUdFd0pWVXpFVk1CTUdBMVVFQ2hNTVJHbG5hVU5sY25RZ1NXNWpNUmt3RndZRFZRUUxFeEIzDQpkM2N1WkdsbmFXTmxjblF1WTI5dE1SMHdHd1lEVlFRREV4UlNZWEJwWkZOVFRDQlNVMEVnUTBFZ01qQXhPREFlDQpGdzB4T0RBek1USXdNREF3TURCYUZ3MHhPVEF6TVRJeE1qQXdNREJhTUJVeEV6QVJCZ05WQkFNTUNpb3VibWR5DQpiMnN1YVc4d2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUURwSDlJbDZrUU05Y3JPDQpaaVJxeTJQUlg4Y2VZY3IwODdLL2l4L3I1SFNKZ2p4TUN3OXM2RTBEQ3lsMGNla0h1TmwvN2FvT0YrQ2NKbktkDQo4aEtQVThwVXVsUjJNdTE0NGY2V1Vkb29wcFYrUTY0V0pyc0w2M0xnVy9kaVJNRmZLOWppU3BTa2VWdmIwRFNTDQpEQnE4UTlNZ2theklrZUhGM1JadGEzUjZCOG5SWFVZNW1uU2pjY1hja2pUQlJLRE9hblNhOSt3cXF1MG45QzF5DQpFVVdzbTNibktQMFI2RDdZNU0zRXNsdG5XN3ZRTEhPSndialQzVC9BM00vSnk0dGJKanVSNUhHTm9ydG1XVS84DQpTRXZ0KzRTbi84MHhhNUlkL3NrblZmd0ZlU3ZITzVscEQzVXZuY1UyajNoeUpqK05jYnUyQUxvRGFiMlFuVjZGDQpaUTErQ1YybkFnTUJBQUdqZ2dLc01JSUNxREFmQmdOVkhTTUVHREFXZ0JSVHloZFovR3ZBQXlFdkdxN2txcWdjDQpnbGJhZFRBZEJnTlZIUTRFRmdRVXJibTJzUXpoK1NmSGVCOEd0Y1RHOE9xS091c3dId1lEVlIwUkJCZ3dGb0lLDQpLaTV1WjNKdmF5NXBiNElJYm1keWIyc3VhVzh3RGdZRFZSMFBBUUgvQkFRREFnV2dNQjBHQTFVZEpRUVdNQlFHDQpDQ3NHQVFVRkJ3TUJCZ2dyQmdFRkJRY0RBakErQmdOVkhSOEVOekExTURPZ01hQXZoaTFvZEhSd09pOHZZMlJ3DQpMbkpoY0dsa2MzTnNMbU52YlM5U1lYQnBaRk5UVEZKVFFVTkJNakF4T0M1amNtd3dUQVlEVlIwZ0JFVXdRekEzDQpCZ2xnaGtnQmh2MXNBUUl3S2pBb0JnZ3JCZ0VGQlFjQ0FSWWNhSFIwY0hNNkx5OTNkM2N1WkdsbmFXTmxjblF1DQpZMjl0TDBOUVV6QUlCZ1puZ1F3QkFnRXdkUVlJS3dZQkJRVUhBUUVFYVRCbk1DWUdDQ3NHQVFVRkJ6QUJoaHBvDQpkSFJ3T2k4dmMzUmhkSFZ6TG5KaGNHbGtjM05zTG1OdmJUQTlCZ2dyQmdFRkJRY3dBb1l4YUhSMGNEb3ZMMk5oDQpZMlZ5ZEhNdWNtRndhV1J6YzJ3dVkyOXRMMUpoY0dsa1UxTk1VbE5CUTBFeU1ERTRMbU55ZERBSkJnTlZIUk1FDQpBakFBTUlJQkJBWUtLd1lCQkFIV2VRSUVBZ1NCOVFTQjhnRHdBSFlBdTluZnZCK0tjYldUbENPWHFwSjdSemhYDQpsUXFyVXVnYWtKWmtObzRlMFlVQUFBRmlIRWtybVFBQUJBTUFSekJGQWlCZFRxNXZrdHJqRzZDWjltK24yRFk3DQptVEdndWJNVWpESHBFY1hJMHgzSU53SWhBTVFnZ0VpT3JGUlh0WHc4VnlZZzRlNVpGUjJhbnRCakdnWE5tWXJCDQp2K24xQUhZQWIxTjJyREh3TVJuWW1RQ2tVUlgvZHhVY0Vka0N3UUFwQm8yeUNKbzMyUk1BQUFGaUhFa3IvZ0FBDQpCQU1BUnpCRkFpRUExUTRVNmVyQ3pPWVJ1OW55UFh6MnFWY2RnL3BwVVdPREVyNWNTbEdJeVJ3Q0lIM2o0ZWMzDQpDbXROaitaN0l1T3R3YjZwMmNPSGlrRHBvNDFWTmlvM3pBcm1NQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUE1DQpCM25FZE5iU3huYmVrTUtHeFo3QlFoTi9uVVRiN0QraVMxMnIxK3BUL0VqdzhUUmEzdWVEWjdWcnNxU3AzaE1tDQpXMmYwMjJkanpocnhFSGkyYWJGb0VUS0Q1dUFSR1F2dDVML2h2bEhmZGtyZVMxSWZxK2YyUVllcU9zcVJlUUxHDQpOUTR6ekRXS1gxeTRBNGpQSi9uQ0diNk16UnlIUytHemhKMU50YnozNDJhQ2IxWER6aXBNSVpZUXhZTDFISjVTDQo5VzRYOGg2c2w4NDJXTmVvajBtYU10ZmVpajhURGlnUnJBUFE5anRYK29lOG1mdG4zOVVyK1ZvRC9FaTl3cWhMDQo1Rlh2WGR3MWxmT0RLNEpzd2JtTzdXMExwNzJnMDBHY3k5a1h1MUpoVkQvVDFFamh0NW92YjdIQ1VWYXkzT0plDQpKUE5Pemh6eTUzUzUzS0hiN0gvSQ0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ0K
    url:  https://1eafed28.ngrok.io/pod
  failurePolicy: Fail
  name: 1eafed28.ngrok.io
  namespaceSelector: {}
  rules:
  - apiGroups:
    - ""
    apiVersions:
    - v1
    operations:
    - CREATE
    resources:
    - pods
```

When our simply Koa.js app, written in Typescript, receives the request it does the following:

1. Clone the incoming pod spec into a new object
1. Make changes to the clone, updating the image location
1. Creates a JSONPatch by comparing the original and the clone
1. Base64 Encodes the JSONPatch data
1. Returns the patch as part of an \`admissionResponse\` object

The code is hopefully nice and simple to follow [so take a look at it here.](https://github.com/lawrencegripper/ClusterLocalAdmissionsController/blob/master/src/server/podCreate.ts) If you'd like a more complex example you can take a look at the `golang` code here in `istio` which uses a similar method to [inject the `istio` sidecars](https://github.com/istio/istio/blob/master/pilot/pkg/kube/inject/webhook.go)(This is what I read in order to write the Typescript example).

That's it, nice and simple.

Note: ngrok approach won't work in an Azure AKS cluster due to networking restrictions, you'll need an ACE Engine cluster or other.. or you can test inside the cluster [with the receive setup as a service](https://github.com/lawrencegripper/ClusterLocalAdmissionsController/blob/master/k8s.yaml) but beware of circular references (pod can't be created because CREATE calls webhook which is received by the pod which can't be created).
