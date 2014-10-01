---
layout: post
title: Change Ubuntu 12.04 SlideShow Desktop Background
date: '2012-08-26T17:17:00.000+03:00'
author: Fred Muya
tags:
- Misc
modified_time: '2013-05-22T13:54:32.111+03:00'
---

This tutorial will explain how to modify the slideshow background in Ubuntu 12.04 (Precise Pangolin) to have your own pictures.

The pictures that are shown are loaded from an XML file that is located under `/usr/share/backgrounds/contest/` called `precise.xml`

The basic outline of the XML file is as follows:

{% highlight xml %}
<?xml version="1.0" encoding="utf-8"?>
<background>
     <starttime>
       <year>2008</year>
       <month>04</month>
       <day>08</day>
       <hour>0</hour>
       <minute>0</minute>
       <second>0</second>
    </starttime>
    <static>
       <duration>{SOME-DURATION-IN-SECONDS}</duration>
       <file>{SOME-FILENAME}</file>
    </static>
    <transition>
       <duration>{SOME-DURATION-IN-SECONDS}</duration>
       <from>{SOME-FILENAME}</from>
       <to>{SOME-FILENAME}</to>
    </transition>
    .
    .
    .
</background>
{% endhighlight %}

The static node defines the picture that is showing. It has the following:

- `duration` - how long the picture will stay before changing
- `file` - the full-path filename of the picture

The transition node defines the change from one picture to another. It has the following:
- `duration` - how long it will take to change from one picture to another
- `from` - the full-path filename of the picture it’s changing from
- `to` - the full-path filename of the picture it’s changing to

The PHP program I've written will allow you to create an XML file that has file-names of JPEG (.jpg) files you want on your desktop background slideshow.

Follow the steps below:

1. Download the zip file [here](http://improved-write-xml.googlecode.com/files/WriteXML.tar.gz) and unpack it in your web directory.
2. Open [http://{YOUR-SERVER}/WriteXML/](http://{YOUR-SERVER}/WriteXML)
3. In the input fields:
  - Enter the full path of the folder where the pictures are. The program only uses the .jpg files in that folder
  - Enter the preferred filename (without the '.xml'). E.g. precise
  - Enter the preferred picture transition duration in seconds (Not too many, maybe 2 or 3 seconds)
  - Enter the preferred picture display duration in seconds (Not more than the transition duration
  - Click 'Create XML'
4. If there were no errors, a page will appear telling you to download the xml file
5. Click on 'here'. Right-click the page and:
  - On Google Chrome, click ‘Save As...’
  - On Firefox, click ‘Save Page As...’
  - On Opera, click ‘Source’. Click on the ‘Save’ button at the top of the page.
6. Open where you downloaded the file and rename it to `precise.xml`.
7. Make a backup of the existing `precise.xml` file. For example, if you're going to put the backup on your desktop, open a terminal window and do:
{% highlight bash %}
$ sudo cp /usr/share/backgrounds/contest/precise.xml
{% endhighlight %}
8. `cd` to where you downloaded the file and do:
{% highlight bash %}
$ sudo mv precise.xml /usr/share/backgrounds/contest/precise.xml
{% endhighlight %}
9. Go to System Settings, Click on Appearance, under the Look tab, select the picture with a clock on the bottom right corner. Under the monitor on the left, it should say: _Ubuntu 12.04 Community Wallpapers (multiple sizes)_

And that’s it! The pictures in the directory you specified should be showing on your desktop background

_**NB:** If your screen is blank, it means that the picture files cannot be read from their location. Give read permissions to the folder where the files are located._

All the best!!!
