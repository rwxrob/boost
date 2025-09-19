# Linux Beginner Boost

The *Linux Beginner Boost* is a quick start into the Linux world for absolute beginners with little or no Linux, terminal, or coding experience. All you need is a computer (we assume Mac or Windows since you are just starting out), a terminal (we use WezTerm), and a container engine (like Podman or Docker).

This repo contains a Linux [Containerfile](Containerfile) that contains a pre-configured Ubuntu Linux OS for sampling and learning.

Check in with [Rob's socials](https://linktr.ee/rwxrob) to see when where he's live and come say hello.

Here's the process to get started:

1. Install [WezTerm](https://wezterm.org) (or another) terminal
1. Install [Podman](https://podman.io/) (or another) OCI container engine
1. Build and run this `boost` container locally

## Install WezTerm

The easiest method is to just go to <https://wezterm.org> and click on your computer operating system to download the regular installer and run it. Note that you will need admin rights on your computer to install anything. If you know of other ways (including `brew` for Mac or `winget` for Windows) go for it.

The Wezterm default configuration is pretty good, but if you want a head start at your own configuration (which requires knowing a bit of Lua language) [here's one](home/ubuntu/.wezterm.lua) that will work on any OS. Read the WezTerm help and instructions to know where to put the file.

## Install Podman

Like WezTerm, the easiest method to install Podman is just to go to <https://podman.io> and follow the instructions to download and run the installer. Again, you'll need admin on your computer to do the installation. You can use `brew` or `winget` if you know how (but don't sweat it if you don't).

## Build and run boost container

```sh
podman build -t boost https://github.com/rwxrob/boost.git
podman run -it --hostname boost --name boost boost
```

If you want to throw away the container after running just add `--rm` to the `run` command above. Otherwise, when you `exit` it will still be there and you can restart and attach later to continue your work until you explicitly remove it:

```sh
podman start -a boost
```

When you are ready to use the `man` pages consider running `sudo unminimize`.

## Beyond Boost

- Customizing [WezTerm](home/ubuntu/.wezterm.lua)
- Creating and customizing containers
- Kubernetes and "cloud native"
- Virtual machines and Proxmox
- Editor custom configuration
- ...
