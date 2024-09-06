---
author: gripdev
category:
  - devops
  - docker
  - k8s
  - kubernetes
  - yaml
date: "2024-02-26T17:43:13+00:00"
guid: https://blog.gripdev.xyz/?p=1651
tag:
  - devops
  - docker
  - k8s
  - kubernetes
  - yaml
title: 'HomeLab: Using 1Password CLI to handle Secrets in Kubernetes/Compose YAML'
url: /2024/02/26/homelab-using-1password-cli-to-handle-secrets-in-kubernetes-compose-yaml/

---
My HomeLab runs a few useful things (like [atuin](https://atuin.sh/), [changedetect.io](https://github.com/dgtlmoon/changedetection.io), [matrix bridges](https://github.com/mautrix/whatsapp) and [homebridge](https://homebridge.io/)) either in via Docker Compose or hosted on a single node [K3s](https://docs.k3s.io/) Kubernetes cluster.

In both cases YAML all sits in a versioned `git` repository. If things go bad I can recreate the lab from scratch without too much pain (along with ZFS snapshots of data drives thanks to [TrueNAS](https://www.truenas.com/)).

Using this basic Infrastructure as Code approach makes this nice, for example rolling back bad changes is a `git reset HEAD~1 --hard` then a `kubectl apply`.

The sticking point was secrets. Nearly all applications have them, for pulling images, encrypting data or logging into other systems. I don't want these checked into the `git` repo or floating around the place, it's easy to `git push` to the wrong remote, for example.

When looking to address this I found lots of "heavyweight" options in this space with Kube, like [Vault provider](https://developer.hashicorp.com/vault/docs/platform/k8s) or the [Azure KeyVault provider](https://learn.microsoft.com/en-us/azure/aks/csi-secrets-store-driver) but the key rule for my HomeLab (and for prod systems I work on) is ...... 🥁

> **"What is the simplest, most reliable solution to the problem"** i.e Lets not build a spaceship when all we need is a bike

With that in mind I started playing with the 1Password CLI, it's something I already have and use for storing secrets, could it do what I need here?

Yeah, it can 🎉 and, in my opinion, in a pretty neat way with [`op inject`](https://developer.1password.com/docs/cli/reference/commands/inject)

Using `atuin` as an example, first I setup a `homenet` entry in my 1Password vault. Then add items, grouped by sections for each application, containing the secrets I need.

Then use on the item use the `Copy Secret Reference` option to get a reference to that secret.

[![](/wp-content/uploads/2024/02/image.png)](/wp-content/uploads/2024/02/image.png)

Now I can paste that reference into the `sercret` YAML for the kube deployment like so 👇 (see [template syntax](https://developer.1password.com/docs/cli/secrets-template-syntax/) for more options/details)

```
apiVersion: v1
kind: Secret
metadata:
  name: atuin-secrets
  namespace: atuin
type: Opaque
stringData:
  ATUIN_DB_USERNAME: atuin
  ATUIN_DB_PASSWORD: "{{ op://Personal/homenet/atuin/db_password }}"
  ATUIN_HOST: "atuin-api.mytailnet-here.ts.net"
  ATUIN_PORT: "80"
  ATUIN_OPEN_REGISTRATION: "true"
  ATUIN_DB_URI: "postgres://atuin:{{ op://Personal/homenet/atuin/db_password }}@postgres/atuin"
immutable: false
```

Instead of using `kubectl apply -f ./my-secret.yaml` to apply this, I can use `cat ./my-secret.yaml | op inject | kubectl apply -f -`  

When you run this command you'll get a prompt from 1Password to allow access, it'll then expand out the `Secret Reference` into the actual secret from your 1Password vault.

With that sorted for a single file, lets expand this to all the YAML K8s config in the repo with a little script

```
#!/bin/bash
set -e
for pathname in $(find ./k3s -type f -name \*.yaml); do
    echo "Applying $pathname"
    cat $pathname | op inject | kubectl --cluster k3s apply -f -
done
```

Note: I tend to avoid using things [like `helm`](https://helm.sh/) directly on the cluster and instead rendering the YAML that `helm` generates into file in the repository, for example `helm install frigate blakeblackshear/frigate -f values.yaml --dry-run=client > frigate.yaml` to give me an applyable YAML for [Frigate](https://docs.frigate.video/). Again, I like simple and not having `helm` in the mix does that, it gets me started and I can then tweak the YAML as I like.

That's it, same approach works nicely for the `docker-compose.yaml` files when running these on VMs. I expand out the template then this is pulled into the Fedora CoreOS VM image with Butane or `rsync`'d up to the right place, I'll blog about this more (hopefully) and update here with a link for that.
