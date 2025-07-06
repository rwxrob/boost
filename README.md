# Linux Beginner Boost

```sh
podman build -t boost https://github.com/rwxrob/boost.git
podman run -it --hostname boost --name boost -v shared://shared boost
```
