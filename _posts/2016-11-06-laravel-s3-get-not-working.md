---
layout: post
title: Laravel Storage 'get' method not working with S3
author: Fred Muya
---

I came across a weird issue when working with the Laravel Storage module. I'd uploaded some files to S3, but for some reason when I used the `exists()` method to check if the file exists, it kept returning false.

It turns out that if your server time (I was working on my Vagrant box) is out of sync, then the method will return false.

I used the `ntpdate` command to update the server time, and the method started returning the correct value.

Now, I need to figure out why my ntp wasn't doing its job.

Happy Coding!
