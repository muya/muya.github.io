---
layout: post
title: Making your Bash Shell More Interesting Using Fortune
author: Fred Muya
---

Every time I open a new shell prompt, a message similar to the one below greets me:
{% highlight bash %}
Weekends were made for programming.
- Karl Lehenbauer
muya@ENCORE ~ 00:27:16 $
{% endhighlight %}

The witty message is provided by fortune, and you too can have something similar.

Using your package manager, install `fortune`, e.g. if you use Homebrew on OS X, run:
{% highlight bash %}
$ brew install fortune
{% endhighlight %}

Test installation by running fortune on your command line:
{% highlight bash %}
$ fortune
Universe, n.:
    The problem.
{% endhighlight %}

I suggest you skim through `man fortune` to get your head wrapped around how fortune works. I'll just highlight a few options:

* `-a` - Will select a random epigram that could be or not be offensive
* `-o` - Will select a random epigram that's offensive

Running fortune without a flag just selects a random, non-offensive epigram.

Quoting the man page, use `-a` and `-o` if and only if you believe, deep in your heart, that you are willing to be offended.

Next, adding it to your bash. Edit your bash prompt config file. For some, it's usually `~/.bash_profile`, others, `~/.bashrc`

Add this to the end of the file:
{% highlight bash %}
fortune
# fortune -o if you are willing to be offended
# fortune -a if you are willing to be offended by surprise  ;)
{% endhighlight %}

This means that the command will be run every time the config file is run (which is usually the case when you open a new bash shell)

Save and exit the file, then run `soure ~/.bash_profile`

Now any time you open a shell, a random epigram will be displayed.

Have fun!
