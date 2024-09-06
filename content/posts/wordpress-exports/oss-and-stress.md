---
author: gripdev
category:
  - uncategorized
date: "2019-03-21T14:59:52+00:00"
guid: http://blog.gripdev.xyz/?p=1189
title: OSS and Stress
url: /2019/03/21/oss-and-stress/

---
I wanted to take a moment to lay out a journey I didn't really realize I was on until this morning. Like many big problems it all came out when dealing with a simple little GitHub issue that was raised on a repo I published about 3 years ago.

It turns out that some of the stuff in this repo was out of date, and the repo wasn't useful anymore. Someone nicely pointed this out to be in an issue, and suggested updating the README to warn others. I said "Sure, happy to accept a PR to do that" to which they closed the issue and didn't raise a PR. I have to admit that this (not contributing a simple PR to update the Readme) upset me a lot, and I over-reacted in my response.

I then realised 2 things…

1.       My over-reaction wasn't really about this issue

2.       Am I really expected to maintain YEARS and YEARS of contributions and repositories while doing my day job?

### How did I end up here?

Well, it starts slow. You publish a few things while you’re working on a pet project, or helping someone out.  You keep doing that for years and years and years, and then sit back and watch the stress poor in.

In the early days I loved getting issues raised, reviewing PRs and fixing stuff up. I’d get 1 or 2 a month, I didn't have a family, I was a junior dev, and I loved programming.

Fast forward 8 years, I'm a dad, a senior dev and I still love programming BUT now I get 5 or 6 issues every week. Each takes 30 mins as it requires me to context switch. The PR's (if they're big) take hours to review and test. I kidded myself that if I added more automation, CI, integration testing that would help... well it did, but it doesn't stop the problem . What it’s allowed me to do is stack up **more** projects and **more** stuff to look after. More things for people to contribute to and raise issues on.

So today I blew up after waking up to failed integration tests from my day job, 2 other PRs and a backlog of 3 issues that need at least 3 days of my time, and then this little issue that “broke” me. My current day job is already eating into my time with my daughter (2 years old and awesome ... don't want to miss anything), and now the things that I’ve published because I thought they were useful to others are causing me stress. Where do I put this stuff? How do I not let people down?

### What am I going to do?

Well this one is harder but I can't see any other way forward. I need to take time out, go through my projects, and mark stuff as "Dead" or "Finished". I need to step back from looking after things I no longer use myself, and make it clear to others that PRs and issues are unlikely to get my attention. Where possible, I need to hand over the code to others, or make it clear that the code I’ve contributed needs to find maintenance from elsewhere.

It'll let people down, it'll mean stuff people might rely on doesn't get updated, but at the moment I'm putting their needs ahead of my own and I can't square the circle. It'll make me sad to let some people down. I think marking some of my repos as “unmaintained” at least makes it clear to others what level of support they’ll get, but it still makes me sad that there will be people out there that may have taken a dependency on something I’ve no longer got the time (and, in some cases, interest) in maintaining.

In future, I also need to make sure that I'm considering the implicit cost of putting things out into the open. Some contributions are easy, but some have an ongoing cost which racks up and compounds over time with other contributions. Spotting these, and being cautious about them, feels like the way to go.

Lastly I have a new appreciation for those who have contributed WAY more than I have and no idea how they keep up with it all!
