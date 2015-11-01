---
layout: post
title: Generating BitStamp Signature in PHP
author: Fred Muya
---

This post will describe how to create the signature required by the BitStamp API for private calls to the API.

In the BitStamp [docs](https://www.bitstamp.net/api/){:target="_blank"}, they've shown how to do this in Python:

{% highlight python %}
import hmac, hashlib
message = nonce + customer_id + api_key
signature = hmac.new(API_SECRET, msg=message, digestmod=hashlib.sha256).hexdigest().upper()
{% endhighlight %}

In PHP, this is made possible by the [`hash_hmac`](http://php.net/manual/en/function.hash-hmac.php){:target="_blank"} function.

This is the function I'm using to do this in PHP

{% highlight php %}
<?php

/**
 * Generates the signature required by Bitstamp API
 */
function generateBitstampSignature($nonce)
{
    $customerID = '123456';
    $bitstampApiKey = 'apikey';
    $bitstampApiSecret = 'apisecret';
    $message = $nonce . $customerID . $bitstampApiKey;

    return strtoupper(hash_hmac('sha256', $message, $bitstampApiSecret));
}
{% endhighlight %}

I hope this saves someone a few seconds somewhere ;)
