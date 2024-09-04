---
author: gripdev
category:
  - bbc-news-mobile
date: "2012-01-07T12:34:18+00:00"
guid: http://gripdev.wordpress.com/?p=125
title: 'Release: BBC News Mobile'
url: /2012/01/07/release-bbc-news-mobile/

---
Hi all,  new version 2.9 is now live on the marketplace.

Thanks for everyone who helped with testing the beta. Special mentions go to:

- derausgewanderte (load errors)
- Ro (& fix, landscape view issues)
- Pia (crash dumps for close bug)
- Andrew h (picture errors)
- Hil Hughes (pin articles feature req)
- Simon H (smaller font sizes for titan)
- Tezza
- gerry
- Pat
- bjorn

Who all spotted bugs/gave detailed feedback without which this release wouldn't have been what it is! Hope I haven't missed anyone out.

## Known bugs:

A couple of bugs crept in and a bug fix release is being certified at the moment.

**Updating the app causes background task error: too many tasks**

- Work around:  un-install/re-install
- Fix: Already going through certification. v3
- Why did this happen? To get the speed boost I made significant changes to the way the app stored settings data. While I did regression test the upgrade path, another fix I implemented before publishing had the knock on effect of causing this error.

**Latest news page wrongly shows latest sports news**

- Workaround: Currently none available
- Fix: Already going through certification. v3
- Why did this happen? Developer error.

**Wrong pictures displaying on articles**

- Workaround: Currently none available
- Fix: Under investigation
- Why did this happen? Under investigation

**No longer able to share via email or SMS**

- Workaround: Currently none available
- Fix: Already going through certification. v3. This version will contain Share via sms/email in the app bar menu.
- Why did this happen? I've recently changed the app to use the integrated share picker. However, this is exclusively for social networks, which I didn't realize at the time.
