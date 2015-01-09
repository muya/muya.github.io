---
layout: post
title: Minify JS, CSS files on Save in Sublime Text
date: '2014-06-16T15:19:00.000+03:00'
author: Fred Muya
tags:
modified_time: '2014-06-16T15:19:48.144+03:00'
---

We all know the advantages of using minified files in your web apps, i.e. faster loading, smaller payload, etc, it would make sense if all editors that we use have the feature built in, so that we don't have to worry about it.

I use [Sublime Text](http://www.sublimetext.com/){:target="_blank"} for most of my work, and much as there are plugins to accomplish the minification of JS & CSS files, an essential feature that I haven't come across is minification on Save.

To solve this, I created a simple plugin to accomplish just this. It requires you to have the Sublime-Minifier extension, [https://github.com/bistory/Sublime-Minifier](https://github.com/bistory/Sublime-Minifier){:target="_blank"}

To start using it:

- Grab this gist: [https://gist.github.com/muya/1cd6ce82490d2c0f2e0a](https://gist.github.com/muya/1cd6ce82490d2c0f2e0a){:target="_blank"}
- In Sublime-Text, click on `Tools->New Plugin...`
- Replace the file content with the code from the gist
- Hit Save!

Here is the Gist:
{% highlight python %}
import sublime, sublime_plugin

class MinifyOnSave(sublime_plugin.EventListener):
  def on_post_save(self, view):
    file_types_to_minify = ['js', 'css']
    filenameParts = view.file_name().split('.')
    if filenameParts[len(filenameParts) - 1] in file_types_to_minify:
      view.run_command('minify_to_file')
{% endhighlight %}
Now, when you save any of your JS & CSS files, a corresponding file (.min.js/.min.css) will be created/updated in the directory.

PS: Atom Editor has this feature for one of the plugins, still not comfortable with it, so Sublime it is!
