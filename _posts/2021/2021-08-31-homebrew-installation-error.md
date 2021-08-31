---
layout: post
title: 'gobuffalo/homebrew-tap is not a valid repository name'
date: '2021-08-31'
author: Fred Muya
excerpt: Resolve Git issue when installing packages via Homebrew
tags:
- homebrew
- brew
- git
- ssh
- https
- macos
- mac
---

Short post this time round - hope you're all staying safe.

A few weeks back I was having a problem installing [Go Buffalo](https://gobuffalo.io/en/) (GoLang web framework) using Homebrew on macOS, and faced the error below:

{% highlight plaintext %}

==> Tapping gobuffalo/tap
Cloning into '/opt/homebrew/Library/Taps/gobuffalo/homebrew-tap'...
fatal: remote error: 
  /gobuffalo/homebrew-tap is not a valid repository name
  Visit https://support.github.com/ for help
Error: Failure while executing; `git clone https://github.com/gobuffalo/homebrew-tap /opt/homebrew/Library/Taps/gobuffalo/homebrew-tap --origin=origin` exited with 128.

{% endhighlight %}

Turns out it was due to a recent Git SSH config change I'd made.

I'd overriden my global Git config to convert all https requests to GitHub to use SSH; i.e. this is what I had in my `~/.gitconfig` file:

{% highlight ini %}

[url "ssh://git@github.com/"]
    insteadOf = https://github.com

{% endhighlight %}

Not sure why this was causing the failure, but reverting the change allowed Homebrew to work again as expected (I commented this out in the git config file).

Hopefully this helps someone out!

Huge credit to [@tconroy](https://github.com/Homebrew/legacy-homebrew/issues/45278#issuecomment-150692549) whose solution pointed me in the right direction.

Until next time, happy coding, and stay safe!
