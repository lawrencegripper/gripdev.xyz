---
author: gripdev
category:
  - how-to
date: "2013-07-06T17:34:09+00:00"
guid: http://gripdev.wordpress.com/?p=349
title: Visual Studio 2013 – How to get rid of the clones
url: /2013/07/06/visual-studio-2013-how-to-get-rid-of-the-clones/

---
Hi all,

A while back I did a talk @wpug about the work I'd done when I re-wrote one of my apps with a new shared core. One of the things I did in the talk was compare the stats from my code before and after the move.

In short I had a play around with the "Analyze" menu in VS2013 - looking at complexity, lines of code and the dependency graphs.

The one I missed that I should have mentioned is the code clone function. Why is this useful, because duplication is [EVIL.](http://en.wikipedia.org/wiki/Don't_repeat_yourself) While the "don't repeat yourself" idea has more to it than duplication it's a good starting point for anyone. If you write something twice it can break twice and need fixing twice. Trust me you'll forget.

![](/wp-content/uploads/2013/07/070613_1732_visualstudi1.png)

The results:

![](/wp-content/uploads/2013/07/070613_1732_visualstudi2.png)

You can see a list of items that are similar grouped into matches, if you select 2 files you can then compare to see what's different between the two.

![](/wp-content/uploads/2013/07/070613_1732_visualstudi3.png)

In this file you can see that I have the EXACT same block of code twice, only lines away from each other.

As they're in the same file the fix is nice and simple:

![](/wp-content/uploads/2013/07/070613_1732_visualstudi4.png)

It's practically all done for you.

Yes there are going to be tougher ones down the road but it's worth a quick look to see if you have any easy wins in the list and working through the harder ones too. Ultimately you'll thank yourself at some point in the future.

N.B I wrote this a while back and just found it in my draft so thought I'd publish
