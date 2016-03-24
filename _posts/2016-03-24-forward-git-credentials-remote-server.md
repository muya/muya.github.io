---
layout: post
title: Make SSH Environment available after sudo-ing to another User
author: Fred Muya
---

If you have 2FA enabled on Github (which you should have), you have to use the SSH method to do git operations related to your repo. This becomes a pain when you're doing a lot of remote work (that you haven't yet automated); because you either have to generate a unique password that you can use during the deploy, or switch off 2FA.

Luckily, GitHub has provided an excellent guide on how to enable SSH Agent Forwarding: https://developer.github.com/guides/using-ssh-agent-forwarding/

However, at times you need to work as root or a different user from the one you logged in remotely as.

To make sure agent forwarding works even after sudo-_ing_ to another user, you need to keep the `SSH_AUTH_SOCK` variable available.

On Ubuntu, you do this by adding the following to your `/etc/sudoers` file:

{% highlight bash %}
Defaults        env_keep+=SSH_AUTH_SOCK
{% endhighlight %}

**Make sure you use `visudo` to edit the file, so that it's validated before saving; otherwise you risk losing superpowers on your server 😂**

Once you do this, log in & out again, and you should be able to access your Github repos using your key. You can do the test given by Github:

{% highlight bash %}
ssh -T git@github.com
{% endhighlight %}

You'll get a response like the one below:
{% highlight bash %}
Hi {github username}! You've successfully authenticated, but GitHub does not provide shell access.
{% endhighlight %}

Happy Coding!
