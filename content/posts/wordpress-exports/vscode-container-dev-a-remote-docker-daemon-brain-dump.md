---
author: gripdev
category:
  - braindump
  - containers
  - vscode
date: "2019-09-17T11:25:48+00:00"
guid: http://blog.gripdev.xyz/?p=1218
tag:
  - braindump
  - containers
  - vscode
title: VSCode Container dev a Remote Docker Daemon [Brain dump]
url: /2019/09/17/vscode-container-dev-a-remote-docker-daemon-brain-dump/

---
Note: This is more a stream of consciousness than a blog post. It's not detailed and mainly for my own memory on how to set this up. Be cautious.

So I've recently built a new home server/lab and wanted to use some of it's power to run my daily dev stuff.

I've been playing a bit with `VSCode Remote for Containers` which lets you define you dev environment as a container and commit it along with your code... it's super nice.

So how about hooking it up to a remote docker daemon so that I can use the power of the new server from my laptop when I'm at home but switch back to my local docker deamon when I'm traveling... turns out the team have thought of that!

There are some great docs here on how to do this in detail: https://code.visualstudio.com/docs/remote/containers-advanced#\_developing-inside-a-container-on-a-remote-docker-host

What will follow is only useful if your like me, running Ubuntu on your laptop and also Ubuntu/Linux on the remote Daemon machine.

First up one of the key problems with this setup is with the file system access as a remote docker daemon on the server can't see files on the Laptop. To get around this I've setup two `alias` which use `rsync` to `push` and `pull` data to and from the server from my laptop.

```
alias docker-remote-push="rsync -rlptv --progress \${PWD} \"lawrence@ubuntudev:\${PWD}/../\""

alias docker-remote-pull="rsync -rlptv --progress \"lawrence@ubuntudev:\${PWD}\" \${PWD}/../"
```

Then I can use `docker context` to setup a context for `local` and my `ubuntudev` remote box to easy switch between them.

Have a look here on how to create a context. https://docs.docker.com/engine/reference/commandline/context\_create/

The last `alias` are to foward the docker for the remote deamon to my local box and switch to the context I've created for it.

```
alias docker-remote-start="ssh -fNL localhost:23750:/var/run/docker.sock lawrence@ubuntudev && docker context use remote"

alias docker-remote-stop="docker context use local"
```

Then all you need to do is add

```
 "docker.host":"tcp://localhost:23750"

```

To your `settings.json` in VSCode and now you are set to go.

Open up a folder in VSCode with a \`.devcontainer\` then run

```
docker-remote-push
docker-remote-start
```

Then use the \`Reopen in container\` option in VSCode and it's done.
