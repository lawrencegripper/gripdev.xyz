---
author: gripdev
category:
  - uncategorized
date: "2022-01-13T20:51:54+00:00"
guid: https://blog.gripdev.xyz/?p=1553
title: What I learnt when a system no one knew how to maintain started failing, and I was on-call
url: /2022/01/13/what-i-learnt-when-a-system-no-one-knew-how-to-maintain-started-failing-and-i-was-on-call/

---
A system is failing. People rely on it. You are on-call to fix it. You don't know how it works, your team don't know how it works and the last person to work on it has left the company. Fun times!

I'll be upfront. This was an intense on-call shift. It wasn't much fun but it did help me learn some new approaches for how to handle these situations. This blog is what I was doing by the end, having learnt from doing the wrong things in places.

**Panic** **and try to find help**

As the realisation dawns that you are meant to fix a system you know nothing about, it's natural to feel some panic. I did.

Next up, try and get help. This is a time to be humble, explain that you'll do your best but be clear you don't know the tech/stack/system etc and reach out widely to see if others do.

You **MUST do this.** Think of it this way, If you struggled with the problem for ages (without asking for help) you are causing unnecessary downtime and pain for users. Imagine, hours into the outage, someone else in the organisation popped up and said "Oh this is an easy problem in stack xyz just do z - why didn't you reach out?". You've not done yourself or the organisation any favours at that point.

Being honest about what you know and don't know. **Don't try and take it all on, reach out for help. Tell people how you are feeling. Get support.**

**Write EVERYTHING down**

This is useful because:

1. In a situation you don't understand, things you think are irrelevant may later become relevant.
1. Anyone who comes to help, whether your awake, asleep or getting a coffee, can see what's been done and what is on the list to try next.
1. After you can look back, as part of a retrospective, to learn from the incident.

**Buy time**

Engineering is hard at the best of times. Doing it with the pressure of an ongoing outage makes it even harder.

Buy yourself some time with tactical hacks. These don't have to be good, they're not there to last forever, they're to give you time. In this case I wrote a script to poke the system, attempt to detect the failure and restart it when it was failing. It worked, sort of, for a bit and gave some head space.

Use the time you gain to do the nitty gritty detailed work of engineering, which I find hard to do with the time pressure associated with an outage.

**Understand the request flow - Read the code, re-read the code**

It's impossible, in my opinion, to fix a system you don't understand. Unless you get incredibly lucky. Once you've brought yourself some time, spend that time wisely.

My personal preference is to look at a typical request flow through the system.

- What does a normal request look like?
- What dependencies are called during its processing?
- Are caches involved? What state are they in?
- Are some requests more substantial than others?
- What code is on the hot path of processing?
- Where is the hardest computational work done?
- Where is the complexity? Note: Sometimes this might be in a library you use not your own code.

All of these help to give you insight into where to investigate.

Another key thing to understand is the history of the system. Systems can sometimes see issues reoccurring or new issues coming up which are new twists on previous issues. It's well worth looking through the historic issues, comments and code to understand some history of the system. It's not practical to know it all, be tactical and look at history for areas you are suspicious might be contributing to the failure.

For example, I ended up reading a commit message from 2018 and finding useful details after doing `git blame` on a file that looked of interest.

**Look at what changed but don't obsess about it**

Thing A was working for X amount of time. Now, at Y time, it's not working any more. What changed between X and Y? There must be one thing. It's an obvious conclusion to make.

Sometimes this give you the perfect answer. "Oh we deployed this change just before it broke!" you roll it back and the world is happy again.

Other times your failure state is based on a complex interplay of issues that have built up over time and combined to cause you pain. **My experience is that this is common. Usually it's not just one thing which caused it, it's an interplay of multiple things.**

Even if the failure is related to a change, it might not be a change you can control. Say the users' usage of the system has shifted. It's usually not practical to email them all and say "Please stop doing X".

**Get data: Metrics** **and Logs**

This is the first change I shipped. Not an attempt at fixing the problem but a way to get more data about the problem.

If I'd started shipping code changes without data I'd be guessing, sometimes you get lucky but most of the time you don't. You need data. Logs, metrics and any other useful sources.

**Get More Data: Debugger**

If possible get an instance that has the issue and attach a debugger. This was crucial to building a picture of the issue in my case. Minimize the impact to users, maybe you can use a secondary and leave a primary serving traffic or attach to a canary instance or a lab.

However you do it, go in with a clear idea of where you want to break and what you want to know when you hit those break points.

**Be Scientific: Build a theory, prove it's right or wrong and repeat**

At this point you've:

1. Asked for help
1. Brought yourself some time
1. Understood the system
1. Looked what might have changed
1. Got data from metrics, logs and debugging

Maybe some more than others.

Start this loop:

- Create a theory about the failure (early on these will be guesses).
- Work out how you'd prove it right or wrong.
- Test the theory.
- If it didn't fix it: Go dig more and build and new theory.
- Repeat

Then the game is simple, repeat the loop above. Add more logs if you need them. Debug more. Write down the outcomes, capture data. Build on past theories.

Start with smaller, easier to verify theories. This lets you run through more of them, and often small things can cause big problems (great addition [from @seveas](https://twitter.com/seveas/status/1481877325075546112?s=20)).

Sometimes a theory will need more data that you don't have, at this point write out a plan for what to do the next time the issue re-occurs to gather the data you need.

**Don't fall into the trap of changing too much at one time. There is a danger you introduce new issues with changes. Be measured. Is during an incident the right time to update all 38 dependencies the app has in one go (skipping QA because everything is on fire)? Probably not.** Theories help with that, a change must have a theory to justify it. That theory should be provable. If a change doesn't improve things don't pile then next one on-top. Reset and go repeat the theory loop.

When you hit a theory that appears correct **always validate the fix did what you expected.** Add metrics, state how you expect them to change after the fix is shipped. Check the system did change the way you expected.

**End**

That's my brain dump on the topic. There are doubtlessly other views on this topic from better authors with more experience in SRE than myself. I'd be really interested to hear thoughts on bits I've missed here or articles by others that related to this topic.
