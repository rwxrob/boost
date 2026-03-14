# Build and run boost container

Open a terminal on your computer and run the following commands. The first command pulls down the boost container source and builds it locally on your computer.

```sh
podman build -t boost https://github.com/rwxrob/boost.git
```

Then actually run the container giving it the name `boost`. If you want to immediately throw away the container instance after you are don you can add `--rm` as well. Otherwise, you can restart and connect to the same container later until you explicitly delete it.

```sh
podman run -it --hostname boost --name boost boost
```

When you are done type `exit` or just close the terminal. You can reconnect later with the following command.


```sh
podman start -a boost
```

When you are ready to use the `man` pages consider running `sudo unminimize`.

Sometimes you might want to [remove it](remove-boost.md) to reset and try again or get the latest.
