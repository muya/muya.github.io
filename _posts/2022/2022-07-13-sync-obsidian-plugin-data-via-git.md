---
layout: post
title: 'Obsidian Configuration: Sync Plugin Data'
date: '2022-07-13'
author: Fred Muya
excerpt: 'How To Sync Plug-In Data in Obsidian via Git'
tags:
- obsidian
- obsidian-plugins
- obsidian-git-sync
---

Welcome to another post! 

This is another Obsidian-themed one, and we'll explore the approach I'm using to sync Obsidian Plugin Data.

We'll go over the `.gitignore` configuration I've been using to sync only relevant data within my `.obsidian` folder.

In case you haven't heard of it before, Obsidian is a text-based personal knowledge management tool, which leverages Markdown files to organise all the notes you take.

I gave it a more thorough introduction in [this other blog post on on creating a custom daily notes file name]({{site.url}}/obsidian-daily-notes-custom-file-name/); check it out!

Now, back to syncing!

The most commonly used approaches are:

- [Obsidian Sync](https://obsidian.md/sync) - this is a premium feature offered by the creators of Obsidian, and it provides a cloud-based sync functionality.
- Git - the popular version control system comes in handy when managing files, and it works great for Obsidian (you can go with an automated approach, like the [Git Sync plugin](https://github.com/denolehov/obsidian-git), or with a more manual approach - i.e. treat your vault like a code repository.
- Sync via a cloud-based file system like Dropbox / Box / Google Drive / etc.

Read about all the supported sync options on the Obsidian Help Page: [Sync your notes across devices](https://help.obsidian.md/Getting+started/Sync+your+notes+across+devices).

Personally, I use Git via the manual approach, since it provides the most flexibility.

When syncing your Obsidian data, it's important to decide whether or not you'd like to sync your Obsidian config folder (usually named `.obsidian`). 

This file is usually found at the root of your vault, and contains configurations for both Obsidian, and all the plugins in use.

Some common files & folders found in this directory:
- `workspace` - contains device-specific configurations for how the Vault is currently operating.
- Various JSON files that store configurations for Obsidian (e.g. hotkeys, appearance, enabled core & community plugins), and its core plugins, such as Daily Notes.
- `themes` - contains files for downloaded themes.
- `plugins` - contains directories with the code, styles, manifest & data for installed plugins.
- `snippets` - contains CSS files that can be used to further personalise Obsidian.
- `cache` - used for any file-based caching within Obsidian.

There may be more or fewer files, depending on what you've enabled in Obsidian.

When using Git to sync your Obsidian folders, it's important to know what you'd like to sync, and what you want to remain on the current device.

In my case, I wanted to be able to open my vault on any device and have a similar experience across them in terms of:
- Theme
- Font
- Hotkeys
- Plugin Settings (Note: I only want the plugin settings, but not the plugin source code; i.e. `data.json` within each of the plugin folders)

With the requirements above, this is my current `.gitignore` set up for Obsidian vaults:

{% highlight markdown %}

# Obsidian
**/.obsidian/cache*
**/.obsidian/workspace*
**/.obsidian/plugins/*/*

# Allow data.json files under plugins - we want to sync plugin settings, if any
!**/.obsidian/plugins/*/data.json

{% endhighlight %}


This config automatically excludes the `cache` & `workspace` folders from being committed.

It also excludes everything under the `plugins` folder, and then explicitly allows the `data.json` files within the specific plugin's folder.

Using the `obsidian-linter` plugin folder as an example: it has 4 folders within it:

![Obsidian Linter Plugin Folder]({{ site.url }}/images/2022-07-13/Obsidian-linter-plugin-folder.jpeg)

With the `.gitignore` file above, the `main.js`, `manifest.json` & `styles.css` files will be excluded from Git, but `data.json` will be included, allowing sharing of configs across devices.

The `.gitignore` file also has some other exclusions, such as common ones for macOS & SublimeText (I used Sublime Text heavily to manage my notes before discovering Obsidian).

And that's it! With this setup, I get the same layout & functionality no matter which device I load my vault from!

This approach has been particularly useful for syncing themes, styling & configs for my most commonly used plugins.

I hope you learnt a thing or two.

Until next time, stay safe, and happy coding!
