# Go 1 Hour Boot Camp

Go is a language initially developed at Google and  rapidly becoming the native language for the cloud and distributed computing. This will be a 1 hour boot camp (or longer) where we go through a Go install, followed by 3 short programs to write in go that would include:


The major components involved are:

## Why Go?

Give a brief overiew of why Go was created, who is using it, and some language overview.

## Installing Go

This is a quick start guide to show you how to install Go, and write your first Hello World program.

## Your first Command Line Program

This tutorial will show you how to create a basic command line program that accepts arguments.

### Compiling with Go (And Cross Compiling)

Quickly show you how easy it is to compile a single static binary for distribution, as well as cross compiling.

## Your First Webserver

Learn how to spin up a web server and how to route incoming requests.

## Basic Concurrency

See a couple of basic concurrency patterns.


## Before You Come To Class

The following is a set of tasks that can be done prior to showing up for the workshop.  We will also do this in class if anyone has not completed it.  However, the more attendees that complete this ahead of time the more time we have to cover additional training material.

## Installing Go

Please refer to the [installing go](https://github.com/corylanou/go-1-hour-boot-camp/blob/master/installing-go.md) document, or feel free to do this step during the bootcamp.

### Editors

**Visual Studio Code**
https://code.visualstudio.com/Updates
https://github.com/microsoft/vscode-go

**Sublime**
http://www.sublimetext.com/
https://github.com/DisposaBoy/GoSublime
http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/

**VIM**
http://www.vim.org/download.php
http://farazdagi.com/blog/2015/vim-as-golang-ide/

**Atom**
https://atom.io/
https://github.com/joefitzgerald/go-plus

**LiteIDE**
http://sourceforge.net/projects/liteide/files/

**Emacs**
https://github.com/creack/dotfiles

For a full list of editors, see the wiki: https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

### Installing the Training Material

While many of the examples can be done using the online playground (http://play.golang.org), some may find it easier to complete them with their local editor.  To do so, you will want to load the training material locally to your machine.  From a command prompt, issue the following commands:

```sh
mkdir -p $GOPATH/src/github.com/corylanou
cd $GOPATH/src/github.com/corylanou
git clone https://github.com/corylanou/go-1-hour-boot-camp.git
```

*NOTE:* This assumes you have Git installed.  If you don’t, you can find the installation instructions here: https://git-scm.com/

## NOTE

This is a repo for a one hour boot camp.  I give this talk rather frequently, so I update the date and location I deliver it at, and bug fixes.

You can view the presentation here: [http://go-talks.appspot.com/github.com/corylanou/go-1-hour-boot-camp/presentation.slide#1](http://go-talks.appspot.com/github.com/corylanou/go-1-hour-boot-camp/presentation.slide#1)

## About the Presenter

[Cory LaNou](https://twitter.com/corylanou) a full stack technologist who has specialized in start-ups for the last 17 years. Cory is part of the Core engineering team that contributes to InfluxDB, a highly scalable, distributed time series database written in Go. Cory has created and led numerous Go workshops and training courses. He has several published online articles related to Go. Cory also helps lead and organizer several community technology meetups and mentors new developers.
