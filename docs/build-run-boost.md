# Build and run boost container

```sh
podman build -t boost https://github.com/rwxrob/boost.git
podman run -it --hostname boost --name boost boost
```

If you want to throw away the container after running just add `--rm` to the `run` command above. Otherwise, when you `exit` it will still be there and you can restart and attach later to continue your work until you explicitly remove it:

```sh
podman start -a boost
```

When you are ready to use the `man` pages consider running `sudo unminimize`.

Sometimes you might want to [remove it](remove-boost.md) to reset and try again or get the latest.
