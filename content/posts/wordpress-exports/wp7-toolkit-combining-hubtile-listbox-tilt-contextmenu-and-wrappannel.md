---
author: gripdev
category:
  - how-to
date: "2011-08-27T08:48:35+00:00"
guid: http://gripdev.wordpress.com/?p=62
title: 'WP7 Toolkit: Combining Hubtile, Listbox, Tilt, ContextMenu and WrapPannel'
url: /2011/08/27/wp7-toolkit-combining-hubtile-listbox-tilt-contextmenu-and-wrappannel/

---
![](http://bugail.com/wp-content/uploads/2011/08/HubTileSample.png)

Found a much clearer set of how-to's by windows phone geek [here](http://www.windowsphonegeek.com/articles/Windows-Phone-HubTile-in-depth-Part1-key-concepts-and-API).

This is a post for the developers out there. I wanted to use the [hubtile control](http://bugail.com/index.php/2011/08/developing-with-mango-hubtiles/) from the [wp7 silverlight toolkit,](http://silverlight.codeplex.com/releases/view/71550) but along side it also use the context menu from the same toolkit, while having the whole lot databound inside a listbox. It took me a while to come up with a solution that I was happy with. I'm not sure if it's following best practices but it seems to work well for what I want. The solution is to use a wrappannel from the same toolkit and then to place this inside "ItemsPanelTemplate" of the listbox. This means that the listbox items that are bound then appear within the wrappanel while still being part of the listbox. Then the context menu and the hubtile control can be put inside the "ItemTemplate", coupled with the [tilt animation,](http://fiercedesign.wordpress.com/2011/02/23/wp7-and-using-the-tilt-effect/) yet again from the toolkit, you end up with a nice tilt-y and scrollable set of hubtiles that have a long press context menu associated with them. If you're curious to see how it all looks, the finished product is included in my latest beta of BBC News Mobile. Hopefully this is a useful how-to for, it's my first shot at once of these so go easy!
