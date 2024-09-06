---
author: gripdev
category:
  - debugging
  - delve
  - golang
  - vscode
date: "2017-11-01T11:54:11+00:00"
guid: http://blog.gripdev.xyz/?p=1077
tag:
  - debugging
  - delve
  - golang
  - vscode
title: Pass arguments to Golang program when debugging with VSCode
url: /2017/11/01/pass-arguments-to-golang-program-when-debugging-with-vscode/

---
I'm doing some work on a golang project, the code takes a relative file path in as an argument.

Now I have [Delve](https://github.com/derekparker/delve/blob/master/Documentation/installation/osx/install.md), [VSCode](https://code.visualstudio.com/) and [VSCode-Go](https://github.com/Microsoft/vscode-go) installed which means I can have a nice interactive debugging session. The only snag I hit was that it wasn't clear how to pass in my arguments when the debugger started my go code.

The trick is the use a double dash to indicate which arguments should go to the code, other args, before this, will go to the delve debugger. Also don't forget to use "cwd" to set the working dir used when debugging.

Here is an example with comments

https://gist.github.com/lawrencegripper/43f4bd28061181a593b0a08f348f2cef

And we're away...

![Screen Shot 2017-11-01 at 11.52.36](/wp-content/uploads/2017/11/screen-shot-2017-11-01-at-11-52-36.png)

N.B The behavior is a bit odd, as the code here suggests that the "--" should be appended for you. https://github.com/Microsoft/vscode-go/blob/master/src/debugAdapter/goDebug.ts#L306. However, this did not work for me, may be fixed in the future so double check!
