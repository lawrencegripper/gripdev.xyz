---
author: gripdev
category:
  - uncategorized
date: "2023-06-28T09:04:16+00:00"
guid: https://blog.gripdev.xyz/?p=1622
title: 'Atuin + Codespaces: Sync command history between Codespaces and local'
url: /2023/06/28/atuin-codespaces-sync-command-history-between-codespaces-and-local/

---
In my normal day I work in my local machine, in codespaces and in devcontainers.

Recently I started using [Atuin](https://atuin.sh/) which is an awesome tool for syncing your command history. Go check it out, it's really nice!

A really cool feature is you can search history in different categories, like commands on this host or commands in this directory. Here I set this up for Codespaces so the `host` is the name of the repository. This means my commands from any instance of that repo's Codespace are sync and searchable when in another new or existing Codespace.

Atuin has a public sync server but I prefer to sync to a local instance I host. This means I need a bit of wiring up to make this work when in a codespace or devcontainer.

When using Codespaces I tend to do the following:

1. Use [Chezmoi](https://www.chezmoi.io/) to configure the [Codespace automatically via my dotfiles](https://docs.github.com/en/codespaces/customizing-your-codespace/personalizing-github-codespaces-for-your-account#dotfiles)
1. Open VSCode connect to the Codespace
1. Open [Kitty](https://sw.kovidgoyal.net/kitty/) and SSH to the Codespace with [gh cli](https://cli.github.com/), to have a terminal window connected to the Codespace

With this flow I need a way to allow calls from within the Codespaces to connect to the Atuin server I'm running on my local network.

To do this I'm going to use [SSH Reverse Port Forwarding](https://blog.devolutions.net/2017/03/what-is-reverse-ssh-port-forwarding/), you've all probably forwarded a port from your local machine to a remote machine before but did you know it supports the other way around too? You can forward a port from the remote machine to a destination the local machine has access too.

The other bit to tackle is authentication from the Atuin instance in the Codespace to the local server. To get this wired up I'm using [SCP](https://linux.die.net/man/1/scp) to copy the session and key files up to the Codespace.

Then I can use the templating features in Chezmoi to customize my Atuin and zshrc files when running in a Codespace. The `ATUIN_HOST_NAME` env lets you override the host name the commands are tracked against. In my case I set this to `codespace/$GITHUB_REPOSITORY` so I can search commands by repo (neat feature of Atuin is you can view command history by host, cwd and more).

I want a nice way to launch this all, this should display a menu which lets me select a Codespace and have all of this done automatically from there. For me I'm using [Rofi](https://github.com/davatorium/rofi) which is a nice launcher for linux but fzf or another would work fine here too.

As a bonus I also use [LanguageTool](https://dev.languagetool.org/http-server) spell/grammer checking service, again self hosting. This, [via VSCode extension](https://github.com/davidlday/vscode-languagetool-linter), gives me spell checking for my Markdown files. I can forward the ports for this in the same way as I did for Atuin.

Lastly, adding in a key binding for [`i3`](https://i3wm.org/) and I can now `WIN+P` and select my codespace then get all the bits I want open with language tools and atuin hooked up. When I move to a new codespace I get to keep my command history 🎉

Putting this all together you get (files trimmed to only the relevant bits).

{{< gist lawrencegripper 28fa370211dae966cc28efc369c902b5 >}}
