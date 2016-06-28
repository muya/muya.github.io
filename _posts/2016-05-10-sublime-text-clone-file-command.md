---
layout: post
title: Create Shortcut to Clone a File in Sublime Text 3
author: Fred Muya
---

I'm putting this here for my future self more than anything else (I know, I'm selfish 😉).

Selecting `File -> New View into File` every time you need to open the same file in a new tab in Sublime Text is too much work, solution: create a shortcut to do that.

Add the following to your key bindings file (`Preferences ->  Key Bindings - User`):

{% highlight json %}
[
    // other shortucts
    ...
    { "keys": ["ctrl+shift+d"], "command": "clone_file" }
    ...
    // even more shortcuts
]
{% endhighlight %}

As a plus, the menu on Sublime's Menu bar updates with the new shortcut!

![Sublime Text Clone File Shortcut]({{ site.url }}/images/2016-05-10/st3-clone-shortcut.png)

Happy Coding!
