# Linux Beginner Boost

```sh
podman build -t boost https://github.com/rwxrob/boost.git
podman run -it --hostname boost --name boost -v shared://shared boost
```

To get man pages and such consider running `unminimize`.
