---
author: gripdev
category:
  - powershell
date: "2015-01-23T09:22:41+00:00"
guid: https://gripdev.wordpress.com/?p=650
title: PowerGist – Source Control for Powershell ISE with Github Gists
url: /2015/01/23/powergist-source-control-for-powershell-ise-with-github-gists/

---
![](/wp-content/uploads/2015/01/012315_0922_powergistso1.png)![](/wp-content/uploads/2015/01/012315_0922_powergistso2.png)

## Install

Get [Chocolatey](https://chocolatey.org/) (awesome command line installer for windows)

Open powershell as administrator and type:

C:\\> choco install powergist

Once that's done simply type ISE and you'll see it pop up on the right.

Login to your Github account and away you go.

**\*Warning – This was a quick project and should be considered Alpha quality to see if it was possible and/or useful. If you find a bug or issue head over to the [github repo](https://github.com/lawrencegripper/PowerGist) site to report it or fix it in a pull request\***

## About

One of the reasons I always advise people writing software to provide an API's or a plugin model is that it allows end users to enhance the functionality of the product. If it doesn't do what they want they can add it making your product better for free. It's also a nice way to POC new features, internally, off the critical path.

So when I sat next to @StuartLeeks and said "I wish Powershell ISE integrated with some kind of source control" he pointed me to the addin model for ISE and said "Bet you can't do it" (well known as a _stealth_ motivation tactic for me within the team). So I set about trying.

The addin model is awesome. It's nice and quick to get up and running and I was done writing a simple integration for Githubs Gist service (if you've never used it think git repos but for snippets) inside a day. I'll go into a bit of detail on how to write addins in a future post, in the meantime feel free to have a look at the source.
