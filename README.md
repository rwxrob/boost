# Linux Beginner Boost

The *Linux Beginner Boost* is a quick start into the Linux world for absolute beginners with little or no Linux, terminal, or coding experience. This repo contains a Linux [Containerfile](Containerfile) that has a bit of everything to have fun while sampling and learning. Check in with [Rob's socials](https://linktr.ee/rwxrob) to see when where he's live and come say hello.

1. Install a terminal (preferably [WezTerm](https://wezterm.org))
1. Install [Podman](https://podman.io/) (or [Docker](https://docker.com) if you must)
1. Build and run the `boost` container locally

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
