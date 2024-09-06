---
author: gripdev
category:
  - uncategorized
date: "2019-09-24T17:56:23+00:00"
guid: http://blog.gripdev.xyz/?p=1227
title: Misusing C# 8 `using` to get Golang `defer` in DotNet Core 3
url: /2019/09/24/misusing-c-8-using-to-get-golang-defer-in-dotnet-core-3/

---
I started out with C# since then I've learned other languages and one of my favorites is Golang.

When [I was reading the release notes from C# 8 I saw the new](https://docs.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-8#using-declarations) [`using`](https://docs.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-8#using-declarations) [declaration and through it was awesome](https://docs.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-8#using-declarations)... I also realized it could be misused to give C# the `defer` keyword from Golang.

**Whats** `defer` **in Golang do?** `defer` in Golang lets you define a function that will get run when the code block exits.

The example below will print:

```
hello world
```

Code: [Try it out here](https://tour.golang.org/flowcontrol/12)

```
 package main  import "fmt"  func main() {     defer fmt.Println("world")     fmt.Println("hello") }
```

I'm a big fan of this approach as I think it's clearer to read than the standard `try finally` pattern as it lets you put the cleanup code directly below the code which is making the mess like:

```
    f := createTempFile("/tmp/defer.txt")     defer deleteTempFile(f)     writeFile(f)          // here the `deleteTempFile` method gets called
```

 **How can we do this in C# 8 with the new** `using` **syntax?**

Well now because `using var x = new thing()` exists you can write a simple class called `defer` which runs a function when the current method exits, just like in `golang`

The interesting thing is that as a `using` statement generates a `try finally` under the covers these defer functions will still run if an exception is thrown.

The example below will print:

```
Hello World! Defer 2 Defer 1 Exception thrown: It's all one wrong
```

Code:

https://gist.github.com/lawrencegripper/afc8ac5de11f4c90cbab96fdafe7b899

Should you do this? Well that's kinda up to you, I don't think it's super nasty but I've not used this in anger so probably worth testing out a bit before going all out.
