---
layout: post
title: Converting between Timezones in Python
author: Fred Muya
---

We all know what a joy it is working with different timezones in our applications, especially when we need to convert different times between them.

To make it even more pleasurable, I created this function some time back to help me with these.

It depends on the `dateutil` library to parse the dates.

{% highlight python %}
import datetime
from dateutil import parser as date_parser
from dateutil import tz
def convert_timestamp_timezone(timestamp, from_tz="UTC", to_tz="UTC"):
    """
    function to convert a string timestamp between timezones
    @timestamp - A string timestamp (dateutil.parser will be used to parse)
    @from_tz - A string, the current timezone as a string.
    @to_tz - A string, the timezone to convert the time to.
    Refer to: http://goo.gl/hmPXML for a list of acceptable TZ strings
    """
    timestamp = date_parser.parse(timestamp)
    from_tz = tz.gettz(from_tz)
    to_tz = tz.gettz(to_tz)
    tz_aware_timestamp = timestamp.replace(tzinfo=from_tz)
    converted_timestamp = tz_aware_timestamp.astimezone(to_tz)
    return converted_timestamp
{% endhighlight %}


Example use:
{% highlight python %}
>>> time_in_tokyo = convert_timestamp_timezone("2015-09-09 09:09:09", "Africa/Nairobi", "Asia/Tokyo")
>>> time_in_tokyo.strftime("%Y-%m-%d %H:%M:%S")
{% endhighlight %}

Make sure you use a timezone string from here: [https://en.wikipedia.org/wiki/List_of_tz_database_time_zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones){:target="_blank"}

Easy to grab gist here: [https://gist.github.com/muya/0054ebb9487f55615daa](https://gist.github.com/muya/0054ebb9487f55615daa)

That's it! Happy coding!
