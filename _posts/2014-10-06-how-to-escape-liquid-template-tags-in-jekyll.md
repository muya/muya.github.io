---
layout: post
title: How To Escape Liquid Template Tags In Jekyll
author: Fred Muya
---


My blog has been on Github Pages for about a week now. So far, so good. I'll say this again, Jekyll + Github Pages, a dev's dream! I especially like the syntax highlighting provided by the Pygments plugin. So sweet, so simple.

One of the tricks I've learnt since is how to highlight Liquid Template code, within the Liquid template you're working on!

This is accomplished using a combination of the HTML syntax highlighting made available by Pygments, and the `{% raw %}{% raw %}{% endraw %}` tag provided by Liquid (Talk about inception!). For example, if you want to display the code snippet below:

{% highlight html %}
{% raw %}{% if post.excerpt %}{% endraw %}
{% raw %}{{ post.excerpt }}{% endraw %}
{% raw %}{% else %}{% endraw %}
{% raw %}{{ post.content | truncatewords:30 }}{% endraw %}
{% raw %}{% endif %}{% endraw %}
{% endhighlight %}

Then this is the code you should have in your markdown file:

![Escaped Template Tag]({{ site.url }}/images/2014-10-06/escaped-template-tag.png)

Easy to select, huh? :stuck_out_tongue_winking_eye:

Grab the gist [here](https://gist.githubusercontent.com/muya/8785e948688a49d83b14/raw/e0f322436e513f165289f12b053b70921815387d/gistfile1.txt)!

Happy hacking!
