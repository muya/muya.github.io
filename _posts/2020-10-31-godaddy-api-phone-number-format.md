---
layout: post
title: GoDaddy API Required Phone Number Format
date: '2020-10-31'
author: Fred Muya
tags:
- godaddy
- godaddy-api
---

GoDaddy have an [API](https://developer.godaddy.com/getstarted) that allows one to perform different actions on their platform.

Some of the API endpoints, e.g. the [Register Domain endpoint](https://developer.godaddy.com/doc/endpoint/domains#/v1/purchase), require an "Address" payload. However, the API documentation doesn't provide any examples.

I was having trouble with the required phone number format, but I finally figured it out. The expected format is:

{% highlight plaintext %}
"+{country_dialling_code}.{rest of phone number}"

# For example, for a Kenyan phone number (country dialling code: 254):
"+254.725669669"

{% endhighlight %}


A full Address payload may look like this:
{% highlight json %}
{
    "addressMailing": {
        "address1": "631  Marietta Street",
        "address2": "Mad River Addrr",
        "city": "Mad River",
        "country": "KE",
        "postalCode": "00100",
        "state": "Nairobi"
    },
    "email": "abeja@some-mail-client.net",
    "nameFirst": "Ab",
    "nameLast": "Ja",
    "nameMiddle": "e",
    "phone": "+254.725669669"
}
{% endhighlight %}

I'll post tips & tricks to work with the GoDaddy API as I come across problems & quirks.

Happy Coding, and stay safe!
