---
author: gripdev
category:
  - cookie
  - invoke-webrequest
  - powershell
date: "2015-05-27T18:14:19+00:00"
guid: https://gripdev.wordpress.com/?p=720
tag:
  - cookie
  - invoke-webrequest
title: Powershell Invoke-WebRequest with a cookie
url: /2015/05/27/powershell-invoke-webrequest-with-a-cookie/

---
Nice and quick post here, mainly so I remember when I need it again, this is a quick sample which shows how to make a web call from PowerShell including a cookie.

Most old methods will suggest using the WebClient object but the new (well newer than the WebClient) Invoke-WebRequest commandlet is a much nicer, in my opinion.

It takes a bit of fooling around to get this up and running so wanted to share.

{{< gist lawrencegripper 6bee7de123bea1936359 >}}
