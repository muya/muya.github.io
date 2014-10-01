---
layout: post
title: Moving To Github Pages
---
I started my first [blog](http://encore254.blogspot.com) in August 2012. My
first post, [Creating Users & Grants in MySQL](http://encore254.blogspot.com/2012/08/create-users-and-grants-in-mysql.html),
was as a result of learning how to do just that. I was a newbie at the time, it was all very fascinating.

I've come a long way since then. Learnt Yii (did a few blog posts on that),
moved to some heavy back-end dev, then to DevOps, changed jobs, and now I'm doing
Python (with Django) at [DumaWorks](http://dumaworks.com), while personally tinkering with
iOS using Swift

Someone once said it's important to keep a blog as a developer for a multitude of reasons.
Mine, I want to give back to the developer community. That one issue I came across that
bugged me for several days, and finally solved; I want to save someone else the hassle.
I also think that writing about stuff I do helps me understand stuff better.

I started with Blogger, since that was the first thing that came to mind (I'm a Google fan).
It work(ed) for a while, but then I wanted to customize it. I felt that it was too bloated for
my needs (writing about code), and that it was too complex to customize. The layout and the
themes took too much work to customize, and I didn't want to build anything from scratch.

A few weeks back I came across GitHub Pages, and this [article](http://www.smashingmagazine.com/2014/08/01/build-blog-jekyll-github-pages/).
Version controlled blogging? A dev's dream! I write my posts in Markdown, and push them to Github.
That's it! [Jekyll](http://jekyllrb.com/){:target="_blank"} is used to build the site and it's published! It couldn't get any simpler than this!

The site is clean, minimal, and I can add code snippets easily (this was especially messy in Blogger).
For example<sub>*</sub>:
{% highlight php startinline=true %}
function generateRandomNumber($seed) {
  // chosen by fair dice roll, guaranteed to be random
  return 4;
}
{% endhighlight %}

I'm moving all my posts here, so soon Encore254 will not be available.

Until next time, see ya!



<small style="font-size: .8em; float: right;">\* Thank you [xkcd](http://xkcd.com/221/)</small>
