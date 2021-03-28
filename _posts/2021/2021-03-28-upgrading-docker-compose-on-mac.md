---
layout: post
title: Upgrading docker-compose on Mac (and using profiles)
date: '2021-03-28'
author: Fred Muya
excerpt: Guide on how to get the latest version of docker-compose when using Docker for Mac (and a note on
    docker-compose profiles)
tags:
- docker
- docker-compose
- mac
- macos
- profiles
- docker-compose profiles
---

A few weeks back, I needed to use the "[profiles](https://docs.docker.com/compose/release-notes/#1280)" feature that
was annouced as part of docker-compose Release 1.28.0. However, I was running an older version of
docker-compose.

Given that the only way to run docker upgrades on MacOS is via Docker for Mac, I thought upgrading my Docker for
Mac app would get me the latest version.

Unfortunately, the Docker for Mac app was showing "No updates available".

To resolve this, I had to manually download the latest installer for Docker for Mac from the
[Release Page](https://docs.docker.com/docker-for-mac/release-notes/), and run it.

Upon installation completion, the docker-compose version was now at 1.28.5 (latest version), and I could use the
"profiles" feature.

(During re-installation, if prompted to "Replace" the existing app, accept it. Data will not be overwritten).


### A Note about docker-compose Profiles
The new [profiles](https://docs.docker.com/compose/profiles/) feature is a great improvement to docker-compose.
It enables "tagging" of different services in a docker-compose file, allowing you to only bring up specific services when you run `docker-compose up`.

Profiles come in handy in applications that may have multiple sub-modules which have different dependencies.

For example, take the docker-compose file below:

{% highlight yaml %}

---
version: "3.9"
services:
  cassandra:
    image: cassandra:latest
    # ... other Cassandra configs
    profiles:
      - module-a
  kafka:
    image: confluentinc/cp-kafka:latest
    # ... other Kafka configs
    profiles:
      - module-b
  mariadb:
    image: mariadb:latest
    # ... other MariaDB configs
    profiles:
      - module-c

{% endhighlight %}

It defines a configuration for a multi-module application setup, whereby:
- Module A needs Cassandra
- Module B needs Kafka
- Module C needs Maria DB

When doing development on Module C, it may not be necessary to bring up the Cassandra & Kafka services; therefore, we can bring up only "mariadb" by adding the `--profile` flag:

{% highlight shell %}

$ docker-compose up --profile module-c

{% endhighlight %}

Upon execution, only the `"mariadb"` service will be brought up, allowing you to preserve resources that might have
been potentially lying idle.


The docker-compose documentation has an excellent guide on using Profles, check it out to learn more:
[Using profiles with Compose](https://docs.docker.com/compose/profiles/)


In this blog post, we demonstrated how to upgrade your docker-compose on MacOS in case you're unable to see an
update for Docker for Mac.<br>
We also highlighted the docker-compose "profiles" feature, and how you can use it to
selectively bring up services.

Happy coding, and stay safe!
