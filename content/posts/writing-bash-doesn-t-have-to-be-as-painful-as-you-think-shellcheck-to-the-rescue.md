---
author: gripdev
category:
  - azure
  - bash
  - devops
  - linux
date: "2019-07-29T16:45:24+00:00"
guid: http://blog.gripdev.xyz/?p=1201
tag:
  - azure
  - bash
  - devops
  - linux
title: Writing Bash doesn't have to be as painful as you think! Shellcheck to the rescue.
url: /2019/07/29/writing-bash-doesnt-have-to-be-as-painful-as-you-think-shellcheck-to-the-rescue/

---
So I've found myself writing lots of `bash scripts` recently and because they tend to do real things to the file system or cloud services they're hard to test... it's painful.

![neverfear](/wp-content/uploads/2019/07/neverfear.jpg)

So it turns out there is an awesome linter/checker for `bash` called [`shellcheck`](https://github.com/koalaman/shellcheck) which you can use to catch a lot of those gotchas before they become a problem.

There is a great plugin for [`vscode` so you get instant feedback when you do something you shouldn't.](https://marketplace.visualstudio.com/items?itemName=timonwong.shellcheck)

Better still it's easy to get running in your build pipeline to keep everyone honest. Here is an example task for Azure Devops to run it on all scripts in the `./scripts` folder.

https://gist.github.com/lawrencegripper/2abfb2580d81ef35b2fccbe4b4884009

Next on my list is to play with the `xunit` inspired testing framework for `bash` [called shunit2](https://github.com/kward/shunit2) but kinda feel if you have enough stuff to need tests you should probably be using `python`.
