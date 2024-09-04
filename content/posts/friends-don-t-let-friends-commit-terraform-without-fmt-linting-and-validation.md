---
author: gripdev
category:
  - '#terraform'
  - devops
  - linux
  - uncategorized
date: "2019-07-16T19:38:20+00:00"
guid: http://blog.gripdev.xyz/?p=1197
tag:
  - '#terraform'
  - devops
  - linux
title: Friends don't let friends commit Terraform without fmt, linting and validation
url: /2019/07/16/friends-dont-let-friends-commit-terraform-without-fmt-linting-and-validation/

---
So it starts out easy, you write a bit of `terraform` and all is going well then as more and more people start committing and the code is churning things start to get messy. Breaking commits block release, formatting isn't consistent and and errors get repeated.

Seems a bit odd right, in the middle of your devops pipe which dutifully checks code passes tests and validation you just give `terraform` a free pass.

![Captain Picard Quotes. QuotesGram](https://proxy.duckduckgo.com/iu/?u=http%3A%2F%2Fmedia-cache-ec0.pinimg.com%2F736x%2F3c%2Fae%2F5a%2F3cae5a1df502615cdb41872423ebf667.jpg&f=1)

The good new is `terraform` has tools to help you out here and make life better!

Here is my rough script for running during build to detect and fail early on a host of `terraform` errors. It's also pinning `terraform` to a set release (hopefully the same one you use when releasing to prod) and doing a `terraform init` each time to make sure you have providers pinned (if not the script fail when a provider ships breaking changes and give you an early heads up).

It's rough and ready so make sure your happy with what it does before you give it a run. For an added bonus the `docker` command below the script runs it inside a `Azure Devops` container to emulate locally what should happen when you push.

https://gist.github.com/lawrencegripper/d5f126279a1991eee5ed2a200234029e

Optionally you can add `args` like `-var java_functions_zip_file=something`  to the `terraform validate` call.

Hope this helps as a quick rough guide!
