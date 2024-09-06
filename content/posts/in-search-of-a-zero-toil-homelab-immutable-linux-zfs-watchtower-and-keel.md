---
author: gripdev
category:
  - cloud
  - containers
  - devops
  - docker
  - fcos
  - fedora-coreos
  - homelab
  - immutable-linux
  - kubernetes
date: "2024-03-16T08:38:51+00:00"
guid: https://blog.gripdev.xyz/?p=1660
tag:
  - cloud
  - containers
  - devops
  - docker
  - fcos
  - fedora-coreos
  - homelab
  - immutable-linux
title: 'In search of a "Zero Toil" HomeLab: Immutable Linux, ZFS, WatchTower and Keel'
url: /2024/03/16/in-search-of-a-zero-toil-homelab-with-immutable-linux/

---
I run a HomeLab for hosting a few bits (for example `atuin`, `omnivore` and `matrix-bridges`)

Like all software these things need to be kept up-to-date. **This saps the fun out of hosting things.**

**I do not want any toil**, here is the goal:

- I don't want to manually do OS updates
- I don't want to manually update versions of software that I'm hosting
- I don't ever want to have to worry about losing data
- **I always** want a simple, quick, rollback when somethings breaks
- I want a good security posture

Previously I had pretty standard lab. It had a VM host (running `libvirt`) and inside that I ran a few VMs.

Some ran `docker-compose` applications, another larger VM ran a simple single node Kubernetes instance (using `K3s`).

This had all the toil, updates etc, and I wanted it gone (without losing any of the functionality).

## Immutable Linux

The underpinning of any lab is the OS.

The traditional option here, to reduce the headache of updates, to configure auto updating. For Ubuntu this is via `apt`.

This has a couple of issues:

1. OS updates, say from Jammy to Focal are still manual
1. No easy quick rollback if things go 💩 … only option is get on the console and start hacking to bring things back

Luckily there is a pretty cool emerging trend in Linux distributions for "Immutable Linux". With immutable linux you don't update the system in place, you stage a full replacement and then swap from old to new.

For folks familiar with blue/green application deployments this should sound familiar. The only difference is this is down at the OS level.

[Ubuntu Core](https://ubuntu.com/core) and [Fedora CoreOS](https://fedoraproject.org/coreos/) are the two distributions that I started looking at in this space.

After some experimentation with Ubuntu Core I abandoned it. It's only accommodating of `snap` based packages (and a large set of customisation functionality is pay walled). This meant running `k3s` or `tailscale` wasn't possible as they're not on the snap store 😭

Luckily Fedora CoreOS (FCOS) doesn't have these limitations. You can, within reason, do nearly everything you would do on a normal linux distribution **plus you get a bunch of awesome stuff**.

It's [auto updating by default](https://docs.fedoraproject.org/en-US/fedora-coreos/auto-updates/), with support for roll back between the immutable system snapshots. All updates, whether full major OS version bump or package level, are handled automagically! 🎉

This is all done using some pretty cool tech - [`rpm-ostree`](https://github.com/coreos/rpm-ostree). You setup a node with `butane` which is similar to `cloud-init` letting you [add users, SSH keys](https://coreos.github.io/butane/examples/#users-and-groups), [Systemd units](https://coreos.github.io/butane/examples/#systemd-units), [configure timeslots for updates](https://github.com/coreos/rpm-ostree) and more. Butane runs once at startup, it's great for bootstrapping.

The defaults in FCOS also have security in mind, it runs with [SELinux](https://www.redhat.com/en/topics/linux/what-is-selinux) enabled, has a very small default footprint with limited set of packages and, as the majority of the OS is readonly, so there is less opportunity for malware to manipulate the system.

**But there's more….** What about managing the host after it's been bootstrapped by Butane?

FCOS lets you manage the software on the node **with a `Dockerfile`**, yeah that's right, [you can build a docker image and the host system files will be built of that](https://www.opensourcerers.org/2023/06/16/using-ostree-native-containers-as-node-base-images/) ( [this repo shows how I do this in detail](https://github.com/lawrencegripper/fcos-native-container-example)). This enables some awesome flows, like update the docker file then have CI build and push at it. At which point the hosts will start updating to this new version, no `puppet` or `ansible` dance. Full declarative management of hosts just like your application container with the same toolchain.

The shifting left here is pretty awesome, **mess up a systemd units syntax and you find out when CI builds the image** not when it's been applied to nodes, and they stop starting the service!

So this becomes my base, FCOS is the hosting layer for the lab. I have a butane file per host and a `fcos-base.Dockefile` which packages ( `tailscale`, `compose`). Together butane and the image control the hosts.

No more toil on the OS level 🤞. Onwards and upwards through the stack we go! 🚀

## Data Snapshots and rollback

It's all well and good having immutable linux but application have state, this is another contender for toil and pain.

Recently I updated an app that has a `postgres` database, the update did a database migration mutating the state of the DB data and **it went wrong**.

The good news is I'm setup to handle this, as well as "immutable linux" I've moved to using `Copy-On-Write` filesystem [ZFS](https://en.wikipedia.org/wiki/ZFS) run on [TrueNAS Scale](https://www.truenas.com/).

In each of the VMs hosted in the lab I use `libvirt`'s `ISCSI` support to mount a ZFS `/mnt/data` directory in which all the applications state is stored.

Why the ISCSI setup at the Hypervisor level? Simple, I don't to have to setup connection details and packages on each VM. By mounting ISCSI at the `libvirt` level all the VMs see normal disks and the hypervisor does all the hard work.

On TrueNAS I have [Periodic Snapshots](https://www.truenas.com/docs/core/coretutorials/tasks/creatingperiodicsnapshottasks/) configured nightly, these don't duplicate data so they're efficient, they only store the diffs.

With this setup the "Bad migration corrupts state" problem becomes a non-problem. Like the Immutable CoreOS I have a snapshot I can rollback to easily and pick up the data from the last snapshot.

That's not always ideal though, maybe you want to use the old data to fixup the current data.

No problem, you don't even have to lose the current state, you can restore the old snapshot and keep the current `HEAD` \- then compare the old vs new data to fix things up

[![](/wp-content/uploads/2024/03/image.png)](/wp-content/uploads/2024/03/image.png)[![](/wp-content/uploads/2024/03/image-1.png?w=1024)](/wp-content/uploads/2024/03/image-1.png)

Here I clone a point in time snapshot into a new mountable ISCSI disk

Now I have current and old side-by-side for fixing bits up 🥳

Another bonus to this setup is, when I abandoned `Ubuntu Core` and moved to `FCOS` no need to copy any app state or config, it was all in `/mnt/data`, I just had swap out the `.qcow2` OS image and reboot the VM, `cd /mnt/data && docker compose up` and app running, unworried about the OS change.

## K3s Cluster Updates

Running a single node Kubernetes cluster is useful, lots of stuff is packaged as `helm` charts and I also don't want to manage a VM for every application I run. K8s gives a neat way, with nice UX, to pack a bunch of small apps onto a cluster. (Bonus: [Tailscale Operator](https://tailscale.com/kb/1236/kubernetes-operator) makes is super easy to expose these as HTTPS endpoints to your tailnet too)

**But as with any abstraction, this doesn't come for free**. K8s can be pretty mammoth 🦣 … giving us another set of things to update and maintain.

The cool thing we avoid most of this pain with the minimalist `k3s` distribution of Kubernetes. It's slimmed down [and has auto-updating](https://docs.k3s.io/upgrades/automated).

For example, it doesn't use the full stack of K8s, [swapping out `etcd` for the simpler `sqlite`](https://docs.k3s.io/datastore). **The upside is a less complex system.**

Let's get it configure to update itself without any toil from us. Once we install the `upgrade-controller` we can `kubectl apply` an update plan, like so 👇, which says "keep yourself up-to-date with the latest stable release". **That's it, no need to spend Sunday updating K8s after, a week of doing it at work 🤩**

```yaml
  apiVersion: upgrade.cattle.io/v1
  kind: Plan
  metadata:
    name: server-plan
    namespace: system-upgrade
  spec:
    concurrency: 1
    cordon: true
    nodeSelector:
      matchExpressions:
      - key: node-role.kubernetes.io/control-plane
        operator: In
        values:
        - "true"
    serviceAccountName: system-upgrade
    upgrade:
      image: rancher/k3s-upgrade
    channel: https://update.k3s.io/v1-release/channels/stable
```

## Auto-updating Applications

I've got two modes of hosting `docker-compose` and `kubenetes` and these are both on-top of `Fedora CoreOS` hosts.

For both I want the docker images kept up-to-date with updates published by me or the upstream owners.

**Warning:** There are downsides here, if someone pushes a bad update, it's going to get picked up and run. Like anything it's worth considering the pro's and con's.

Let's [take `synapse`, the Matrix server,](https://github.com/element-hq/synapse) if they release a new version I want it, I don't want to have to think about it. If it goes wrong, I can move back to the old version, even if it nukes my data we have that covered through the ZFS snapshots.

For images/components I'm building I want similar, when I `docker push` a new image and I want that picked up and run.

To do this for `docker compose`, I use [`Watch Tower`, it's a little go utility](https://containrrr.dev/watchtower/) which polls the docker repository for new versions and connects to the hosts `docker socket` to update running containers to these new versions.

As FCOS has SELinux by default containers are prevented from accessing the docker socket and host files (good right!). We need to explicitly allow only watch tower an exception to the docker socket with [`--security-opt label=disable`](https://github.com/containrrr/watchtower/pull/1917) ( [Docs PR here](https://github.com/containrrr/watchtower/pull/1917)). Note `:z` after the mounts too, again by default containers can't mount host files, [`Z` modifies the selinux label of the host file or directory being mounted](https://docs.docker.com/storage/bind-mounts/#configure-the-selinux-label) into the container. The end command is 👇

```
$ sudo docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock:z --security-opt label=disable --restart=always containrrr/watchtower
```

To finish it off, I've set it up [its notifications](https://containrrr.dev/watchtower/notifications/), using my [`SMTP -> Signal` bridge](https://github.com/lawrencegripper/signald-smtp-bridge), to ping me a note when it does it. That's it, no more manual update of docker-compose based apps. I get these nice messages on [Signal](https://www.signal.org/) when it updates something:

[![](/wp-content/uploads/2024/03/image-2.png?w=1024)](/wp-content/uploads/2024/03/image-2.png)

**What about the Kubernetes application you ask?**

Here [I use `Keel` which is a similar utility](https://keel.sh/), instead of watching the docker socket it monitors the `deployments` in my K8s cluster.

`Keel` has some nice features, above and beyond `watchtower`

1. It [support `semver` policies which control](https://keel.sh/docs/#policies) what updates are applied via `annotation` on the Kubernetes deployment.
1. You can use the [`approval` feature](https://keel.sh/docs/#approvals) if there is something you'd like to first 👀 before applying

Put together this looks like 👇

```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: atuin
    namespace: atuin
    annotations:
      keel.sh/policy: major     # update policy (available: patch, minor, major, all, force)
      keel.sh/trigger: poll     # enable active repository checking (webhooks and GCR would still work)
      keel.sh/approvals: "1"    # required approvals to update
      keel.sh/pollSchedule: "@every 1h"
```

As with `watchtower` I configured `keel`'s notifications to use the [`SMTP -> Signal`](https://github.com/lawrencegripper/signald-smtp-bridge) bridge to get notifications of updates going out.

[![](/wp-content/uploads/2024/03/image-3.png?w=1006)](/wp-content/uploads/2024/03/image-3.png)

A nice added bonus is that Kell it has a UI showing you what updates have been applied

[![](/wp-content/uploads/2024/03/image-5.png?w=1024)](/wp-content/uploads/2024/03/image-5.png)

**What about package updates on Fedora CoreOS?**

Well… guess what … I can do the same, auto updating to docker images (and I love it).

With `rpm-ostree` [I can rebase the node to use my custom OCI image](https://www.opensourcerers.org/2023/06/16/using-ostree-native-containers-as-node-base-images/). With that done OS updates are handled through docker images.

I build the docker image and push it to the repository (in this case I rebuild it daily to pickup upstream updates and package updates). I have a simple Systemd Timer which runs on the box and periodically runs `rpm-ostree upgrade --reboot` this will pull a newer container version, if one exists, and update the node to it (if you don't want to reboot the node you can do `rpm-ostree upgrade && rpm-ostree apply-live` then restart affected services).

Here is an example `Dockerfile` for CoreOS. It adds `tailscale` and `compose` into a base image. I then build another version on top of that with `k3s` for the Kube nodes.

```dockerfile
FROM quay.io/fedora/fedora-coreos:stable as homenet-fcos-base
RUN date >> /etc/homenet-fcos-build-at.txt
ADD ./rpm-repos/tailscale-stable.repo /etc/yum.repos.d/
RUN rpm-ostree install tailscale && ostree container commit
# Add bits we'd like in our base images /etc or /usr folders (remember ostree-native containers don't suppor other paths like /var)
COPY overlay-base/ / # This has my systemd units etc to copy into the image
RUN systemctl enable tailscale-configure.service && systemctl enable tailscaled.service && systemctl enable update.service && ostree container commit
# Build a k3s specific image
FROM homenet-fcos-base as homenet-fcos-k3s
ENV INSTALL_K3S_BIN_DIR=/usr/bin
ENV K3S_KUBECONFIG_MODE="644"
ENV INSTALL_K3S_SKIP_ENABLE="true"
COPY overlay-k3s/ /
RUN curl -sfL https://get.k3s.io | sh - && ostree container commit
RUN systemctl enable k3s.service && ostree container commit
# Bring k3s specific config

```

The only difference you'll notice is [`ostree container commit`](https://coreos.github.io/rpm-ostree/container/#using-ostree-container-commit) is run after each `RUN` command. This validates and does some clean-up.

Other than that normal docker rules apply.

I love that this shifts left the work of updating OS and it's packages to CI docker build step, have you ever had an update fail mid run because of a network drop of host glitch? In theory this resolves that issue, an image either builds or does, it's either pulled onto the host and staged for update or it's not.

[This repo is an example](https://github.com/lawrencegripper/fcos-native-container-example) of this all hooked together

## Summary

I've talked about how a stack of Fedora CoreOS, Docker Compose and Kubernetes can be setup for a minimal toil home lab.

Using docker images to update everything from the raw host to the running applications and keeping those up-to-date with `keel`, `watchtower` and `rpm-ostree`.

I also covered how I mitigate issues with application state data during updates through the use of `zfs` on `TrueNAS` with periodic snapshots.

As a final note, before you think about following me down this road. One thing **I'm always trying to balance with the HomeLab is complexity**. I'm fairly ruthless, if a component, approach or setup isn't robust it's removed.

This new stack hasn't had 6 months to bake yet, the jury is still out. I'm new to this space and still learning, mistakes have likely been made.

I'm interested to see how this new setup behaves over the next 6 months. It has more complex than what went before it, on the flip side... it's also tackling more problems. So far I'm hopeful I'm onto something. Only time will tell if I'm right.
