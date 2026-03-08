# Boost Your Beginner Tech Career

The *Beginner Boost* is a quick start into the tech world for absolute beginners with little or no experience in tech. All you need is a computer, a terminal emulator, and a container engine. Getting these is described in sections below and in [the videos](https://youtube.com/playlist?list=PLrK9UeDMcQLrf-bVcwlx8lbOtXtgk4Nsr&si=Aeshy2nqmpzhfbb7) (which are constantly updated).

Check in with [Rob's socials](https://linktr.ee/rwxrob) to see when and where he's live and come say hello or stop by the [forum here on GitHub](https://github.com/rwxrob/boost/discussions).

>[!IMPORTANT]
>Please help and contribute any way you can. Financially [GitHub sponsorship](https://github.com/sponsors/rwxrob) is best, but Twitch cheers, subs, etc. are always welcome. You can also help by just providing positive comments with useful chapter markers on YouTube videos that need them. You can also just help people in the [discussion forum](https://github.com/rwxrob/boost/discussions).

## Overview

The content of the Boost itself is organized as follows:

1. Understand the tech industry (culture, lifestyle, companies, jobs)
2. Understand yourself (interests, strengths, challenges)
3. Discover your ideal job and companies that need it
4. Learn what your ideal job requires and how to prove you can do it
5. Develop the specific skills, knowledge, and abilities required by the job
6. Manage and maintain your tech career effectively

## Common technologies relevant to most tech jobs

What if you are still figuring out what your ideal job is? Consider learning any of the following right now while you figure it out. These apply universally to pretty much any tech job these days.

- AI prompting and web research
- Linux/Unix terminal (not desktop)
- Git (git) and GitHub (gh)
- Markdown and static web site creation
- Coding fundamentals (bash, Python, Go, JavaScript, AI vibing)
- Virtualisation (containers, virtual machines)
- Network fundamentals

Enterprise specific:

- Kubernetes
- Ansible
- Slurm

## Get started exploring Linux terminal

This repo allows you to immediately start exploring the Linux operating system and terminal command line interface (CLI) by running it safely in a container (configured in the [Containerfile](Containerfile)). When combined with a laptop, once you are setup, you can explore and learn anywhere — even without an Internet connection. Here's how to get started:

1. Get admin rights on your computer
2. Learn to install software on your computer
3. Install [WezTerm](https://wezterm.org) or another terminal emulator
4. Install [Podman](https://podman.io/) or another OCI container engine
5. Build and run this `boost` container locally

These steps are discussed below.

### Get admin rights on your computer

Even though Linux is contained and separated from your host computer, you still need to be able to install a few things on it as described in the following sections. This requires something called _admin rights_. If you did the setup on your computer yourself when you bought it then you likely already have these rights and don't necessarily need another account created. If not, you'll need to have the person who did setup your computer change your rights or add a new account with those rights on the computer that has them. There's no way around this.

> [!NOTE]
> If you cannot get admin rights because you are using a computer you do not control, say at work or school, you can find the person who can install things and only need to install two things for this Boost: a terminal emulator (i.e. WezTerm) and an OCI container engine (i.e. Podman). The Boost has been designed to only require these two things to make it easier for schools to use the Boost without violating school computer security policies.

### Learn to install software on your computer

Installing software is a basic computing skill, yet unfortunately very few people learn it because it isn't formally taught anywhere. Thankfully, there are many videos on how to do it (including ours).

There are two general ways to install stuff:

1) The point and click way (beginners)
2) The software package manager way (pros)

Most of the time just going to the web page of the thing you want to install and following the directions does the trick. This means downloading a thing and then clicking on it or clicking through the buttons that pop up. Pretty much everyone can do this once they have admin access. For now this is enough to install the stuff you need to get going, but you might want to take the time now to learn about the more professional method of software installation.

However, learning to use a software package manager is easier, more common, and preferred for professional work. The overall winner for Mac systems is `brew` and for Windows (these days) is `winget` (but Windows has many).

### Install WezTerm

The easiest method is to just go to <https://wezterm.org> and click on your computer operating system to download the regular installer and run it.

The Wezterm default configuration is pretty good, but if you want a head start at your own configuration (which requires knowing a bit of Lua language) [here's one](home/ubuntu/.wezterm.lua) that will work on any OS. Read the WezTerm help and instructions to know where to put the file.

### Install Podman

Like WezTerm, the easiest method to install Podman is just to go to <https://podman.io> and follow the instructions to download and run the installer. Again, you'll need admin on your computer to do the installation. You can use `brew` or `winget` if you know how (but don't sweat it if you don't).

### Build and run boost container

```sh
podman build -t boost https://github.com/rwxrob/boost.git
podman run -it --hostname boost --name boost boost
```

If you want to throw away the container after running just add `--rm` to the `run` command above. Otherwise, when you `exit` it will still be there and you can restart and attach later to continue your work until you explicitly remove it:

```sh
podman start -a boost
```

When you are ready to use the `man` pages consider running `sudo unminimize`.

### Remove container to get latest

Sometimes you will want to remove a previously created `boost` container and restart by rebuilding it (in the section above). Here's how.

> [!WARNING]
> You will lose all your work done in the `boost` container when you remove it unless you have saved it to GitHub or some other way. Use extreme caution.

```sh
podman rm boost
```

You are now able to create a fresh new `boost` container using the latest from this repo.

## Old beginner boost content

- https://rwxrob.github.io/skilstak/boost
