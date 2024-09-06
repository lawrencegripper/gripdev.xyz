---
author: gripdev
category:
  - docker
  - golang
  - mlock
  - terraform
  - vscode
date: "2020-07-14T12:18:40+00:00"
guid: http://blog.gripdev.xyz/?p=1321
summary: |-
  I recently upgrade my machine and and installed the latest Ubuntu 20.04 as part of that.

  Very smugly I fired it up the new install and, [as I use devcontainers,](https://code.visualstudio.com/docs/remote/containers) looked forward to not installing lots of devtools as the Dockerfile in each project had all the tooling needed for VSCode to spin up and get going.

  Sadly it wasn't that smooth. After spinning up a project which uses `terraform` I found an odd message when running `terraform plan`

  > failed to retrieve schema from provider "random": rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: EOF
  >
  > error from `terraform plan`

  Terraform has a `provider` model which uses GRPC to talk between the CLI and the individual providers. `Random` is one of the HashiCorp made providers so it's a really odd one to see a bug in.

  Initially I assumed that the downloaded provider was corrupted. Nope, clearing the download and retrying didn't help.

  So assuming I'd messed something up I:

  1. Tried changing the docker image using by the devcontainer. Nope. Same problem.
  2. Different versions of `terraform`. Nope. Same problem.
  3. Updated the Docker version I was using. Nope. Same problem.
  4. Restarted the machine. Nope. Same problem.

  Now feeling quite frustrated I finally remembered a trick I'd used lots when building my own terraform providers. I enabled debug logging on the terraform CLI.

  `TF_LOG=DEBUG terraform plan`

  This is where it gets interesting...
tag:
  - docker
  - golang
  - mlock
  - terraform
title: 'Terraform, Docker, Ubuntu 20.04, Go 1.14 and MemLock: Down the rabbit hole'
url: /2020/07/14/terraform-docker-ubuntu-20-04-go-1-14-and-memlock-down-the-rabbit-hole/

---
I recently upgrade my machine and and installed the latest Ubuntu 20.04 as part of that.

Very smugly I fired it up the new install and, [as I use devcontainers,](https://code.visualstudio.com/docs/remote/containers) looked forward to not installing lots of devtools as the Dockerfile in each project had all the tooling needed for VSCode to spin up and get going.

Sadly it wasn't that smooth. After spinning up a project which uses `terraform` I found an odd message when running `terraform plan`

> failed to retrieve schema from provider "random": rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: EOF
>
> error from `terraform plan`

Terraform has a `provider` model which uses GRPC to talk between the CLI and the individual providers. `Random` is one of the HashiCorp made providers so it's a really odd one to see a bug in.

Initially I assumed that the downloaded provider was corrupted. Nope, clearing the download and retrying didn't help.

So assuming I'd messed something up I:

1. Tried changing the docker image using by the devcontainer. Nope. Same problem.
1. Different versions of `terraform`. Nope. Same problem.
1. Updated the Docker version I was using. Nope. Same problem.
1. Restarted the machine. Nope. Same problem.

Now feeling quite frustrated I finally remembered a trick I'd used lots when building my own terraform providers. I enabled debug logging on the terraform CLI.

`TF_LOG=DEBUG terraform plan`

This is where it gets interesting...

The log read as follows:

```
2020-07-13T19:32:19.720Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: plugin address: network=unix address=/tmp/plugin278719704 timestamp=2020-07-13T19:32:19.720Z
2020-07-13T19:32:19.720Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: runtime: mlock of signal stack failed: 12
2020-07-13T19:32:19.720Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: runtime: increase the mlock limit (ulimit -l) or
2020-07-13T19:32:19.720Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: runtime: update your kernel to 5.3.15+, 5.4.2+, or 5.5+
2020-07-13T19:32:19.720Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: fatal error: mlock failed
2020-07-13T19:32:19.722Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4:
2020-07-13T19:32:19.722Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: goroutine 7 [running]:
2020-07-13T19:32:19.722Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: runtime.throw(0xf85069, 0xc)
2020-07-13T19:32:19.722Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4:    /opt/goenv/versions/1.14.0/src/runtime/panic.go:1112 +0x72 fp=0xc000601608 sp=0xc0006015d8 pc=0x432dd2
2020-07-13T19:32:19.722Z [DEBUG] plugin.terraform-provider-random_v2.3.0_x4: runtime.mlockGsignal(0xc000803800)
```

So the go version `1.14.0` was throwing a panic while the `random` provider was trying to execute and it was complaining about.

```
runtime: mlock of signal stack failed
```

Looking around the internet I found an issue on Github explaining that this is t [o workaround a Kernel bug](https://github.com/golang/go/issues/37436#issuecomment-591033727) but, the odd thing, was that ubuntu 20.04 appeared to have a version of the kernel which didn't have this bug.

[So it looks like there is a bug in the workaround that Go](https://github.com/golang/go/issues/40184) put in place which means the Ubuntu Kernel is wrongly detected as having the issue.

Now, you might think, why not use a different base docker image for the container as then you'll have a different distro/kernel - not so lucky as the kernel for the docker host is [shared between all containers running on it.](https://github.com/docker-library/golang/issues/320#issuecomment-591103426)

> Good reminder that containers aren't Virtual Machines. Stuff is shared with the host and not everything is isolated.

So it looks like the only way around it, for now, is to set my `memLock` limit higher. Sounds easy right? You'd think you just update that ... oh yeah where would I update that for it to affect the container?

I tried:

1. [`/etc/security/limits.conf`](https://ss64.com/bash/limits.conf.html.) Nope. Same problem
1. Adding an upped limit into the `docker.service` systemd definition. Nope. Same problem.

Then I found the `--ulimit` option for the `docker run` command. The following command would set a infinite memlock limit and we wouldn't observe the issue with docker.

```1.14-rc-alpine

I wondered can you set this in the docker daemon so it's wired up to all containers started? Yes it turns out you can, thanks oracle.

So setting this into your /etc/docker/daemon.json does the trick after restarting the docker daemon sudo systemctl restart docker.service

{
 "default-ulimits": {
  "memlock": {
      "Name": "memlock",
      "Hard": -1,
      "Soft": -1
    }
 }
}

Now 4 hours later I can run terraform plan like normal people inside my devcontainer. Wow, wasn't expecting that journey.

```
