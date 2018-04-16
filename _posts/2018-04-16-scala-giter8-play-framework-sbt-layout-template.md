---
layout: post
title: Giter8 Template for Scala Play! Framework with SBT Layout
author: Fred Muya
---

I've been using [Scala Play! Framework](https://www.playframework.com/documentation/2.6.x/ScalaHome) for a little over a year now, and I've come to favor the [SBT layout](https://www.playframework.com/documentation/2.6.x/Anatomy#Default-SBT-layout) over the default layout.

In all that time, I haven't come across a template that I can quickly use to get myself up & running, so I decided to roll out my own, using [Giter8](http://www.foundweekends.org/giter8/Combined+Pages.html)!

It's been on my list for a while, and since I'm about to start on a new project, I decided this was the time to do it.

I present, [`muya/play-sbt-scala-seed.g8`](https://github.com/muya/play-sbt-scala-seed.g8) (and the crowd goes mad with joy 🎉) -- hehe :)

To use it, simply do:
{% highlight bash %}
$ sbt new muya/play-sbt-scala-seed.g8
{% endhighlight %}

It'll ask you a few questions, if you don't feel like answering, it has some sensible defaults, press ENTER to continue.

This is still v1 of it, I plan to update it to match the official, non-SBT one provided by [Play!](https://github.com/playframework/play-scala-seed.g8)

Anyway, feel free to use it for your Play Framework projects that need an SBT layout.

Happy Coding! I hope this shaves off a few minutes off your coding time.
