---
layout: post
title: Fetching Redis INFO using Redigo
date: '2020-10-18'
author: Fred Muya
tags:
- golang
- redis
- go
- redigo
---

[Redigo](https://github.com/gomodule/redigo) is a [Go](https://go.dev/) client for the [Redis](https://redis.io/) database.

It may be useful to be able to fetch & use Redis Server information within your Go program. The library currently doesn't offer an explicit method to query for the information returned by the [Redis INFO command](https://redis.io/commands/info).

This post outlines how to query Redis INFO when using the [redigo library](https://github.com/gomodule/redigo) in a GoLang project. As an example, we'll be fetching the `uptime_in_seconds` field from the INFO response.

First, ensure you have a redis connection to work with. In our example, this connection is named: `conn`

Send the `"INFO"` command, and pass it to the [`String`](https://godoc.org/github.com/gomodule/redigo/redis#hdr-Reply_Helpers) helper method provided by redigo.

{% highlight go %}
infoResponse, getInfoErr := redis.String(conn.Do("INFO"))

// Expected value: string blob with result of "INFO" command
{% endhighlight %}

Next, we'll use a regular expression to define the field we want to extract from the INFO response. In our case, `uptime_in_seconds` field, which is: "Number of seconds since Redis server start". Given that the value is expected to be an integer, our regex can explicitly search using a number.

{% highlight go %}
findUptimeValueRegex := regexp.MustCompile(`uptime_in_seconds:(\d+)`)

// Expected value: uptime_in_seconds:403
{% endhighlight %}

And then, we'll split the string using `":"` to get the uptime value

{% highlight go %}
uptimeDetails := strings.Split(findUptimeValueRegex.FindString(infoResponse), ":")

// Handle error in case the response was empty
if len(uptimeDetails) < 2{
    // return error
}
{% endhighlight %}

Finally, we convert the uptime value to an integer
{% highlight go %}
strconv.ParseInt(uptimeDetails[1], 10, 64)
{% endhighlight %}

Putting it all together, we can have a function that accepts the Redis connection, and return the uptime value:
{% highlight go %}
func getRedisServerUptime(conn redis.Conn) (int64, error) {
    // Send INFO command
    infoResponse, getInfoErr := redis.String(conn.Do("INFO"))
    if getInfoErr != nil {
        return 0, getInfoErr
    }

    // Define regex to extract required value
    findUptimeValueRegex := regexp.MustCompile(`uptime_in_seconds:(\d+)`)

    // Parse value
    uptimeDetails := strings.Split(findUptimeValueRegex.FindString(infoResponse), ":")

    // Handle error in case the response was empty
    if len(uptimeDetails) < 2{
        return 0, errors.New("uptime not found in redis server INFO")
    }

    // Convert value to an integer
    strconv.ParseInt(uptimeDetails[1], 10, 64)
}
{% endhighlight %}

This approach can be used to extract any field from the INFO result.


I'll see if I can propose this as an addition to the library, since it could come in handy.

Happy Coding, and stay safe!


