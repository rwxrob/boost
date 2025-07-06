# Linux Beginner Boost

1. Install a terminal (preferably [WezTerm](home/ubuntu/.wezterm.lua))
1. Install Podman (or Docker if you must)
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
