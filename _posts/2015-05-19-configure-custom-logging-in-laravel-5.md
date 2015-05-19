---
layout: post
title: Configuring Custom Logging in Laravel 5
author: Fred Muya
---

I recently started working on a Laravel 5 project, and I'd like to share how I set up my custom file logging.

I pooled the information below from several sources (acknowledged at the bottom of the article).

I organize my log files per project, under a directory, e.g. `/var/log/applications/app_name/app_name.log`, etc.

Below is a sample log that I'd like to accomplish:

![Sample Log Format]({{ site.url }}/images/2015-05-19/log-extract.png)

The format is similar to:
```
date log_type channel message extra_params
```

Custom Logging in Laravel 5 can be achieved by overriding the `Illuminate\Foundation\Bootstrap\ConfigureLogging` class.

To get started, create a file under the `/bootstrap` folder and name it: `ConfigureLogging.php`

Add the code below to the file:
{% highlight php %}
<?php namespace Bootstrap;

use Monolog\Logger as Monolog;
use Monolog\Formatter\LineFormatter;
use Illuminate\Log\Writer;
use Illuminate\Contracts\Foundation\Application;
use Illuminate\Foundation\Bootstrap\ConfigureLogging as BaseConfigureLogging;
use Monolog\Handler\StreamHandler;


class ConfigureLogging extends BaseConfigureLogging
{
    /**
     * Configure the Monolog handlers for the application.
     *
     * @param  \Illuminate\Contracts\Foundation\Application  $app
     * @param  \Illuminate\Log\Writer  $log
     * @return void
     */
    protected function configureSingleHandler(Application $app, Writer $log)
    {
        // Stream handlers
        $logPath = '/var/log/applications/app_name/app.log';
        $logLevel = Monolog::DEBUG;
        $logStreamHandler = new StreamHandler($logPath, $logLevel);

        // Formatting
        // the default output format is "[%datetime%] %channel%.%level_name%: %message% %context% %extra%\n"
        $logFormat = "%datetime% [%level_name%] (%channel%): %message% %context% %extra%\n";
        $formatter = new LineFormatter($logFormat);
        $logStreamHandler->setFormatter($formatter);

        // push handlers
        $logger = $log->getMonolog();
        $logger->pushHandler($logStreamHandler);
    }
}

{% endhighlight %}

The code above shows how you can override the `'single'` logging for Laravel. If you prefer another logging mode, e.g. `daily`, then you can replace `configureSingleHandler` with `configureDailyHandler`. Check [here](http://laravel.com/docs/5.0/errors#configuration){:target="_blank"} for a full list of available logging modes.

If your app is using a different namespace, be sure to use that one on line `1`, I've just named mine `Bootstrap` for demonstration purposes.

Next, in your `composer.json` file, add namespace details to the `autoload` object under `psr-4`:
{% highlight json %}
"autoload": {
    ...
    "psr-4": {
        "App\\": "app/",
        "Bootstrap\\": "bootstrap/"
    }
    ...
},
{% endhighlight %}


Then, we need to replace the `ConfigureLogging` bootstrappers in `Http/Kernel` and `Console/Kernel` to use our custom one. We'll do this by overriding their respective constructors and doing an `array_walk` to the bootstrappers property.

Add the code below to `/app/Http/Kernel.php`:
{% highlight php %}
<?php
...
use Illuminate\Contracts\Foundation\Application;
use Illuminate\Routing\Router;
...

...
public function __construct(Application $app, Router $router)
{
    parent::__construct($app, $router);

    array_walk($this->bootstrappers, function(&$bootstrapper)
    {
        if($bootstrapper === 'Illuminate\Foundation\Bootstrap\ConfigureLogging')
        {
            $bootstrapper = 'Bootstrap\ConfigureLogging';
        }
    });
}
{% endhighlight %}

Then this code to `/app/Console/Kernel.php`:
{% highlight php %}
<?php
...
use Illuminate\Contracts\Foundation\Application;
use Illuminate\Contracts\Events\Dispatcher;
...

...
public function __construct(Application $app, Dispatcher $events)
{
    parent::__construct($app, $events);

    array_walk($this->bootstrappers, function(&$bootstrapper)
    {
        if($bootstrapper === 'Illuminate\Foundation\Bootstrap\ConfigureLogging')
        {
            $bootstrapper = 'Bootstrap\ConfigureLogging';
        }
    });
}
...
{% endhighlight %}

Note that we're using `Illuminate\Routing\Router` in `/app/Http/Kernel.php` and `Illuminate\Contracts\Events\Dispatcher` in `/app/Console/Kernel.php`

Now you can use the logger in your application, e.g.
{% highlight php %}
<?php
...
use Log;
...

...
public function goLooney()
{
    Log::info('I am Bugs Bunny!');
}
{% endhighlight %}

This should appear as below in your log file:
![Successful logging]({{ site.url }}/images/2015-05-19/logged-success.png)

*PS:* The extra square brackets are serialized representations of empty arrays. The `LineFormatter` class allows you to pass extra arguments to the log, e.g. if you want to pass an array:
{% highlight php %}
<?php
...
public function goLooney()
{
    $looneyFacts = array(
        'favoriteFood' => 'orange carrots',
        'bestFriend' => 'Daffy Duck'
        );
    Log::info('I am Bugs Bunny!', $looneyFacts);
}
{% endhighlight %}

it will be added to the log:
![Successful logging]({{ site.url }}/images/2015-05-19/extra-array.png)

If the extra brackets are really bugging you (:grin:), check out how you can work around it in this StackOverflow question: [http://stackoverflow.com/questions/19935899/laravel-logging-extra-square-brackets-at-end-of-log-lines](http://stackoverflow.com/questions/19935899/laravel-logging-extra-square-brackets-at-end-of-log-lines)

I hope this is helpful to someone! Back to code!

###### I pooled the information above from:
- [Laravel Official Docs](http://laravel.com/docs/5.0/errors#configuration){:target="_blank"}
- [Monolog Official Docs](https://github.com/Seldaek/monolog/blob/master/doc/usage.md){:target="_blank"}
- [Laravel Forums](http://laravel.io/forum/02-06-2015-laravel5-how-to-change-logs-path){:target="_blank"}
- Laracasts Discussions
    - [https://laracasts.com/discuss/channels/general-discussion/advance-logging-with-laravel-and-monolog](https://laracasts.com/discuss/channels/general-discussion/advance-logging-with-laravel-and-monolog){:target="_blank"}
    - [https://laracasts.com/discuss/channels/general-discussion/error-on-overriding-configurelogging-bootstrap-class?page=1#reply-42802](https://laracasts.com/discuss/channels/general-discussion/error-on-overriding-configurelogging-bootstrap-class?page=1#reply-42802){:target="_blank"}
- Witty log contents thanks to [Fortune](http://en.wikipedia.org/wiki/Fortune_(Unix)){:target="_blank"}
