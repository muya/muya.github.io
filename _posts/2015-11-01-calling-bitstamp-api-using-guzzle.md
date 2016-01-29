---
layout: post
title: Calling BitStamp API using Guzzle
author: Fred Muya
---

I've recently been working on integrating to the BitStamp API from a Laravel application, and I was having some problems calling it using the [Guzzle PHP HTTP Client](http://guzzle.readthedocs.org/en/latest/){:target="_blank"}.

I'm sharing how I was finally able to make a call to the API.

The API supports 2 types of calls:

* public calls, e.g to get the market ticker
* private calls, e.g. to get your account balance. The private calls require you to pass some extra details required for authentication, i.e. `key`, `nonce` & `signature`.

The signature is generated using an algorithm they've defined in their [docs](https://www.bitstamp.net/api/){:target="_blank"}. I've written a post to demonstrate how to do this in PHP [here]({{ site.url }}/generate-bitstamp-signature-in-php).

I have a function that will handle calling different endpoints on the API, so I just have to send pass the following:

* `endpoint`
* `requestType`
* `queryParams` for `GET` if required,
* a flag, `requiresAuth` to determine if we need to pass additional authentication params, and;
* the `requestBody`, which contains any `POST` parameters required.

{% highlight php %}
<?php
use GuzzleHttp\Client as GuzzleClient;

function callBitStampApi($endpoint, $requestType = 'GET', $queryParams = [], $requiresAuth = false, $requestBody = [])
{
    $bitstampApiClient = new GuzzleClient([
        'base_uri' => 'https://www.bitstamp.net/api/'
    ]);
    $requestParams = [
        'query' => $queryParams
    ];

    $requestParams['form_params'] = $requestBody;

    if ($requiresAuth) {
        // generate nonce
        $mt = preg_split('/ /', microtime());
        $nonce = $mt[1] . substr($mt[0], 2, 6);
        $requestParams['form_params']['key'] = env('BITSTAMP_API_KEY');
        $requestParams['form_params']['nonce'] = $nonce;
        $requestParams['form_params']['signature'] = generateBitstampSignature($nonce);
    }

    $resp = $bitStampApiClient->request($requestType, $endpoint, $requestParams);

    return json_decode($resp);
}
?>
{% endhighlight %}

For calls that require authentication, the `key`, `nonce` & `signature` should be passed as form_params.

One thing to note, if you pass the `endpoint` parameter without a trailing slash, e.g. `ticker` instead of `ticker/`, you'll get an error saying:

{% highlight json %}
{"error": "Missing key, signature and nonce parameters"}
{% endhighlight %}

An example of a call to this function would be:
{% highlight php %}
<?php
// make call to ticker endpoint
$marketTickerData = callBitStampApi('ticker/', 'GET');

// make call to account balance endpoint
$accountBalData = callBitStampApi('balance/', 'POST', [], true);
{% endhighlight %}

That's all for now people! Back to code!
