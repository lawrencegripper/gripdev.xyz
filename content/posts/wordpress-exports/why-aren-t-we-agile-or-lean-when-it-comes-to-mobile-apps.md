---
author: gripdev
category:
  - thoughts
date: "2014-11-13T18:24:27+00:00"
guid: http://gripdev.wordpress.com/?p=619
title: Why aren’t we agile or lean when it comes to mobile apps?
url: /2014/11/13/why-arent-we-agile-or-lean-when-it-comes-to-mobile-apps/

---
I've been writing apps for a long time now, stretching back to the Windows Mobile days, something completely new dawned on me in the last few months ago as I took had a bit of break for Wedding planning and moving house – For a long time I'd not been using agile, I've not been lean.

As it the natural reaction to these things I first went to blame others "Store certification forces me to release complete apps", "I can't iterate often with the certification process, I'd end up cutting corners to release quickly" and "I don't have a single customer I have lots - how do I manage a backlog or understand user stories."

While all of these are true they're also fixable and this is a brief look at how I've gone about fixing these over the last 6 months.

## "Store certification forces me to release complete apps"

For this I'm going to use the HypeMix example, this is a music app that I've written to let you listen to Hypem.com from a native client.

In the original release, which came to the store as Windows 8 Launched, I did a big bang. I looked at some features I wanted, set about making them and released the completed product. I was completely focused on the features I wanted and valued, I didn't have a view to the future and I didn't incorporate any user feedback or scope for it in the future.

The outcome was that the app was released, did what I wanted but not much of what others wanted. As I had been solely focused on getting it out and in the market I'd also written it in a way that made it nearly impossible to adapt. Tightly coupled classed, interwoven XAML and generally high levels of technical dept.

So when I had some time off I started the architecture from fresh. I approached it pragmatically and developed in a lean way.

1. I wrapped up past nastiness and legacy I had created in reasonable interfaces. I couldn't afford to re-write the lot so I limited its impacted. This allowed me to quickly refactor and extend without losing ground.
1. I released often to a small set of testers and as often as I could to the store. While the store implies that the app should be complete that only means you need to have complete Features advertised, not all the features you intend to write.
1. I addressed issues of technical debt by using a pub/sub models and MVVM rigorously to loosely couple components and create a maintainable codebase.

**What was the outcome?**

In the space of three months I completely revamped the app, livened up the interface and released 3 updates to the store. MVVM combined with Pub/Sub has now left me with an app I can extend, flex and maintain easily. Users are happy to see the progress. I can respond quickly to requests.

So while the first of the three releases, maybe even the current release, isn't at the level I would normally have been happy with I'm seeing a great response from the users.

This for me has been the major learning from HypeMix. Users love progress. So what - you broke a feature - if you fix it the next week and add 3 more users will average these out and most will like it. You can add in unit testing and TDD to reduce the chances of causing regression issues as you move forward.

## "I can't iterate often with the certification process, I'd end up cutting corners to release quickly"

This is just something that you tell yourself. Hit the minimum bar for store certification and you can push out an update in 10mins flat. Make sure you have privacy statements, assets and icons in your source control and a quick set of smoke tests that you run. Once you have this you're sorted. Yes the first publish is slow. If you wait another 4 months the next is also slow but if you do this every month and get into the routine it's a non-issue.

When it comes to cutting corners my take is you need to be lean and pragmatic. This may seem at odds with my HypeMix example but it's not. HypeMix could have been a complete failure, if I'd originally sunk loads of time into it and over engineered the solution I'd have less time for other projects, lost momentum and emphusiasm. Striking a balance and being pragmatic is key. If you swing too far into hacker or perfection you'll either never ship anything or make a Frankenstein monster that no one wants!

## "I don't have a single customer I have lots how to I manage a backlog or understand user stories."

For this point I'm going to look at "BBC News Mobile" and now "Paperboy".

The level of feedback I got on both of these projects because a welcome problem, but I couldn't respond well. What I missed was correlation of feedback and usage data. As soon as I looked at features and feedback next to each other the task became much much simpler.

Put in simple terms, power users are amazing and give loads of feedback but tend toward less used features. Normal users give hardly any feedback and focus on the major features of the app.

When you combine something like flurry analytics with uservoice you can target your response. Working out the weighting is hard but for example, if a bug is reported on a highly used feature by 2 users then you better listen and fix it. If 4 users give feedback on a lesser used feature then by all means fix it but don't do it before the other one!

In short, not only can you get this feedback effectively with services like uservoice, you can also react intelligently by combining it with other data sources such as usage data in your app.
