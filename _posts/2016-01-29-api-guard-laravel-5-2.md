---
layout: post
title: Using API Guard with Laravel 5.2
author: Fred Muya
---

If you've been having problems with using [chrisbjr's ApiGuard](https://github.com/chrisbjr/api-guard) library in your project after upgrading to Laravel 5.2, I've outlined the steps to take in order to have it working; particularly if you are getting the error below:

{% highlight bash %}
exception 'BadMethodCallException' with message 'Method [beforeFilter] does not exist.'
{% endhighlight %}

The error above appears because the `beforeFilter` method was deprecated in Laravel 5.1, and removed in 5.2; Laravel now uses Middleware.

This change was [merged](https://github.com/chrisbjr/api-guard/pull/85) into the master branch of the ApiGuard module, but for some reason is not available when you do composer update. The fix involves updating your `composer.json` file to use `dev-master`.

First, update your `composer.json` file as follows:

{% highlight json %}
...
"require": {
    ...
    "laravel/framework": "5.2.*",
    "chrisbjr/api-guard": "dev-master",
    ...
},
...
{% endhighlight %}

Next, run `composer update` for your project.

Finally, add the apiguard middleware to the `$routeMiddleware` array in `app/Http/Kernel.php` as shown below:

{% highlight php %}
<?php
protected $routeMiddleware = [
    ...
    'apiguard' => \Chrisbjr\ApiGuard\Http\Middleware\ApiGuard::class,
    ...
];
?>
{% endhighlight %}
That's it! You should be able to get back to creating amazing apps with Laravel! Till next time, bye!


I wouldn't have been able to figure this out without the resources below; [Asanteni sana!](https://translate.google.com/#sw/en/Asanteni%20sana){:target="_blank"}:

- [Laravel Github Issue](https://github.com/laravel/framework/issues/11640){:target="_blank"}
- [Pull Request on the ApiGuard repo](https://github.com/chrisbjr/api-guard/pull/85){:target="_blank"}
- And of course, [SO](http://stackoverflow.com/questions/19009334/when-dev-master-should-be-used-in-composer-json){:target="_blank"}

