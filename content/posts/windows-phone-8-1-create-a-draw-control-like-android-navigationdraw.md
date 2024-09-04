---
author: gripdev
category:
  - windows-store
date: "2014-06-01T19:30:24+00:00"
guid: http://gripdev.wordpress.com/?p=555
title: Windows Phone 8.1 - Create a Draw control like Android NavigationDraw
url: /2014/06/01/windows-phone-draw-control-like-android-navigationdraw/

---
Hi,

This is something I wanted to do with my HypeMix app, the aim was that the user would be able to navigate around the app but, on any page, open the draw to see the currently playing track etc.

Luckily, now Windows Phone 8.1 is in line with WinRT, we have the idea of Frames for Navigation.

Using this I created a FramePage, this is the first page the app navigates to and it contains the Draw Content, Frame (for subsequent pages we navigate to) and some animations for show and hide.

Once my custom FramePage has loaded it sets up the app so all navigation happens inside this page, which contains the draw overlay.

As there are quite a few moving parts, rather than copy and paste LOADS of snippets, I've put the whole thing up on GitHub for you to play with.

\[embed\]https://github.com/lawrencegripper/WindowsPhone8--NavigationDrawerExample\[/embed\]

Ps. Quick and dirty demo hacked together from the actual app, no lovely UI, just some nice bright colors to demonstrate different panels and pages. In my app I've combined this with a pub/sub model and a navigation service to avoid the code behind and make things a bit nicer, left out of this demo to aid simplicity.

Hopefully this is useful to others!
