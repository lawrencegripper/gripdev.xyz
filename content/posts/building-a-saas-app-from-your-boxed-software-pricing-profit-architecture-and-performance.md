---
author: gripdev
category:
  - thoughts
date: "2017-01-22T07:28:52+00:00"
guid: https://gripdev.wordpress.com/?p=854
summary: |-
  While working as a consultant I spent time with a number of large companies helping them move from a tradition boxed software model to providing SaaS applications.

  One of the first hurdles to overcome is the pricing model for their new SaaS app and the cultural change to pricing SaaS apps vs traditional boxed software.

  In the old world the sales, product and marketing team would have a discussion about the value proposition of the product, take into account cost of creating the code, understand the price elasticity and market demand.

  Once that’s done the dev team would produce a machine minimum spec required to run the solution and burn the code to a disk or zip up a download for the customers.

  At that point all is done, trouble would brew if future versions needed significantly higher spec hardware. That might put pressure on a dev team to optimize but the pricing decision wasn’t part of this process.

  Now let’s look at this is a SaaS world, the pricing discussions has to take into account the running costs of hosting the solution. There are two approaches at this point, ignore this completely and continue down the old path OR treat this as a lever you can control and use it to make the product more successful.

  Let’s take two examples of decisions that now, in the SaaS world, have a huge impact on the cost of running your solution and, ultimately, your profit margin.
title: Building a SaaS App from your Boxed Software – Pricing, Profit, Architecture and Performance
url: /2017/01/22/building-a-saas-app-from-your-boxed-software-pricing-profit-architecture-and-performance/

---
While working as a consultant I spent time with a number of large companies helping them move from a tradition boxed software model to providing SaaS applications.

One of the first hurdles to overcome is the pricing model for their new SaaS app and the cultural change to pricing SaaS apps vs traditional boxed software.

In the old world the sales, product and marketing team would have a discussion about the value proposition of the product, take into account cost of creating the code, understand the price elasticity and market demand.

Once that’s done the dev team would produce a machine minimum spec required to run the solution and burn the code to a disk or zip up a download for the customers.

At that point all is done, trouble would brew if future versions needed significantly higher spec hardware. That might put pressure on a dev team to optimize but the pricing decision wasn’t part of this process.

Now let’s look at this is a SaaS world, the pricing discussions has to take into account the running costs of hosting the solution. There are two approaches at this point, ignore this completely and continue down the old path OR treat this as a lever you can control and use it to make the product more successful.

Let’s take two examples of decisions that now, in the SaaS world, have a huge impact on the cost of running your solution and, ultimately, your profit margin.

## Multitenant vs Single Tenant

Maybe you have an existing boxed product and you want to make a fast move to the cloud to offer a no-hassle hosted option to your customers. The risk of doing this slowly is that you may lose market share to others who move first or worse still, maybe new entrants who aren’t burdened by their existing code and have written for SaaS from the start.

You could start a complete re-write, creating a multitenant solution with low hosting costs, built for the SaaS world.

The opportunity cost here is the time and resources needed to execute this strategy and the % risk of failure.

Alternatively, you could create a single tenant solution, using your existing code. Maybe creating IaaS VMs and masking the complexity from your customers - offering a SaaS solution without large scale re-writing of your existing product.

The trade-off here is the higher hosting cost associated with a single tenant solution. Many use this as justification for starting a ground up re-write for their software, for some this is correct but for many this is a failure to account for the costs/risk combo involved in this endeavor.

So how does this affect pricing and profit? Well the multitenant rewrite involves high expenditure over an indeterminate period of time in hope of future profits. The single tenancy example sacrifices current profit margin to get a foothold the market. Depending on the market and your business either of these could be the correct option.

## ROI and Solution Performance

In the boxed software world - code quality, performance and efficiency were focused on ensuring good user experience and happy customers.

That’s not the case anymore. If you’re moving to SaaS, they directly affect your profit margin. Releasing inefficient code or designing inefficient architectures costs you money.

So what’s the best response? Go nuts on performance testing, micro optimize every line of code and jam as much on one box as possible?

Well yes and no, just like with multitenancy you can use this as a lever. Maybe you’re losing ground to a competitor, to close the gap you need to increase the speed with which you create new features. You know the margins you have to play with and the future delivery roadmap so you can take a calculated risk to lower the focus on performance in favor of shipping the new features needed to close the gap.

Actually it may transpire that the uptake on the feature is very small but it helps your product compete. At this point you may decide not to revisit the feature and improve its performance, due to the low usage. In this case you’ve saved time and effort which would be exerted optimizing unnecessarily.

Again, we’ve minimized upfront investment of resources in favor of potentially higher costs while you get solid data to justify the time and effort required to optimize the solution.

## Summary

Successful SaaS isn’t about technology, business or operations in isolation. It’s about all of these working together. It’s about making decisions with their proper context. It’s about having meetings where an key decision maker, dev lead, product manager and sales lead all take time to understand the implications of their actions and explain them to each other. It’s about trade-offs, margins, code performance, velocity, hosting costs, market research, competitors … you get the gist.

Above all, it’s about being tactical. Be aware of what decisions cost, monitor the outcome and take calculated risks.
