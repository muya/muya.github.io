---
layout: post
title: Configuring Multiple Loggers in Python
author: Fred Muya
---

At times you may want to have multiple loggers in your application, e.g. one for a payments module and another for a messaging module. I'm going to explain how to set this up in your application, as well as show some examples for usage.

First, create a file, maybe call it `logger.py`

Import some required libraries

{% highlight python %}
import logging

from logging import FileHandler
from logging import Formatter
{% endhighlight %}

For example, if our app is called `wasted_meerkats`, and we have 2 modules, `payments` and `messaging`, we could set up different loggers as follows:

{% highlight python %}
LOG_FORMAT = (
    "%(asctime)s [%(levelname)s]: %(message)s in %(pathname)s:%(lineno)d")
LOG_LEVEL = logging.INFO

# messaging logger
MESSAGING_LOG_FILE = "/tmp/wasted_meerkats/messaging.log"


messaging_logger = logging.getLogger("wasted_meerkats.messaging")
messaging_logger.setLevel(LOG_LEVEL)
messaging_logger_file_handler = FileHandler(MESSAGING_LOG_FILE)
messaging_logger_file_handler.setLevel(LOG_LEVEL)
messaging_logger_file_handler.setFormatter(Formatter(LOG_FORMAT))
messaging_logger.addHandler(messaging_logger_file_handler)

# payments logger
PAYMENTS_LOG_FILE = "/tmp/wasted_meerkats/payments.log"
payments_logger = logging.getLogger("wasted_meerkats.payments")

payments_logger.setLevel(LOG_LEVEL)
payments_file_handler = FileHandler(PAYMENTS_LOG_FILE)
payments_file_handler.setLevel(LOG_LEVEL)
payments_file_handler.setFormatter(Formatter(LOG_FORMAT))
payments_logger.addHandler(payments_file_handler)
{% endhighlight %}

To test this out, you could create another file, called `wasted_meerkats.py` and put the following code in it:
{% highlight python %}
from logger import messaging_logger
from logger import payments_logger

messaging_logger.info("The meerkats are drunk!")

payments_logger.info("Who knows where they got the money?!")
{% endhighlight %}

Create the log directories above
{% highlight bash %}
$ mkdir /tmp/wasted_meerkats && touch /tmp/wasted_meerkats/messaging.log /tmp/wasted_meerkats/payments.log
{% endhighlight %}

In a separate tab, tail the logs
{% highlight bash %}
tail -f /tmp/wasted_meerkats/*log
{% endhighlight %}

Then test it out!
{% highlight bash %}
$ python wasted_meerkats.py
{% endhighlight %}

The log output should look something like this:
{% highlight bash %}
==> /tmp/wasted_meerkats/messaging.log <==
2015-10-17 20:55:05,942 [INFO]: The meerkats are drunk! in wasted_meerkats.py:4

==> /tmp/wasted_meerkats/payments.log <==
2015-10-17 20:55:05,942 [INFO]: Who knows where they got the money?! in wasted_meerkats.py:6
{% endhighlight %}

Having all the loggers in one file can help organize your project better.

I've put all the code used here as a gist here: [https://gist.github.com/muya/2dff1cd8c5b42f1dabab](https://gist.github.com/muya/2dff1cd8c5b42f1dabab)

Happy coding people!
