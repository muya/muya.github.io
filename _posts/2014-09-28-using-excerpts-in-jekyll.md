---
layout: post
title: Using Excerpts in Jekyll
author: Fred Muya
---

I was migrating my blog posts from Blogger to Github Pages today. Not the most interesting thing to be doing on a Sunday (Blogger exporter, y u export to HTML? I WANT MARKDOWN!)

I'm using the brilliant example given [here](http://www.smashingmagazine.com/2014/08/01/build-blog-jekyll-github-pages/){:target="_blank"}
as the basis for my blog, and the items that load on the front page are the
first few words of each posts. This is determined by this tag in the `index.html` file:
{% highlight html %}
{% raw %}{{ post.content | truncatewords:30 }}{% endraw %}
{% endhighlight %}
which means the first 30 words of the post will be used as the excerpt.

This brings an issue when you have images within the first 30 words of your post. The result is your image is loaded on the landing page of your blog, and the formatting for excerpts below the image is messed up.

Fortunately, Jekyll supports excerpts in YAML's front matter. Here's how to do it:

- Add `excerpt: "Your awesome excerpt"` to your post's front matter, i.e. between the dashes at the top of your page
{% highlight yaml %}
---
excerpt: 'Awesome blog post here'
---
{% endhighlight %}

- Then in your `index.html`, modify the part that's rendering the excerpt to:

{% highlight html %}
{% raw %}{% if post.excerpt %}{% endraw %}
{% raw %}{{ post.excerpt }}{% endraw %}
{% raw %}{% else %}{% endraw %}
{% raw %}{{ post.content | truncatewords:30 }}{% endraw %}
{% raw %}{% endif %}{% endraw %}
{% endhighlight %}

This will render the excerpt where it's present, and fall back to the first 30 words where it's not.

Hope this helps! Back to blogging like a hacker!
