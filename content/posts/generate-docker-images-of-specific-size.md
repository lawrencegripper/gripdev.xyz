---
author: gripdev
category:
  - uncategorized
date: "2020-02-24T12:14:36+00:00"
guid: http://blog.gripdev.xyz/?p=1284
title: Generate docker images of specific size
url: /2020/02/24/generate-docker-images-of-specific-size/

---
For some testing I'm doing I need a set of images of a specific size to simulate pulling larger vs smaller image.

[Here is a quick script I put together](https://gist.github.com/lawrencegripper/5c25d5fdd13a3233144d87e972b52fb2) for generating a 200mb, 600mb, 1000mb and 2000mb image (tiny bit larger as alpine included). Took a while to work out best to use `/dev/urandom` not `/dev/zero` as with `zero` the images got compressed for transfer.

![](/wp-content/uploads/2020/02/image.png?w=1024)

https://gist.github.com/lawrencegripper/5c25d5fdd13a3233144d87e972b52fb2
