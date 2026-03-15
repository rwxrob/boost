# Build, run, remove boost container

Containers are everywhere in IT these days. You must understand them. What better way than to learn how to use the Linux terminal from within a container designed to help beginners get started. Here's how.

> [!CAUTION]
> Containers are designed to be _ephemeral_ meaning they get recreated all the time and anything on them is reset. It's good to get into the mindset now when thinking about them. Nothing is saved. Source code is pulled dynamically from GitHub and other repos. Configuration is passed through environment variable and mounting things into the container. This property of containers is essential to understand when working with them in an enterprise IT setting.

## Build the boost container

[Open a terminal](open-a-terminal) on your computer and run the following commands. The first command pulls down the boost container source and builds it locally on your computer.

```sh
podman build -t boost https://github.com/rwxrob/boost.git
```

## Run a named boost container and interact with it

Then actually run the container giving it the name `boost`. If you want to immediately throw away the container instance after you are don you can add `--rm` as well. Otherwise, you can restart and connect to the same container later until you explicitly delete it.

```sh
podman run -it --hostname boost --name boost boost
```

When you are done type `exit` or just close the terminal. You can reconnect later with the following command.

## Reconnect to boost container to continue working later

```sh
podman start -a boost
```

When you are ready to use the `man` pages consider running `sudo unminimize`.

# Remove boost container to get latest

Sometimes you will want to remove a previously created `boost` container and restart by rebuilding it (in the section above). Here's how.

> [!WARNING]
> You will lose all your work done in the `boost` container when you remove it unless you have saved it to GitHub or some other way. Use extreme caution.

```sh
podman rm boost
```

You are now able to create a fresh new `boost` container using the latest from this repo.
