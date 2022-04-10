---
layout: post
title: 'Obsidian Daily Notes: Create Custom Filename'
date: '2022-04-10'
author: Fred Muya
excerpt: 'How To Use a Custom Filename for Obsidian Daily Notes'
tags:
- productivity
- obsidian
- obsidian-daily-notes
- daily-notes
- markdown
---

Hello! Welcome to another post!

In this one, we'll be learning how to set a custom file name for your Obsidian Daily Notes.

## What is Obsidian?

[Obsidian](https://obsidian.md/) is officially billed as: "A second brain, for you, forever". 

At its core, it's a knowledge base based on local folder(s) of Markdown files.

Think about it as a management system for your notes, which can be written in plaintext, leveraging Markdown to provide some structure.

This means that notes can be stored directly on your file system, readable in Markdown, and accessible wherever (e.g using Git to sync the files across platforms).

Obsidian builds upon the files by allowing addition of connections between related files, tagging, conversion to presentations, and a whole host of possibilities through the core & community plugins - while keeping it all in Markdown!

While outlining Obsidian's features (the ones I've come across so far) would require several blog posts, I wanted to introduce the "Daily Notes" feature, and show how you can set a custom name for your files.

To provide some context: a few years back when I started doing a lot of (Dev)Ops work, I started maintaining an Operations Log for all the stuff that I did daily, including SQL queries ran, drafts for notes, small code snippets, etc.

I would organize these into daily "changelog" files, and periodically commit them to a private Git repository; this is a practice I maintain to this day! (This blog post is being drafted in one of these changelog files).

The overall structure of the directory where the change log files are stored is:

{% highlight plaintext %}

├── dev
│   ├── 2021
│   │   └── Apr-2021
│   └── 2022
│       ├── Apr-2022
│       └── Mar-2022


{% endhighlight %}

- A top level directory to classify the scope of the files - usually named after the organization I'm working with (or `dev` for the personal workspace)
- 2nd level directory for the current year
- 3rd level directory for the current month
- Files live in the the "month-level" directories

{% highlight plaintext %}

./
├── 2021
│   └── Apr-2021
│       └── CL_05042021
└── 2022
    ├── Apr-2022
    │   └── CL_10042022.md
    └── Mar-2022
        ├── CL_20032022.md
        └── CL_27032022.md

{% endhighlight %}

The filenames are formatted as: 

{% highlight plaintext %}

# {CL}_{today's date}
# For example, for 10th Apr 2022
CL_04102022

# After Obsidian
CL_04102022.md

{% endhighlight %}

Before Obsidian, these were simply text files without an extension. After I started using Obsidian, I started setting the `.md` extension on the files.

## Daily Notes

The "Operation Changelog" is my version of a "Daily Note", in that I'll usually create one every day depending on what I'm working on.

There's even a custom alias configured on the system to quickly create one:

{% highlight plaintext %}

cl='touch CL_`date +%d%m%Y`'

{% endhighlight %}

This creates an empty file with the filename `CL_{today's date}` in the current directory.

As I started to explore Obsidian more, I decided to explore their ["Daily Notes"](https://help.obsidian.md/Plugins/Daily+notes) feature.

It's a "core plugin", which means it ships with Obsidian by default.

![Obsidian Daily Notes Core Plugin]({{ site.url }}/images/2022-04-10/obsidian-daily-notes-core-plugin.png)


Once you enable it, it allows you to quickly create or access your daily note in clicking on the "Daily Note" button in Obsidian's side bar:

![Obsidian Daily Notes Sidebar Button]({{ site.url }}/images/2022-04-10/obsidian-daily-notes-access-daily-note.png)


You also get access to the plugin's settings:

![Obsidian Daily Notes Sidebar Button]({{ site.url }}/images/2022-04-10/obsidian-daily-notes-configurations.png)


- Date format - How you want the date to appear in your daily note filename
- New File Location - Where you want your daily note to be created
- Template File Location - If you'd like your note to be created with certain content already in place, you can choose a template to start from (this is extremely useful for files with similar formats)
- Open on Startup - An option to automatically create and open your daily note when you access your Obsidian vault

The part we're most interested in is the "Date format", which determines how your file will be named.

In this case, we needed to preserve the naming format:
- `CL` - prefix to denote that this is a "ChangeLog file"
- `_` - an underscore to separate the name sections
- `DDMMYYYY` - the current date, in the format: `DDMMYYY`; e.g. for 10th April 2022, we'd want: `10042022`

Obsidian's date formatting is based off [Moment.js](https://momentjs.com/), the legendary JavaScript Date library.

Providing the date format was quite straightforward, however I was originally stumped about how to add the `CL_` part.

By referencing the [Moment.js documentation on formatting dates](https://momentjs.com/docs/#/displaying/format/), I was able to figure out that we could wrap the custom part of the file name in `[]`, and it would be output as-is.

Therefore, the required date format ends up being:

{% highlight plaintext %}

[CL_]DDMMYYYY

{% endhighlight %}

![Obsidian Daily Notes Date Format Config]({{ site.url }}/images/2022-04-10/obsidian-daily-notes-filename-format.png)


And now, when one clicks on "New Daily Note", the file gets created in the format required!


![Obsidian Daily Notes Create New]({{ site.url }}/images/2022-04-10/obsidian-create-new-daily-note.gif)


And that's it! That's how you create a custom filename format for your daily notes!


I hope this was insightful, and you learnt something new today! I'll be sharing more Obsidian features, tips & tricks as I come across them.


Until next time, happy coding, and stay safe!
