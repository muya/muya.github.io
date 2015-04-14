---
layout: post
title: Allow External Access for Flask Apps on AWS
author: Fred Muya
---

If you are using AWS for your deployment, you may come across a scenario whereby you have to allow access to your server via non-standard ports.

I was recently pulling my hair out when deploying a Flask application that's running on `5001`. It took some time before I came across clear documentation on how to do this.

First of all, make sure that in your code, you've set up the Flask app to
allow external connections, i.e. by setting host to `0.0.0.0`, e.g.

{% highlight python %}

app.run("0.0.0.0")

{% endhighlight %}

If you still can't access the port you need (I use telnet to check for access to a given port), follow the steps below:

- Log in to AWS, and go to your instances list
- Select your server, and click on the **Description** tab that pops up below
- There is a Security Groups option, click on *view rules* to see the ports that have been allowed access
- If the port you require is not there (or is mis-configured), click on the Security Group for your instance (on the Instances table) - It's one of the last columns
- Choose **Inbound** tab in the popup at the bottom, then **Edit**
- The popup allows you to define rules. In my case, I used:
    - Type - Custom TCP Rule
    - Protocol - TCP
    - Port Range - 5001
    - Source - Anywhere (tailor as you wish)
- Click on Save

That's it!

Confirm that it works by doing a telnet to your server, e.g.:
{% highlight bash %}
$ telnet muya.co.ke 5001
{% endhighlight %}
