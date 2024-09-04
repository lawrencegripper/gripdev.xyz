---
author: gripdev
category:
  - paperboy
date: "2013-06-22T14:37:00+00:00"
guid: http://gripdev.wordpress.com/?p=319
title: 'Paper Boy Preview Coming - Sneak peak. '
url: /2013/06/22/paper-boy-preview-coming-sneak-peak/

---
Hi All,

So its been a busy couple of months my end, I've started a new role and getting up to speed has meant the time I had to make changes to the BBC News Mobile code base was limited.

I'm happy to say though that I have been working on it, over the last couple of evenings and weekends, and have a preview to release in the near future. All being well the "Preview" version will go live on the store early next week after I've come back from Glastonbury festival (escaping from all tech for a week). The preview is your chance to get to grip with the app and shape its future before it eventually replaces the BBC News Mobile app through the store update process.

First things first, yes its very orange - hopefully it isn't too visually offensive!

New Panorama - Kinda like the old one with some tweaks

![home](/wp-content/uploads/2013/06/home.png?w=180)[![mostread](/wp-content/uploads/2013/06/mostread.png?w=180)](/wp-content/uploads/2013/06/mostread.png)![sections](/wp-content/uploads/2013/06/sections.png?w=180)

New Reader view:

Feed "groups" allow multiple news sources to be combined into a since view and browsed neatly. This is where the most significant amount of time has been spent as I wanted to keep the previous performance but with multiple feeds being loaded this can get tricky. I've setup a database to allow the app to query all items and pull back those associated with a group then step through updating those and adding in an new items while your reading. This has the up side of hopefully improving offline reading as well, although this is currently still it's early stages.

The article reader is now using the lovely instapaper mobilizer to provide a great, supportable and robust in app reading experience. Massive thanks to instapaper for the awesome mobilizer! [http://mobilizer.instapaper.com/m](http://mobilizer.instapaper.com/m).

![newspapers](/wp-content/uploads/2013/06/newspapers.png?w=180)![reader](/wp-content/uploads/2013/06/reader.png?w=180)

Settings and sections has also had a large amount of change, you can now use the selectable jump list to build out your own groups (more functionality to come here) and also select your favor home and most read feeds for the main panorama screen.

Breaking news, fueled by twitter, is also included by is currently a work in progress so not fully functional in this build.

[![settings2](/wp-content/uploads/2013/06/settings2.png?w=180)](/wp-content/uploads/2013/06/settings2.png)[![settings](/wp-content/uploads/2013/06/settings.png?w=180)](/wp-content/uploads/2013/06/settings.png)

So that's pretty much the whole of it right now. Check back on the 1st of July for a public beta, I'll post a link that anyone can click and install on a WP8 device with no need to register or dev unlock, then give me lots and lots of feedback please!

For the devs out there this release gave me an opportunity to play with the sterling nosql OOB to store and query items in the app, well worth a look, [http://sterling.codeplex.com/](http://sterling.codeplex.com/). I also moved to a new model using Reactive extensions to update ViewModels, hopefully get a chance to put together a post on this soon, also very fun to play with once you get your head around it.
