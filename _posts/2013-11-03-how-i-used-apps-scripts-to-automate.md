---
layout: post
title: How I Used Apps Scripts to Automate Daily Report Generation
date: '2013-11-03T20:16:00.002+03:00'
author: Fred Muya
tags:
modified_time: '2013-11-03T20:16:56.831+03:00'
---

Repetitive tasks… I don’t like them. They get monotonous.That’s what was happening a few days ago, when I was required to create some daily reports on our systems’ health, and upload them to a Google Drive Spreadsheet, i.e., run a few queries on the db, copy the data to the Spreadsheet (probably via a text editor, because Drive doesn't know any other delimiters other than the tab :-(

Since the reports will be required daily for the next few weeks, I decided to automate the process. Our InfoSec team said they don’t want us to expose any APIs to any external systems without them being over a private network (P2P, VPN, etc), and APIs would take too much time to build, test, deploy, etc. So here’s what I came up with:

- Write a PHP script that will be running my queries, saving them in JSON files, and sending the generated files to my email (that is allowed :-)
- Setup the PHP script to as a cronjob on our servers
- Setup a Gmail filter to detect and give a label to the aforementioned email (We use Google Apps for business)
- Create a Script (on our Google Apps domain), to search my inbox for the email using the given label, fetch the attachments, parse the data and load it into Spreadsheets, then mark the email as read, then archive it.
- Setup the triggers on the script to run daily (like a cron on the cloud)
- Then finally, no more zombie-mode for 10-15 minutes each morning

I’d like to share the stuff I did, all the way from the PHP script, to the Apps Script app. Any improvements, criticism, etc will be much appreciated. Let’s get to code…

Everything will be available on [GitHub](https://github.com/muya/jsc-mentions)

For the data, I won’t give you anything from our servers, however we’ll be analyzing how frequently the JSC case came up on the news based on platform and location (don’t ask about the data set, someone gave me a challenge to make Law technical, it’s a start :)

So, to get started, check-out/clone the project from [https://github.com/muya/jsc-mentions](https://github.com/muya/jsc-mentions) onto any folder on your system. I’ll assume a Unix-based system because… what else?
{% highlight bash %}
$ git clone https://github.com/muya/jsc-mentions.git
{% endhighlight %}

I’ve provided some generic functions that you may re-use in any of your projects, (thank me later!)

#### DATA SETUP
The following files can be found under the **`db`** folder:

- `JSCDatabase.sql` - Schema file that will create the database, and the tables required.
- `JSCDatabase_with_data.sql` - Schema file that will create the database, the tables and load the data

_**NB:** If you choose to use the schema without data, you can run the load data scripts under the test folder from the terminal._

I’ll take you through some basic configurations in `config/_config.php`:

- Ensure you’ve created the path for `LOG_DIRECTORY` on your system
- `EMAIL_RECIPIENTS` is a pipe-separated list of email addresses that will receive the data in JSON format
- `MAILER_USERNAME` & `MAILER_EMAIL` are the full email addresses of the account that will be used to send the email.

`cd` into the directory where you cloned the project and run the command below (make sure you have a working Internet connection):
{% highlight bash %}
$ php loadJSCData.php
{% endhighlight %}
If all goes well, you should have an email in your `EMAIL_RECIPIENTS` account, having 2 attachments, one for mentions per location and another having mentions per platform.

Next step, is set up on Google Apps Scripts:

- Open [Google Drive](https://drive.google.com/), and create a new Script
- Select `Blank Project`
- Replace the contents of `Code.gs` with the file contents of the file in `appsScripts/LoadJSCMentionsData.gs`
- Click on `Select Function` & Choose `loadJSCMentionsData`
- Run!

_**NB:** The script will ask for permissions to Manage your email & Google Drive. (Don’t worry, encore254 is NOT affiliated with NSA, NIS, or any other Alphabet bodies)_

A Spreadsheet will be created in your root Drive folder with the data loaded from the JSON files sent to your email.

You’re free to dissect, critic & improve the any aspects of the code. Also suggest any improvements, corrections, etc. via comments on this blog or on [GitHub](https://github.com/muya/jsc-mentions).

Special thanks to [GDG Nairobi](http://www.gdgnairobi.info/) for the Hackathon that introduced me to Google Apps Scripts.
