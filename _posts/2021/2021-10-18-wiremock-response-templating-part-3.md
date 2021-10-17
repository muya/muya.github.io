---
layout: post
title: 'Using Wiremock to Mock API Responses: Part 3 - Response Templating using Handlebars Helpers'
date: '2021-10-18'
author: Fred Muya
excerpt: Making use of Wiremock's Response Templating for dynamic responses using Handlebars helpers
tags:
- wiremock
- api
- mock
- response-templating
- tutorial
- wiremock tutorial
- wiremock handlebars
- handlebars
---

In the previous 2 parts of this series, we:
- [Part 1]({{ site.url }}/wiremock-response-templating/): introduced Wiremock, and showed how it could be used to mock a simple response from a dummy "Songs API"
- [Part 2]({{ site.url }}/wiremock-response-templating-part-2/): showed how to return dynamic responses using details from the incoming request (i.e. request parameters)

In this part, we'll continue with the topic of dynamic responses, but this time making use of Handlebars Helpers, as well as other built-in helpers, to generate even more dynamic responses.

This post is full of practical examples to ease understanding of Handlebars usage.

**Note:** Before proceeding, please ensure you've enabled response templating in Wiremock; this has been outlined in [Part 2]({{ site.url }}/wiremock-response-templating-part-2/) of the series.

Let's take another look at our Songs API response which we're using for a demo:

{% highlight json %}

{
    "metadata": {
        "searchQueryParam": "liked",
        "requestMethodType": "GET"
    },
    "data": [
        {
            "id": "8efe58ee-a5cf-4926-bcd2-92c7e2ec82ba",
            "songInfo": "Jam Now, Simmer Down - Blinky Bill",
            "likedOn": "2021-06-19",
            "listenCount": 34
        },
        {
            "id": "724d520f-15b0-4c2f-9999-a28b2531195c",
            "songInfo": "Dunia Ina Mambo - Just a Band",
            "likedOn": "2020-12-04",
            "listenCount": 182
        }
    ]
}



{% endhighlight %}

Let's say, for example, that we want the `songInfo` to be in `UPPER CASE`; i.e. have the API return:

{% highlight json %}


{
    "metadata": {
        "requestMethodType": "GET"
    },
    "data": [
        {
            "id": "8efe58ee-a5cf-4926-bcd2-92c7e2ec82ba",
            "songInfo": "JAM NOW, SIMMER DOWN - BLINKY BILL",
            "likedOn": "2021-06-19",
            "listenCount": 34
        },
        {
            "id": "724d520f-15b0-4c2f-9999-a28b2531195c",
            "songInfo": "DUNIA INA MAMBO - JUST A BAND",
            "likedOn": "2020-12-04",
            "listenCount": 182
        }
    ]
}


{% endhighlight %}

We can make use of the `upper` helper which is enabled by Handlebars:

{% highlight json %}

{
    "metadata": {
        ...
    },
    "request": {
        "method": "ANY",
        "urlPattern": "/api/songs\\?dynamic=true"
    },
    "response": {
        ...
        "jsonBody": {
            "metadata": {
                "searchQueryParam": "{% raw %}{{ request.query.search.0 }}{% endraw %}",
                "requestMethodType": "{% raw %}{{ request.method }}{% endraw %}"
            },
            "data": [
                {
                    "id": "8efe58ee-a5cf-4926-bcd2-92c7e2ec82ba",
                    "songInfo": "{% raw %}{{ upper 'Jam Now, Simmer Down - Blinky Bill' }}{% endraw %}",
                    "likedOn": "2021-06-19",
                    "listenCount": 34
                },
                {
                    "id": "724d520f-15b0-4c2f-9999-a28b2531195c",
                    "songInfo": "{% raw %}{{ upper 'Dunia Ina Mambo - Just a Band' }}{% endraw %}",
                    "likedOn": "2020-12-04",
                    "listenCount": 182
                }
            ]
        }
    }
}

{% endhighlight %}

**Note:** we're using a new URL pattern (`dynamic=true`) to allow differentiation between this and the previous version in Part 2.

Handlebars also provides a helper to get the current timestamp.

For example, let's say we wanted to add a `"requestedAt"` field to the response metadata. For this, we can use the `now` helper:

{% highlight json %}

{
    "metadata": {
        ...
    },
    "request": {
       ...
    },
    "response": {
        ...
        "jsonBody": {
            "metadata": {
                ...
                "requestedAt": "{% raw %}{{ now }}{% endraw %}"
            },
            "data": [
                {
                    ...
                },
                {
                    ...
                }
            ]
        }
    }
}

{% endhighlight %}


The response would be:

{% highlight json %}
{
  "metadata": {
    "requestMethodType": "GET",
    "requestedAt": "2021-09-26T20:31:32Z"
  },
  "data": [
    {
      "id": "8efe58ee-a5cf-4926-bcd2-92c7e2ec82ba",
      "songInfo": "JAM NOW, SIMMER DOWN - BLINKY BILL",
      "likedOn": "Monday, June 21, 2021",
      "listenCount": 34
    },
    {
      "id": "724d520f-15b0-4c2f-9999-a28b2531195c",
      "songInfo": "DUNIA INA MAMBO - JUST A BAND",
      "likedOn": "Saturday, December 4, 2021",
      "listenCount": 182
    }
  ]
}
{% endhighlight %}

Handlebars also provides date helpers to allow for formatting of timestamps.

For example, let's say we wanted the `likedOn` date format to be something like 'Saturday, 4th December, 2021'. To do this, we can use the `dateFormat` helper:


{% highlight json %}

{
    "metadata": {
        ...
    },
    "request": {
        "method": "ANY",
        "urlPattern": "/api/songs\\?dynamic=true"
    },
    "response": {
        ...
        "jsonBody": {
            ...
            "data": [
                {
                    "id": "8efe58ee-a5cf-4926-bcd2-92c7e2ec82ba",
                    "songInfo": "{% raw %}{{ upper 'Jam Now, Simmer Down - Blinky Bill' }}{% endraw %}",
                    "likedOn": "{% raw %}{{ dateFormat '2021-06-19' full }}{% endraw %}",
                    "listenCount": 34
                },
                {
                    "id": "724d520f-15b0-4c2f-9999-a28b2531195c",
                    "songInfo": "{% raw %}{{ upper 'Dunia Ina Mambo - Just a Band' }}{% endraw %}",
                    "likedOn": "{% raw %}{{ dateFormat '2020-12-04' full }}{% endraw %}",
                    "listenCount": 182
                }
            ]
        }
    }
}


{% endhighlight %}


This would produce:

{% highlight json %}
{
  "metadata": {
    "requestMethodType": "GET",
    "requestedAt": "2021-10-17T15:27:32Z"
  },
  "data": [
    {
      "id": "8efe58ee-a5cf-4926-bcd2-92c7e2ec82ba",
      "songInfo": "JAM NOW, SIMMER DOWN - BLINKY BILL",
      "likedOn": "Monday, June 21, 2021",
      "listenCount": 34
    },
    {
      "id": "724d520f-15b0-4c2f-9999-a28b2531195c",
      "songInfo": "DUNIA INA MAMBO - JUST A BAND",
      "likedOn": "Saturday, December 4, 2021",
      "listenCount": 182
    }
  ]
}
{% endhighlight %}


`upper`, `now`, `dateFormat` are only a subset of all the available helpers.

For a full list, check out the [`StringHelpers.java`](https://github.com/jknack/handlebars.java/blob/master/handlebars/src/main/java/com/github/jknack/handlebars/helper/StringHelpers.java) file on the Handlebars repo.

Additionally, I've provided a "kitchen sink" API that demonstrates all of these helpers.


{% highlight json %}

{
    "metadata": {
        "description": "Provides a demo API utilizing all the supported Handlebars String Helpers in Wiremock",
        "blogPost": "https://blog.muya.co.ke/wiremock-response-templating-part-3/"
    },
    "request": {
        "method": "GET",
        "urlPattern": "/handlebarsKitchenSink"
    },
    "response": {
        "status": 200,
        "headers": {
            "Content-Type": "application/json"
        },
        "jsonBody": {
            "abbreviate": "{% raw %}{{ abbreviate 'Truncate long sentence up to # of characters and add ellipses' 28 }}{% endraw %}",
            "capitalize": {
                "capitalize first letter of all words": "{% raw %}{{ capitalize 'ONLY first letter capitalized' }}{% endraw %}",
                "capitalize first letter of all words AND lower case other characters": "{% raw %}{{ capitalize 'FULLY first letter capitalized' fully=true }}{% endraw %}"
            },
            "capitalizeFirst": "{% raw %}{{ capitalize 'only first string' }}{% endraw %}",
            "center": {
                "center a string": "{% raw %}{{ center 'centerAStringWithEmptySpaces' size=40 }}{% endraw %}",
                "center string with padding": "{% raw %}{{ center 'centerAStringWithAsteriskPadding' size=40 pad='*'  }}{% endraw %}"
            },
            "cut": {
                "remove number 7 from a string": "{% raw %}{{ cut 'string 7 with 7 number 7s' '7' }}{% endraw %}",
                "remove spaces from a string": "{% raw %}{{ cut 'string with spaces' }}{% endraw %}"
            },
            "dateFormat": {
                "display current date and time in custom format": "{% raw %}{{ dateFormat (now)  format='yyyy-MM-dd HH:mm:ss'}}{% endraw %}",
                "display current date in full format": "{% raw %}{{ dateFormat (now) format='full' }}{% endraw %}",
                "display current date in medium format": "{% raw %}{{ dateFormat (now) format='medium' }}{% endraw %}",
                "display current date in short format": "{% raw %}{{ dateFormat (now) format='short' }}{% endraw %}",
                "parse timestamp from value in specific format, and display in full format": "{% raw %}{{ dateFormat (parseDate '2021-06-21' format='yyyy-MM-dd') format='full' }}{% endraw %}"
            },
            "defaultIfEmpty": {
                "set NOTHING as value provided value is empty ": "{% raw %}{{ defaultIfEmpty '' 'NOTHING'  }}{% endraw %}",
                "set empty string if value provided is empty": "{% raw %}{{ defaultIfEmpty ''  }}{% endraw %}"
            },
            "join": {
                "join a list of items with custom joiner (last item in list is considered the joiner)": "{% raw %}{{ join 'a' 'b' '-' }}{% endraw %}",
                "join a list of items with custom joiner and prefix": "{% raw %}{{ join 'a' 'b' '-' prefix='[' }}{% endraw %}",
                "join a list of items with custom joiner and suffix": "{% raw %}{{ join 'a' 'b' '-' suffix=']' }}{% endraw %}",
                "join a list of items with custom joiner, prefix and suffix": "{% raw %}{{ join 'a' 'b' '-' prefix='[' suffix=']' }}{% endraw %}"
            },
            "ljust": {
                "left align a given string in a 30 width space": "{% raw %}{{ ljust 'left aligned' size=30 }}{% endraw %}",
                "left align a given string in a 30 width space with padding": "{% raw %}{{ ljust 'left aligned' size=30 pad='*' }}{% endraw %}"
            },
            "lower": "{% raw %}{{ lower 'CHANGE VALUE TO LOWER CASE' }}{% endraw %}",
            "now": {
                "display current date time": "{% raw %}{{ now }}{% endraw %}",
                "display current date time in custom format": "{% raw %}{{ now format='yyyy-MM-dd HH:mm:ss.SSSSSS' tz='Africa/Nairobi' }}{% endraw %}"
            },
            "numberFormat": {
                "format number in currency format": "{% raw %}{{ numberFormat 30 'currency' }}{% endraw %}",
                "format number in currency format with locale": "{% raw %}{{ numberFormat 30 'currency' 'fr'}}{% endraw %}",
                "format number in custom decimal format": "{% raw %}{{ numberFormat 3000000 '#,###,##0.000' }}{% endraw %}",
                "format number in integer format": "{% raw %}{{ numberFormat 30 'integer' }}{% endraw %}",
                "format number in percent format": "{% raw %}{{ numberFormat 30 'percent' }}{% endraw %}",
                "format number with defined maximum integer and fraction digits": "{% raw %}{{ numberFormat 4542.3733 maximumFractionDigits=3 maximumIntegerDigits=2 }}{% endraw %}",
                "format number with defined minimum integer and fraction digits": "{% raw %}{{ numberFormat 0.37 minimumFractionDigits=3 minimumIntegerDigits=2 }}{% endraw %}"
            },
            "replace": "{% raw %}{{ replace 'Replaces placeholder with another string' 'another string' 'provided replacement' }}{% endraw %}",
            "rjust": {
                "right align a given string in a 30 width space": "{% raw %}{{ rjust 'right aligned' size=30 }}{% endraw %}",
                "right align a given string in a 30 width space with padding": "{% raw %}{{ rjust 'right aligned' size=30 pad='*' }}{% endraw %}"
            },
            "slugify": "{% raw %}{{ slugify 'Creates A Slug Useful For Blog Post URLs' }}{% endraw %}",
            "stringFormat": {
                "boolean": "{% raw %}{{ stringFormat 'isSet: %b isNotSet: %b' 'yes' null }}{% endraw %}",
                "string": "{% raw %}{{ stringFormat 'applies %s formatting capabilities. All Java formatting options supported.' 'string' }}{% endraw %}"
            },
            "stripTags": "{% raw %}{{ stripTags '<span>Removes all (X)HTML tags</span>' }}{% endraw %}",
            "substring": {
                "substring from 3rd (exclusive) to 7th character": "{% raw %}{{ substring '0123456789' 3 7 }}{% endraw %}",
                "substring from 5th character (exclusive)": "{% raw %}{{ substring '0123456789' 5 }}{% endraw %}"
            },
            "upper": "{% raw %}{{ upper 'change value to upper case' }}{% endraw %}",
            "yesno": {
                "set true/false/null to yes/no/maybe": "{% raw %}{{ yesno true }} | {{ yesno false }} | {{ yesno null }}{% endraw %}",
                "set true/false/null to ndio/hapana/labda (Swahili)": "{% raw %}{{ yesno true yes='ndio' }} | {{ yesno false no='hapana'  }} | {{ yesno null maybe='labda' }}{% endraw %}",
                "set true/false/null to sí/no/quizás (Spanish)": "{% raw %}{{ yesno true yes='sí' }} | {{ yesno false no='no'  }} | {{ yesno null maybe='quizás' }}{% endraw %}"
            }
        }
    }
}



{% endhighlight %}


NB: Some special notes on exclusions from the kitchen sink:

- There seems to be a bug with handling of JSON that has special characters (e.g. newline characters) between Handlebars and Wiremock. As a result, the following helpers have not been added to the kitchen sink: `wordWrap`
- `numberFormat` has flags for `parseIntegerOnly` and `roundingMode`, but these don't seem to be in use


We'll use these exclusions as use-cases later in this series to build custom extensions for Wiremock.


## Summary
In this part of the series, we dove deeper into Response Templating in Wiremock, seeing how to make use of Handlebars Helpers to enable us to generate even more dynamic responses.

We also showed how (almost) all of the string helpers can be used via a "Kitchen Sink API" (this is available on this series' accompanying repository on GitHub: [`wiremock-docker-demo`](https://github.com/muya/wiremock-docker-demo)).

## References
Shout out to [rodolpheche](https://github.com/rodolpheche/wiremock-docker) for building and maintaining the Wiremock docker image.

All the code used in this blog post is available in this repository: [wiremock-docker-demo](https://github.com/muya/wiremock-docker-demo)

## Coming in Part 4: Using Custom Extensions in Wiremock
In the next part, we'll see how to load custom extensions to Wiremock, for when we want to enable very customized functionality in our mocked APIs.

We'll use the custom extensions to go around some of the exclusions noted above.


Until then, happy coding, and stay safe!
