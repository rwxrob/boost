# Linux Beginner Boost

1. Install a terminal (preferably WezTerm)
1. Install Podman (or Docker if you must)
1. Build and run the `boost` container locally

```sh
podman build -t boost https://github.com/rwxrob/boost.git
podman run -it --hostname boost --name boost boost
```

To restart and attach later:

```sh
podman start -a boost
```

When you are ready to use the `man` pages consider running `sudo unminimize`.
