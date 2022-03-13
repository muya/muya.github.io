---
layout: post
title: Muya's DevTools
date: '2022-03-13'
author: Fred Muya
excerpt: Collections of useful tools I use on my computer daily
tags:
- productivity
- devtools
---

**Before You Continue: This is the "lite version" of this blog post, and doesn't have GIFs showing off the different functionality. [Access the "full" version of this page here]({{ site.url }}/productivity-tools-with-animations/)**

To kick off the blog posts for the year, I thought it might be fun to share some tools that I use daily, and that I find invaluable to my workflow.

My day involves doing software engineering at a "Staff" level, which means that in addition to coding, I also do a lot of outward communication to other teams, code reviews, coordinate between team members, among other things.

Therefore, the tools I'll share today *will not* be centered around coding.

The tools are divided into three: 
- IDEs & IDE-related plugins to cover the "coding" side of things
- Useful apps I use daily
- CLI - for all terminal-related stuff (I promise I'm not that kind of hacker 😅)


Disclaimer: I work primarily on macOS, so most of the apps will lean towards that, but I'll highlight similar `*nix` tools that I've used as well, where applicable.

Enjoy!

## IDEs
A coder's best friend, and in many situations, a daily driver.

### Sublime Text
Yes, yes - I know VSCode is a thing, but for some reason I've never been able to stop using Sublime, at least for scenarios where I don't need a fully featured IDE.

Through its plugins, it offers a wide array of tooling that is also found on VSCode, but with one major advantage - it's really fast, and doesn't hog any resources on my computer (last I checked, VSCode was built using the Electron framework, and that was / is known to be quite a resource hog).

Sublime Text is built with Python, which provides a "native feel" in terms of how it works.

As always, to each their own - maybe one day I'll give VSCode another chance.

One of my favourite things about Sublime Text is the "unified command" access via `Cmd+Shift+P` - this command gives you access to all the commands that can be run on the editor, including those made available by plugins.
This, combined with the fuzzy search function means that you can get access to your most commands with only a few taps on your keyboard.

The most recently used commands are also kept highlighted, making access even faster.

Plugins that I use commonly:
- [Pretty JSON](https://packagecontrol.io/packages/Pretty%20JSON) - Offers a lot of helpers for minifying, prettifying, validating and linting JSON
- [Markdown Editing](https://github.com/SublimeText-Markdown/MarkdownEditing) - provides utilities to help you work with Markdown (I'm utilizing it to write this post!), including writer friendly themes
- [SideBarEnhancements](https://github.com/titoBouzout/SideBarEnhancements) - enhances the right-click menu on the navigation pane to offer more file access options
- [GitGutter](https://packagecontrol.io/packages/GitGutter) - Provides additional context to the editor to show changes against the most recent `git HEAD`


Check out [Package Control](https://packagecontrol.io/) for a full list of available plugins to help you write your code.

### JetBrains Family of IDEs
For when I need to do "heavy" development work (read: when I need a fully featured IDE), the JetBrains IDEs are my tools of choice.

I primarily write code in GoLang and Scala, and GoLand & IntelliJ Ultimate are my preferred tools of choice.

What makes these IDEs amazing, despite them being system resource hogs at times, are the plugins that allow you to enhance your experience while using them.

My favourites:
- [IdeaVim](https://plugins.jetbrains.com/plugin/164-ideavim) - makes Vi(m)'s  (yes, the CLI text editor) commands available for use in the IDE - this way you get the best of both worlds: speedy access to functionality if you're comfortable with Vim, and access to the IDEs built-in tools!
- [GitHub Copilot](https://plugins.jetbrains.com/plugin/17718-github-copilot) - (cue all the "AI is taking all our jobs" narratives). In case you haven't come across it, GitHub Copilot is a helper that uses AI to automatically generate code for you - think of it as an automatic pair programming companion.

    When GitHub Copilot was [initially announced](https://github.blog/2021-06-29-introducing-github-copilot-ai-pair-programmer/), I was a little skeptical about how good it would be, but I signed up for the beta anyway - and holy @#@! was I amazed by how good it's been. Not only does it accurately deduce the code it thinks you're about to write, but it's also contextual. I've been so impressed by this that I'm constantly tempted to start a #github-copilot channel at work just to share the cool stuff. Until then, please enjoy my geeking out on Twitter instead :-D

    <div class="jekyll-twitter-plugin" align="center">
        <div class="jekyll-twitter-plugin"><blockquote class="twitter-tweet" data-theme="dark"><p lang="en" dir="ltr">GitHub Co-Pilot keeps blowing my mind at every turn! Amazing!! 🤯 <br /><br />In this screengrab, we&#39;re extracting images from the Blogger API <a href="https://t.co/tPTZsMQJmt">pic.twitter.com/tPTZsMQJmt</a></p>&mdash; Muya (@fred_muya) <a href="https://twitter.com/fred_muya/status/1469653717653860358?ref_src=twsrc%5Etfw">December 11, 2021</a></blockquote>
    <script async="" src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
        </div>
    </div>

- [GitLink](https://plugins.jetbrains.com/plugin/8183-gitlink) - this plugin allows you to right-click on any file, and open it in your preferred code hosting platform (GitHub, GitLab, BitBucket, etc); it's so handy especially in situations where you want to share the location of an exact line with someone else on the team.
- [Material Theme UI](https://material-theme.com/docs/introduction/) [Free with Paid Options] - what's an IDE without some personalization; I've always been a fan of Google's Material UI, and much as it didn't take off on the web, I really enjoy it within IntelliJ. This plugin provides an extreme level of customization to make your IDE feel at home; after all, this is where you probably spend most of your time during the day, and you might as well make it personal.
    + [Atom Material Icons](https://plugins.jetbrains.com/plugin/10044-atom-material-icons) - somewhat related, this plugin changes all the IDE icons into Material Icons - this completes the look!

## Apps
With computers being our main tools of trade, there are some apps that make it a joy to work on my machine - check them out!

### Brave Browser
Have you ever wanted "Chrome, but not resource-hog Chrome"? I've found Brave Browser to be a viable alternative, and I've been using it for the past few years since its launch.

[Brave](https://brave.com/) is a privacy focussed browser that blocks trackers, ads and a bunch of other malicious JavaScript on your browser by default. By just doing this, your web pages load much much faster (up to 3-6x faster than Chrome / Firefox, according to their website).

My second favourite feature is "Brave Rewards", which is an opt-in feature that allows you to earn frequent flyer-like tokens ([BAT](https://basicattentiontoken.org/)) for viewing privacy-respecting ads, should you wish to.

The BAT you earn can be used to pay creators whose content you enjoy (like this website, eh? 😉).

Check out the full host of features here: [https://brave.com/features/](https://brave.com/features/).

The coolest part of this whole experience is that all the functionality that's available on Google Chrome, such as extensions & [browser profiles](https://www.chromium.org/user-experience/multi-profiles/), is also available on Brave, since they're both based on the Chromium Open Source project!

Brave is also available on mobile!

#### Valuable Extensions
While we're on the topic of browsers, there are some extensions that I really enjoy using:
- [ProductHunt](https://chrome.google.com/webstore/detail/product-hunt/likjafohlgffamccflcidmedfongmkee) - the only extension that I've ever had as a "New Tab" page - it showcases the latest products listed on [Product Hunt](https://www.producthunt.com/), usually a nice way to skim through new tech & services
- [Vimium](https://chrome.google.com/webstore/detail/vimium/dbepggeogbaibhgnhhndojpepiihcmeb) - Again with the Vim-in-everything. This one provides additional tools to navigate your browser using Vim shortcuts!
- [Save to Pocket](https://chrome.google.com/webstore/detail/save-to-pocket/niloccemoadcdkdjlinkgdfekeahmflj) - this is the browser companion to [Pocket](https://getpocket.com/), which is an amazing "read it later" service, allowing you to save articles for later.
- [Pushbullet](https://chrome.google.com/webstore/detail/pushbullet/chlffgpmiacpedhhbkiomidkjlcfhogd) - [Pushbullet](https://www.pushbullet.com/) is a service that allows you to send media between devices. This is extremely useful, especially if your workflow is spread across multiple platforms (Android / iOS / macOS / Windows / etc).
- [Zoom Closer](https://chrome.google.com/webstore/detail/zoom-closer/appjbedfhcmpknanmbndpojcllfaemal) - If you use Zoom frequently, this is a useful extension that automatically closes those Zoom windows that open when you're joining a call; no more lingering Zoom call tabs!

### Cleanshot X (Paid)
Cleanshot X was my find of the year for 2021! This is an amazing screen capturing tool that gets stuff done! (It's what I used for all the screen-grabs in this post!)

A non-exhaustive list of features (i.e. my favourites):
- Screenshotting (is that a word?)
- Annotation on just captured screenshots (including cropping, combining images, drawing, blurring, adding text, adding sequential numbers, highlighting)
- Scrolling Capture - allows you to take a screenshot beyond your screen real estate. If the app you're using scrolls, then you can take a screenshot of everything!
- Capture Text (OCR) - do you want to extract text from an image? Take a screenshot using the OCR functionality, and it automatically copies the identified text to the clipboard!
- Screen recording (both GIFs & videos) - with the ability to pause your recordings (e.g. to allow you to enter sensitive information). The screen recording also comes with an option to show any mouse clicks and / or keyboard inputs.

My few words here do not do this tool enough justice, so I'll leave this walk-through here: [CleanShot X for Mac](https://www.youtube.com/watch?v=FZbICrBKWIU)

This tool has an annual subscription, but I think it's worth it if you can get it! (Maybe you can even expense it for work?)

Skitch, described below, is a close, second, free option - it was my tool of choice before I discovered CleanShot X!


#### Skitch
[Skitch](https://evernote.com/products/skitch) is another screen grabbing tool that I can recommend as an alternative to CleanShot X. It has most of the same great features, and it works great!

### Numi
[Numi](https://numi.app/) is an enhanced calculator app for macOS. It provides a slew of features, with my favourites being:
- Unit Conversion for almost any unit: `cm` → `inch`, `Celsius` → `Fareinheit`, etc
- Currency Conversion, including support for cryptocurrencies
- Variable assignment, which allows you to do "light coding" within the calculator

Check out the full feature set [here](https://github.com/nikolaeu/numi/wiki/Documentation)!

For my set-up, I've enabled a global shortcut that allows me to bring up Numi from within any app using `Alt + N`, making it available at my fingertips!

### Dash
[Dash](https://kapeli.com/dash) is an API Documentation Browser, and a Code Snippet Manager. 

I use it primarily as the latter, and particularly taking advantage of the "Expand Snippets" feature, which allows you to automatically place a snippet anywhere across your system.

For example, if you have a database command that you frequently use, e.g.

{% highlight plaintext %}

mysql -h localhost -p -u user

{% endhighlight %}

You can set up an abbreviation (or a trigger) to automatically fill this snippet. E.g. you can configure:

{% highlight plaintext %}

`db

{% endhighlight %} 


The "Expand Snippets" feature becomes even more powerful with the "placeholders", which autofill useful content:
- `@clipboard` automatically pastes the contents of your keyboard within the automatically expanded snippet
- `@cursor` moves the cursor to the set position after the snippet is expanded
- `@date` and `@time` add a timestamp to your snippet

### Slash

**Pricing**: [Paid, with trial period in the form of "tasks"](https://getslash.co/pricing).

[Slash](https://getslash.co/) is a To-Do app, and one of the best I've come across. It's marketed as a helper that helps you stay organized and focussed on your work, and I can attest to this.

It works by allowing you to choose your tasks for the day, allowing you to prioritize them, then you can "get slashing!".

Once you start, a small bar appears at the bottom of your screen, showing you your current task, and how long you've been at it. One of my favourite things about this is that you can set it to bounce every 10 minutes, which is a great way to ensure your attention is still on the task!

The most satisfying aspect of it though, has to be the animation + random GIF that shows when you mark a task as complete! 


### ItsyCal
[ItsyCal](https://www.mowglii.com/itsycal/) is a menu bar app that provides quick access to a calendar. It's simple and non-intrusive, which is why I like it.

It has options to display the current week, and provides integration to your Calendar, allowing you to view events that are coming up.

### Clocker
[Clocker]() is another menu bar app that provides a quick way to check the time in different cities across the world. This is especially useful if you're working across multiple timezones.

Again, the simplicity of this app is a reason why I enjoy it so much.

My favourite feature is the "Time Scroller" that allows you to move time forward / back to see the corresponding time in another city!

### Tower (Paid, with 30-day trial)
Before telling you about this one, let me make it clear: knowledge and proficiency in the Git CLI tool is an invaluable skill to have, and I highly recommend learning the basics at the very least!

That said, sometimes a graphical user interface is useful to help you perform some tasks quickly.

This is where [Tower](https://www.git-tower.com/) comes in!

It's a Git app that organises all your projects, and allows you to perform most of the actions provided by Git.

Some of my favourite things about the app:

- Provides a visual way to resolve merge conflicts (and integrates to your diff tool of choice)
- Multiple windows for multiple projects
- Solid, well-crafted UI that's a joy to use
- An amazing staging area, that allows you to craft your commits meticulously - it even allows staging of specific lines only.

### f.lux
Have you ever been using your computer in the evening / late at night, and the screen brightness becomes a recipe for your lack of sleep that night?

This is because most screens are designed to be as bright as possible, which is okay for using them during the day, but wreak havoc if you're using your devices later at night.

[f.lux](https://justgetflux.com/) is built to help solve this by automatically adjusting the colour of your computer's display to match the time of day based on your location: warm at night, and like sunlight during the day.

The app lives on your menu bar, and is non-intrusive (noticing a pattern here 😅).

Additionally, it has a "Movie Mode" setting, which temporarily disables the changes to allow you to enjoy your movies with the correct lighting.

### CloudMounter
[CloudMounter](https://cloudmounter.net/) is an application that allows you to mount your cloud storage (Google Drive, Dropbox, FTP, Amazon S3, whatever) on your computer, making the cloud storage accessible directly from your computer. This effectively makes your cloud storage accessible directly on your desktop.

It's simple, and it works great!

The service is free for one integration, and requires a paid upgrade to connect more than one source.


## CLI Tools
For those spending a lot of time using the command line, having CLI tools that make your life easier is a boon!

### iTerm (macOS) / Terminator (*nix)
The default terminal applications that come with operating systems are useful, but these two up the game a lot!

- [iTerm](https://iterm2.com/) - for macOS
- [Terminator](https://gnome-terminator.org/about/) - for other *nix Operating Systems


These applications provide a significant level of customization options that go beyond the cosmetic, and make you a much more productive user.

The customization options available are just too numerous to name in a post like this, so I highly recommend that if you're not using these, stop what you're doing, download them, and make them the default terminal.

Then depending on your style, explore the docs ([iTerm](https://iterm2.com/features.html) / [Terminator](https://gnome-terminator.readthedocs.io/en/latest/)) OR go to Settings, and check out each option available - you'll have fun, I promise!

The feature that I consider to be the highest multiplying factor in your productivity, however, is the Split Pane feature! Get the apps for this, if nothing else!

For my setup, I've hooked up shortcuts (`Alt+h` / `Alt+v`) to split my window panes horizontally and vertically, respectively. Check it out in action (enjoy the fortune jokes in my terminal):

### fortune
[Fortune](https://en.wikipedia.org/wiki/Fortune_%28Unix%29) is a Unix program that displays a random (sometimes funny) message - also called fortune cookies.

It's been traditionally used to set a "Message of the Day (MOTD)" when logging in to servers, but there's no reason you shouldn't use it on your terminal.

For my setup, I've linked added it to my environment configuration file (`.zsrc`) to have it run every time I open a new terminal window.

_Note: some messages might not be safe for work - so if you regularly use your terminal at work, it's best not to enable offensive fortunes (offensive messages are off by default)_

Check out the screen-grab above for some messages.

### tabset
What did [Biggie Say? Mo' Terminals, Mo' Problems](https://www.youtube.com/watch?v=ss142Aix2Bo)? No? Okay!

Well, with more terminals running, it's always nice to have a visual identifier of what's running within them.

[tabset](https://www.npmjs.com/package/iterm2-tab-set) is an iTerm add-on that enables naming and colouring of tabs and windows - which is especially useful if you have many of them open.

### OhMyZSH
With Apple moving away from `bash` to `zsh` as the default shell, I was finally forced to try it out, and what's using `zsh` without making use of the [OhMyZSH](https://ohmyz.sh/) framework?

You can think of OhMyZSH as a framework built on top of the `zsh` shell (equivalent of Play! to Scala, or Laravel to PHP).

I like this quote on their home-page:

> Oh My Zsh will not make you a 10x developer...but you may feel like one!

The main provisions of OhMyZSH are the plugins, and the themes!

My favourites:

- [`git`](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/git) - provides additional context to your CLI when working with Git-hosted projects, including aliases and functions. The functions can then be used to further customize your CLI. For example, the info it provides can be used to customize your theme, and add some context about your current branch. In my case, I use it to display the current branch, and whether or not the current project is "clean" (i.e. doesn't have modified files)
  ![OhMyZSH git plugin]({{ site.url }}/images/2022-03-13/zsh-git-plugin.png)
- [`zsh-z`](https://github.com/agkozak/zsh-z) - allows you to quickly jump to directories you've used recently, or that you use frequently (a concept called, believe it or not, [Frecency](https://en.wikipedia.org/wiki/Frecency)). Instant productivity boost!
- [`zsh-vi-mode`](https://github.com/jeffreytse/zsh-vi-mode) - The default OhMyZSH comes with a [default vi-mode](https://github.com/ohmyzsh/ohmyzsh/tree/master/plugins/vi-mode), but this plugin makes the vi(m) mode more user-friendly.

### fzf
[fzf](https://github.com/junegunn/fzf) is a general purpose command line fuzzy finder - you feed it any input, you get instant fuzzy search functionality. 

Have a list of files in a directory? Feed it into `fzf`, and you can search faster

Want to quickly search through the results of a `locate` call? `fzf` to the rescue!

It even integrates to your "Search Command History (usually bound to `Ctrl+R`)" to enable fuzzy search:


### More CLI Tools
[@amilajack](https://twitter.com/amilajack/status/1479328649820000256?s=20&t=lz6cZG12bTXJKu4EFgiFvQ) wrote an excellent thread on some amazing CLI tools (it's where I found out about `fzf`) - check them all out!

<div class="jekyll-twitter-plugin" align="center">
    <div class="jekyll-twitter-plugin"><blockquote class="twitter-tweet" data-theme="dark"><p lang="en" dir="ltr">Life is to short to use dated cli tools that suck<br /><br />Try these instead 🧵</p>&mdash; Amila Welihinda (@amilajack) <a href="https://twitter.com/amilajack/status/1479328649820000256?ref_src=twsrc%5Etfw">January 7, 2022</a></blockquote>
    <script async="" src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
    </div>
</div>


## Finally...
You made it to the end! Congratulations!

I really hope you enjoyed this post, and found a tool or two that enhance how you work! 

After all, you're only as good as your tools, so it makes sense that you choose the right ones!

Do you have additional suggestions that you can't live without? Let me know!

Happy Coding, and stay safe!
