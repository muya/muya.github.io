---
layout: post
title: Create Users and Grants in MySQL
date: '2012-08-10T08:37:00.001+03:00'
author: Fred Muya
tags:
- MySQL
modified_time: '2012-08-26T17:25:18.762+03:00'
---

This blog post provides instructions on how to set up multiple users on MySQL. This is an advisable practice since it allows you to restrict access to certain databases to certain users only.

It is also advisable to create another super user (same as root) as a backup for super administrative rights to MySQL.

To create users, the current MySQL user has to have permissions to create on the mysql database. The root user is usually sufficient.

The basic command for creating users in MySQL is:
{% highlight mysql %}
CREATE USER 'username'@'host' IDENTIFIED BY 'password';
{% endhighlight %}

For example, to create a user `john` on `localhost` with the password as `pass`, run:
{% highlight mysql %}
CREATE USER 'john'@'localhost' IDENTIFIED BY 'pass';
{% endhighlight %}

There is also another way to create users on MySQL. This method involves creating users while at the same time giving grants to tables. Grants in MySQL are similar to permissions.

The command is:
{% highlight mysql %}
GRANT ALL ON db_name.* TO 'username'@'host' IDENTIFIED BY 'password';
{% endhighlight %}

This grants all permissions (`INSERT`, `UPDATE`, `SELECT`, `DELETE`, `DROP`, etc) to the user on all tables in the `db_name` database.

You can grant permissions to specific tables only, e.g.
{% highlight mysql %}
GRANT ALL ON db_name.tbl_1 TO 'username'@'host' IDENTIFIED BY 'password';
{% endhighlight %}

This will give permissions to the user on `tbl_1` table only.

If the user already exists, you don’t have to specify the IDENTIFIED BY part. However, you may use it if you want to change the user’s password.

Alternatively, you can give grants on certain operations to a user, e.g:
{% highlight mysql %}
GRANT SELECT, UPDATE ON db_name.* TO 'username'@'host' IDENTIFIED BY 'password';
{% endhighlight %}

To replicate the root user, you can run:
{% highlight mysql %}
GRANT ALL ON *.* TO 'new_root_user'@'host' IDENTIFIED BY 'password';
{% endhighlight %}

This gives all permissions on all tables of all databases to `new_root_user`.

So, to give permissions on a certain table to a user, simply run:
{% highlight mysql %}
GRANT ALL ON db_name.* to 'new_user'@'host' IDENTIFIED BY 'password';
{% endhighlight %}
