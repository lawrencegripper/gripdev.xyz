---
author: gripdev
category:
  - quick-post
date: "2019-03-27T10:08:22+00:00"
guid: http://blog.gripdev.xyz/?p=1191
title: Avoiding pushing secret stuff to Git by accident
url: /2019/03/27/avoiding-pushing-secret-stuff-to-git-by-accident/

---
So it seems like a brain dead simple one. Don't push secrets by accident, make sure you check and update the projects `.​gitignore` to ignore sensitive files but the reality is different.

One example, you use Terraform and set the ignore file to ignore the state file. Then later another developer moves the folder the Terraform is in and updates the ignore. Now when you merge you get the updated `.gitignore` and if you don't pay attention all your state files get pushed in your next commit.

#### Whats the solution?

Global Git Ignores! Yes they exist and are easy to use. [Check out this guide](https://help.github.com/en/articles/ignoring-files#create-a-global-gitignore)

So using this you can setup a nice rule like this:

\[code lang=text\]
\*.private\*
private.\*
\[/code\]

Now next time you create a file you NEVER want to end up in a commit all you have to do is name it `secretstuff.private.env` and your safe.

It's saved me loads and I can't recommend it enough - also you can update your global with more specific stuff like Terraform or whatever else you want.
