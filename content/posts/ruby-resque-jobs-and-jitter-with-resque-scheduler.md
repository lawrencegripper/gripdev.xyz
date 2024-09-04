---
author: gripdev
category:
  - uncategorized
date: "2023-08-21T12:51:51+00:00"
guid: https://blog.gripdev.xyz/?p=1637
title: 'Ruby: Resque Jobs and Jitter with `resque-scheduler`'
url: /2023/08/21/ruby-resque-jobs-and-jitter-with-resque-scheduler/

---
[Resque](https://github.com/resque/resque) is a background job processor for Ruby. Sometimes you need to do something that'll take a long time and you don't want that happening as part of the HTTP request lifecycle. It helps you do that.

But what happens when you want to do LOTS of things at once, you want to avoid the [thundering herd problem](https://en.wikipedia.org/wiki/Thundering_herd_problem) by spreading that work out rather than doing it all at once.

[Jitter](https://en.wikipedia.org/wiki/Jitter) is one way to do this, taking an example using Resque, say you want to recalculate something for products and you have 10,000 products. Here we can define a resque job to do the recalculation, taking in the product ID and then queue a job for each but we don't want them to all start at the same time due to the load this would put on the DB. To fix that we add jitter, saying start these jobs somewhere between now and 120seconds time, picking a random duration per job.

To do this for Resque we can use [resque-scheduler](https://github.com/resque/resque-scheduler), which lets us queue work for the future, along with [`before_enqueue` resque hook](https://github.com/resque/resque/blob/master/docs/HOOKS.md) to let a job declare its jitter with, for example, `@jitter_milliseconds = 10` on the job class.

Here is an example of this all wired up (note: We ended up using a slightly different code for the prod implementation due to requirements on our side, be sure to test the below works for you before adopting).

https://gist.github.com/lawrencegripper/8e701b0d201e65af0f8bc9b8b0b14207
