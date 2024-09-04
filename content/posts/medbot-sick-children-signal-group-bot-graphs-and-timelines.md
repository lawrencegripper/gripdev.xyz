---
author: gripdev
category:
  - uncategorized
date: "2021-06-21T14:36:38+00:00"
guid: https://blog.gripdev.xyz/?p=1514
title: 'MedBot: Sick children + Signal Group + Bot = Graphs and Timelines'
url: /2021/06/21/medbot-sick-children-signal-group-bot-graphs-and-timelines/

---
This is a brain-dump rather than a fully fleshed out blog. Most of the code was written with an unwell small human sleeping on me and python isn't my best language, it's very much a hack.

[![](/wp-content/uploads/2021/06/screenshot-2021-06-19-at-08.06.37.png)](/wp-content/uploads/2021/06/screenshot-2021-06-19-at-08.06.37.png)

I have two kids, both have asthma and chest issues. Unfortunately, these are things you manage rather than cure, they're more prone to normal colds escalating quickly and need more medical interventions in general.

My oldest hasn't started school yet but has spent more time in hospital already than I have in my entire life.

"How does this relate to coding Lawrence?", Glad you asked. We keep a track of the medication, temp, pulse ox and other key events in a Signal Group.

We've found that between swapping parents, sleepless nights and different hospital wards/doctors its easy for things to get lost.

This has worked really well in the past, Signal keeps things tracked, it's quick and easy. You can write down whatever you want. If your offline it'll sync up later.

When you swap parents or see a new doctor you can do a quick rundown of what's happened in the last x hours, chase up missed doses etc just by scrolling up the chat.

What was new this time round was that both of my kids where ill at the same time, both with chest infections. Both needed medication, observations on temp, pulse ox etc and the group got messy fast.

So I decided to write something to make things nicer. Partly because I thought it would help, partly because having something to focus on helped dissipate the nervous energy of seeing your kids ill and not being able to do much about it.

The aim is a bot to pickup the messages on the group and then store them and build out views/graphs.

The stack I used is:

- Docker/VSCode devcontainer to run stuff
- [SignalD as the interface to Signal Messenger](https://gitlab.com/signald/signald)
- [Semaphore as the Bot Library](https://github.com/lwesterhof/semaphore)
- Sqlite as the data store
- Python, Pandas, Matplotlib for the graphs
- [Labella py for timelines](https://github.com/GjjvdBurg/labella.py)

First up, massive shout out to Finn for the work on Signald and to Lazlo for the Semaphore bot library that builds on it. Both of these where awesome to work with and made this project easy.

The basic aim is for the bot to listen on the group, pickup updated then pull out the relevant information and store it in a sqlite db.

I used the 'reaction' in Signal to show that the bot has successfully picked up an item and stored it, you can see this as the 💾 added to the messages below.

Last when someone sends a message 'graphs' the bot should build out graphs and share them back to the group.

[![](/wp-content/uploads/2021/06/image-2.png)](/wp-content/uploads/2021/06/image-2.png)

What does this code look like? See the Semaphone examples for a full fledged starting point (seriously they're awesome). In the meantime, I'll show my specific bits. It's surprisingly small, I added a handler to the bot to detect messages that had a temperature in them using a regex and insert them into the `temperature` table in sqlite.

https://gist.github.com/lawrencegripper/00c57d7d0152b49859eb84101b40e9a7

Then for graphing I tried out something a bit different. I used a Juypiter notebook to author and play with the code then I used `jupyter nbconvert graphs.ipynb --to python` to output the notebooks code as a python file.

This was a nice mix for a side/hack project, I could iterate quickly in the notebook but still have that code callable from the bot easily.

The handler and graph rendering look like this, I was seriously impressed with `pandas` `datafame`, I've not used it much in the past and being able to easily read in from sqlite was a big win.

https://gist.github.com/lawrencegripper/9581ee2c654a8171af0e376628200467

[![](/wp-content/uploads/2021/06/image-5.png)](/wp-content/uploads/2021/06/image-5.png)

Last was drawing the timelines, [labella was awesome here](https://github.com/GjjvdBurg/labella.py/blob/master/examples/timeline_kit_1.pdf), I had to hack a bit but it does awesome stuff like let you pick a colour for the item based on it's content. With this I could label different types of medication with different colours on the timeline.

https://gist.github.com/lawrencegripper/b2c31fc910b88353debc202e61426905

What does this look like when drawn? (Granted I've picked rubbish colors).

[![](/wp-content/uploads/2021/06/image-3.png)](/wp-content/uploads/2021/06/image-3.png)

It gives a chronologically accurate timeline with each medicine or item type easily distinguishable. This is useful to take in how things are going over 24 hours and also spot issues with missed doses.

So that's it really, I haven't published the full set of code as it's got more specific stuff to them in there, but hopefully this is a useful overview and drop comments if you'd find this interesting/useful. If there is enough interest I can clean stuff up to make this sharable.
