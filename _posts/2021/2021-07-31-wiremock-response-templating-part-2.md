---
layout: post
title: 'Using Wiremock to Mock API Responses: Part 2 - Response Templating using Request Parameters'
date: '2021-07-31'
author: Fred Muya
excerpt: Making use of Wiremock's Response Templating for dynamic responses using request parameters
tags:
- wiremock
- api
- mock
- response-templating
- tutorial
- wiremock tutorial
- wiremock request parameters
---

In the [first post]({{ site.url }}/wiremock-response-templating/), we introduced Wiremock, and showed how it could be
used to mock a simple response from a dummy "Songs API".

![Liked Songs API Response]({{ site.url }}/images/2021-07-31/liked-songs-api-response.png)

In this part, we'll show how to have the mocked API return dynamic responses.

Having dynamic responses from your mocked API allows you to make it more similar to a real API returning actual data.

In our "Songs API", we can make the different fields dynamic in the following ways:

- `id` - have uniquely generated UUIDs for each request
- `songInfo` - pick one from a list of songs. This could allow you to test for songs with different characters in them.
- `likedOn` - have a random date for when a song was liked
- `listenCount` - have a random number of listens

Using ["Response Templating"](http://wiremock.org/docs/response-templating/), Wiremock allows us to define dynamic responses.

## Enabling Response Templating in Wiremock
To utilize the Response Templating functionality, it has to be explicitly enabled.

It can be enabled on a per-mapping basis, or across the board for all mappings.

For our case, we'll enable it globally, using the `--global-response-templating` flag.

Since we're using Docker, we'll add this to the `docker-compose` file:

{% highlight yaml %}

version: '3.9'
services:
  wiremock:
    image: rodolpheche/wiremock
    volumes:
      - $PWD/wiremock:/home/wiremock
    ports:
      - 8080:8080
    command:
      - --verbose
      - --global-response-templating



{% endhighlight %}


If you'd like to enable it on a per-mapping basis, use the `--local-response-templating` flag instead, and add the
following snippet to each of your mapping files:

{% highlight json %}

{
    "request": {
        ...
    },
    "response": {
        ...
        "transformers": ["response-template"]
        ...
    }
}

{% endhighlight %}

You'll need to restart the Docker container after making this change.


## Configuring Dynamic Responses

Dynamic responses can be created in 3 ways, depending on your use case:

1. Using request parameters
2. Using [Handlebars Helpers](https://handlebarsjs.com/)
3. Using custom transformer properties

In this series, we'll review options 1 & 2.

### Using Request Parameters
Wiremock allows you to use the incoming request details as part of the mocked response.

For example, let's say we wanted to update our Songs API response to include some metadata about the incoming request.

We've updated the request to include query parameters to demo the functionality.

We'd like to include the search parameters and the request method type in the metadata.


###### Sample Request
{% highlight plaintext %}

GET /api/songs?search=liked

{% endhighlight %}

###### Expected Response
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

To have Wiremock return the response above, this is what the mapping will look like:

{% highlight json %}

{
    "metadata": {
        "description": "Demonstrates Wiremock's ability to simulate an API response with templated request parameters",
        "blogPost": "https://blog.muya.co.ke/wiremock-response-templating-part-2/"
    },
    "request": {
        "method": "ANY",
        "urlPattern": "/api/songs\\?search=([a-z]*)"
    },
    "response": {
        "status": 200,
        "headers": {
            "Content-Type": "application/json"
        },
        "jsonBody": {
            "metadata": {
                "searchQueryParam": "{% raw %}{{ request.query.search.0 }}{% endraw %}",
                "requestMethodType": "{% raw %}{{ request.method }}{% endraw %}"
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
    }
}


{% endhighlight %}


There are a lot of changes here, so let's go through them:

#### Stub Metadata
We've added a `metadata` key on the JSON object.

This is referred to as the Stub Metadata, and it allows one to add any information they need to the response template,
for example, to help with documentation.

It also enables some advanced querying capabilities against Wiremock.

Note that any details can be added to the `metadata`.

In our case, we've added the `description` and `blogPost` keys, to provide additional context about the mapping.

If you'd like, you can use `desc` instead of `description`, or you could add a `mockedService` key to provide info on
which service is being mocked by the given stub.

Read more about it in the docs: [Stub Metadata](http://wiremock.org/docs/stub-metadata/).


#### Request Method
We've set the request method to `"ANY"`.

The `"ANY"` request method type enables this mock to accept different types of HTTP request methods (`POST`, `GET`,
`PUT`, `PATCH`, etc).

We've done this to allow us to demonstrate the `requestMethodType` in the response. When we make a request to the
API, we can use any HTTP method.

This can be useful in cases where you want to mock the same response for different method types.

#### Request URL Pattern
We've used `urlPattern` instead of `urlPath`. This showcases another powerful aspect of Wiremock: URL matching based
on Regular Expressions (Regex).

In our case, we'd like to mock an API that accepts any value for the `search` query parameter.

The regex set in `urlPattern` allows us to match and extract the values passed.

Read more about regex pattern matching in the [Wiremock Docs: Request Matching](http://wiremock.org/docs/request-matching/).


#### Reference Request Parameters
With all the changes in place, we're now able to use request parameters in our response.

{% highlight json %}

{
    ...
    "response": {
        ...
        "jsonBody": {
            "metadata": {
                "searchQueryParam": "{% raw %}{{ request.query.search.0 }}{% endraw %}",
                "requestMethodType": "{% raw %}{{ request.method }}{% endraw %}"
            },
            ...
        }
    }
}


{% endhighlight %}


In this example, we're extracting the value of the `search` query parameter, and setting it in `metadata.searchQueryParam`.

We're also extracting the request HTTP method, and setting it in `metadata.requestMethodType`.

If you're following along:
- Copy the full payload above, and add it to a new file under the `mappings` directory; you can name the file `likedSongs-RequestParameters.json`
- Reset wiremock so that it loads the new settings (remember, this can be done by restarting the Docker container, or
by sending a "reset" command to Wiremock; see the first part of this series for details)

Now, you can make an API request as follows, and the response will include the newly added metadata:

{% highlight plaintext %}

# POST request with query param "liked" (response truncated for brevity)
curl --request POST --url 'http://localhost:8080/api/songs?search=liked'

{
  "metadata": {
    "searchQueryParam": "liked",
    "requestMethodType": "POST"
  },
  "data": [
    ...
  ]
}

# GET request with query param "saved" (response truncated for brevity)
curl --request GET --url 'http://localhost:8080/api/songs?search=saved

{
  "metadata": {
    "searchQueryParam": "saved",
    "requestMethodType": "GET"
  },
  "data": [
    ...
  ]
}

{% endhighlight %}

With this mapping, we're able to read values from the request, and use them in our response.

There are more request parameters that are accessible, and you can view them all in the docs: [Wiremock Response Templating: The request model](http://wiremock.org/docs/response-templating/).

## Summary
In this part of the series, we introduced Response Templating in Wiremock. We showed how to enable it, and how to use values from the request as part of the response body.

We also introduced "Stub Metadata", which can be used to provide additional information about a mapping.

Finally, we showed how to perform URL matching based on Regular Expressions, providing flexibility on how you route incoming requests.


## References
Shout out to [rodolpheche](https://github.com/rodolpheche/wiremock-docker) for building and maintaining the Wiremock docker image.

All the code used in this blog post is available in this repository: [wiremock-docker-demo](https://github.com/muya/wiremock-docker-demo)



## Part 3: Dynamic Responses using Handlebars Helpers
Read [Part 3]({{ site.url }}/wiremock-response-templating-part-3/) of this series to learn how to use [Handlebars Helpers](https://handlebarsjs.com/) to allow us to generate dynamic data for our responses.


Until then, happy coding, and stay safe!
