---
layout: post
title: 'Fix Images Not Loading on Jekyll Blog Running on Localhost'
date: '2022-08-28'
author: Fred Muya
excerpt: 'How To Fix Images not Loading on Jekyll Blog Running on Localhost in Brave Browser <br><br><strong>TL;DR:</strong> Set Brave Shields to Down'
tags:
- jekyll
- jekyll-images
- brave-browser
---

¡Hola! Welcome to another post.

In this (short) one, we're documenting a fix for images not loading for your Jekyll blog running on localhost, when using Brave Browser.

I came across this issue while drafting a new post locally, and noticed that images were not loading.

In the browser console, I was seeing `ERR_BLOCKED_BY_CLIENT` errors. This indicated to me that the Brave was probably blocking the images.

I was able to resolve this by setting Brave Shields to down for my blog. 

Alternatively, accesssing the blog on `0.0.0.0` instead of `localhost` resolves the issue.

The issue seems related to Brave being aggressive with external images blocking.

Until next time, stay safe, and happy coding!
