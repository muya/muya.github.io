---
layout: post
title: Unshorten URLs with Expand URL Chrome Extension
excerpt: 'Chrome extension for previewing Shortened URLs before opening them'
date: '2014-05-11T18:08:00.002+03:00'
author: Fred Muya
tags:
- Apps
modified_time: '2014-05-11T18:08:59.247+03:00'
---

I’ve been having a problem opening some shortened links on my browsers (Chrome & Firefox),
whereby in some cases I end up with this screen:

![No Data Received Error Chrome]({{ site.url }}/images/2014-05-11/error-on-shortened-url.png)

Not sure if it’s just me but…

I decided to create a solution while at the same time tinkering with Chrome Extensions.

I developed an Extension called ExpandURL that you can use to ‘expand’ shortened links before opening them in your browser. I used some of the URL ‘unshortening’ APIs available to accomplish this, including:

- UrlEx.org

There are two ways to use the extension:

- Automatically re-direct to the ‘expanded’ URL after clicking
- Preview the ‘expanded’ URL before navigating to it (InfoSec guys, I see you!)

You can adjust these settings using the Options page of the Extension.

While I was using this to solve an actual problem that I was having, it was also an experiment on how to make a good user experience as possible.

I’ll publish a post soon on my experience building Chrome Extensions.

The extension is available on the Chrome Web Store: [https://chrome.google.com/webstore/detail/expand-url/llnlabdcccdihhemekhhddgbaonclebe](https://chrome.google.com/webstore/detail/expand-url/llnlabdcccdihhemekhhddgbaonclebe")

You can check out the application’s source code on my [GitHub](https://github.com/muya/expand-url)
