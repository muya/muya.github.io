---
layout: post
title: Using Gii (Part 1)
date: '2012-08-10T08:47:00.000+03:00'
author: Fred Muya
tags:
- Yii Framework
modified_time: '2012-08-23T08:04:21.160+03:00'
---

In this section, we are going to use the tool that comes with Yii to enable us to generate Models, Views and Controllers for our project. You can learn more about [MVC frameworks](http://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller){:target="_blank"}  on Wikipedia.

Gii is a GUI that is used to automatically generate our models, views and controllers without writing a single line of code. This simplifies our work a lot. In order to continue, we need to make a few configuration changes.

Firstly, we need to set up our database connection. Yii allows us to configure this in one place for use throughout the application. Navigate to `protected/config/main.php` under the project folder. This is the main configuration file for the application. It’s simply a PHP file that returns an array of configuration options. The main ones include:

- `basePath`
- `name`
- `preload`
- `import`
- `modules`
- `components`

Some of the configurations have arrays and arrays of arrays within them. The purpose of most of these will become evident as we continue building our project. For this step, we’ll deal with the `modules` and `component` properties.
Gii is one of the modules. To enable it, simply uncomment the code that represents the array defining it, so that it looks like:
{% highlight php startinline=true %}
'gii'=>array(
 'class'=>'system.gii.GiiModule',
 'password'=>'Enter your password here',
 // If removed, Gii defaults to localhost only. Edit carefully to taste.
 'ipFilters'=>array('127.0.0.1','::1'),
),
{% endhighlight %}

`class` defines the class that loads the Gii module. `password` defines the password you will use to access the gii module. `ipFilters’ defines an array of IP addresses that are allowed to access the Gii module. This is important for security reasons: since Gii is an essential system tool that can be used to create data manipulating functions, it is important to limit its access to development servers only.

Change the `password` property to something harder to guess.

The component property holds configurations for application components, especially those that are used frequently in the application. One of these is the db connection, defined by `db` property.

By default, Yii defines an SQLite connection. If you use SQLite, you’re good to go. However, if you use MySQL, you will need to comment out the part that defines the SQLite connection, and uncomment the part that defines the MySQL connection. It looks something like:
{% highlight php startinline=true hl_lines=1 %}
'db'=>array(
  'connectionString' => 'mysql:host=localhost;dbname=studentportal',
  'emulatePrepare' => true,
  'username' => 'root',
  'password' => 'password',
  'charset' => 'utf8',
),
{% endhighlight %}

Simply update the highlighted (above) parts with your own configurations.
Note: It is advisable to create users in MySQL for specific databases to prevent access to all databases by one user. For instructions on how to create users in MySQL and grant them permissions on databases and tables, you can check [this](http://encore254.blogspot.com/2012/08/create-users-and-grants-in-mysql.html)

With Gii enabled and the database set up, we are ready to proceed. Go to [http://{YOUR_HOST_NAME}/StudentPortal/index.php/gii](http://localhost/StudentPortal/index.php/gii)

In the form displayed, enter the password you set up in the config/main.php file. On successful login, the page displayed welcomes you to Yii Code Generator...you are good to go...

In the next section, we will be using Gii to create the models, views and controllers.
