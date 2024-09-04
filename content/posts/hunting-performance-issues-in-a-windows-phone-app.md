---
author: gripdev
category:
  - paperboy
date: "2013-09-07T13:48:57+00:00"
guid: http://gripdev.wordpress.com/?p=376
title: Hunting Performance issues in a Windows Phone App
url: /2013/09/07/376/

---
Hi All,

I spent the morning today working on a performance issue I had with paperboys startup time. As they're isn't a huge amount of info out there on how to do this on the phone I thought I'd blog along as I worked through the problem.

First port of call is the Windows Phone application analysis tool in Visual studio. This will hopefully help me work out what calls are going on at the time when I see the delay.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf1.png)

You've got a couple of settings, here I'm using the execution profiling to see if I can spot what's going on when the apps loading up.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf2.png)

Sure enough it shows a massive gap where the applications responsiveness is really low right after the app starts up.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf3.png)

Selecting that that time slice and you'll be able to drill down further into what's going on.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf4.png)

Head into Frames->Functions and you'll see the functions that are being called at that time.

Sort by method name to see your methods, otherwise you may find it hard to separate them from the phones calls which are also collected.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf5.png)

Investigating those methods there is one that looks dodgy and so we'll jump into it to take a more detailed look at what its up too. You can do this by clicking on the method name and it'll jump you to the code.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf6.png)

So now we've got a hunch that it's this code we're going to have to switch tacks and get more hands on with the code.

Instrumenting the code will let you get detailed info from the method that you're interested in. As a simple way to do this on the phone I adapted this code from code overlord. [http://codeoverload.wordpress.com/2012/01/15/timing-a-method-in-c/](http://codeoverload.wordpress.com/2012/01/15/timing-a-method-in-c/)

I ended up with this:

![](/wp-content/uploads/2013/09/090713_1407_huntingperf7.png)

All the lines highlighted are added in for the profiling, the rest is the method as it was.

Spinning this up on the phone and then glancing at the output window it became pretty clear where things were going pear shaped.

loading existing items
 Staring databse query -- 1 ms
 Adding context to item -- 25 ms
 Check for duplicates -- 43 ms
PROFILER :: Total 4689 ms

As it turns out I was loading all the items ever seen on the homefeed since the app was first installed, this wasn't working out well!

So I did this:

![](/wp-content/uploads/2013/09/090713_1407_huntingperf8.png)

The first time I ran this on a phone that had lots of items in the database unsurprisingly things where still slow. loading existing items
 Staring databse query -- 1 ms
 Adding context to item -- 25 ms
 Add items -- 38 ms
PROFILER :: Total 4990 ms

Then then second time I ran it I got a significant improvement as all the older and duplicate items are no longer loaded.

loading existing items
 Staring databse query -- 1 ms
 Adding context to item -- 27 ms
 Add items -- 42 ms
PROFILER :: Total 1865 ms

Sweet so we've making progress, this is a fairly big difference we're down from 4.9 seconds to 1.8 seconds. Odd though, this still seems to take longer that it should I mean the items are actually returned really quickly… or are they?

Well no they're not. I'm using SterlingDatabase which is using lazy evaluation so the items aren't actually retrieved until I start the foreach loop. Making the change below, for testing, forces it to be evaluated ahead the loop and surprise surprise it takes longer! Now we can see the real cost of the database query highlighted in yellow below.

\_newsRepo.GetItemsByFeedUri(this.feedUri).ToList() asIEnumerable<INewsItem>;
loading existing items
Staring databse query -- 1 ms
 Adding context to item -- 357 ms
 Add items -- 374 ms
 check is old or duplicate -- 397 ms
add to collection -- 414 ms
 check is old or duplicate -- 1497 ms
 add to collection -- 1513 ms
 check is old or duplicate -- 1565 ms
 add to collection -- 1578 ms

Even with this though it's only an extra 357ms in 1865ms so it's not all that much. Adding an item seems to take an age the first time (highlighted in red), what's going on here? Looking at the Observable collection I'm doing lazy loading here too so that explains that as the first add command creates a new observable.

![](/wp-content/uploads/2013/09/090713_1407_huntingperf9.png)

The creation of that observable collection appears to cost us ~1000ms.

So now we've got:

350ms returning the items from the DB
1000ms creating a new ObservableCollection
13ms per item for checking duplicates and age and adding to the collection.

We've saved ~3000ms by not repeating these checks repeatedly on every load and instead doing some simple housekeeping.

Hopefully this is useful for others grappling with app performance, not to say this is the only way to attack the problem but this is now I set about doing it this morning!

P.s Anyone using paperboy this update is going into the store now and will mean a really nice performance improvement when opening the app!
