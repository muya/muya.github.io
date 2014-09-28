---
layout: post
title: Apache Issue after Installing OS X Yosemite
date: '2014-07-26T16:49:00.000+03:00'
author: Fred Muya
tags:
modified_time: '2014-08-06T13:27:03.831+03:00'
excerpt: Solving syntax error in Apache httpd.conf file after upgrading to Yosemite Beta
---

I recently updated to OS X Yosemite, and I was having issues with my Apache
installation. Specifically, this is the error I was getting when i
ran `apachectl -t` (to test Apache configuration):
{% highlight bash %}
httpd: Syntax error on line 527 of /private/etc/apache2/httpd.conf: Syntax error on line 8 of /private/etc/apache2/other/+php-osx.conf: Cannot load /usr/local/php5/libphp5.so into server: dlopen(/usr/local/php5/libphp5.so, 10): Symbol not found: _unixd_config Referenced from: /usr/local/php5/libphp5.so Expected in: /usr/sbin/httpd in /usr/local/php5/libphp5.so
{% endhighlight %}

And according to this Github issue [here](https://github.com/liip/php-osx/issues/117),
it seems to be something to do with renaming of a file in Apache 2.4 (which Yosemite ships with).
I followed the suggestion there, and you're required to uncomment this line
(remove the '#' at the beginning of the line) in your `/etc/apache2/httpd.conf` file:
{% highlight bash %}
LoadModule php5_module libexec/apache2/libphp5.so
{% endhighlight %}

Make sure you run the command as root, then restart apache. All should work great now!
