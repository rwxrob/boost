# Build, run, remove boost container

Containers are everywhere in IT these days. You must understand them. What better way than to learn how to use the Linux terminal from within a container designed to help beginners get started. Here's how.

> [!CAUTION]
> Containers are designed to be _ephemeral_ meaning they get recreated all the time and anything on them is reset. It's good to get into the mindset now when thinking about them. Nothing is saved. Source code is pulled dynamically from GitHub and other repos. Configuration is passed through environment variable and mounting things into the container. This property of containers is essential to understand when working with them in an enterprise IT setting.

## Install WSLv2 on Windows

If you have Windows you likely need to install WSLv2 in order to use podman and containers. This is an extra step but is actually a good thing since WSL is _actually_ Linux and not Linux running inside a container. This will be useful later when you want to use Linux from your Windows computer without 
throwing it away and running Linux at the same time.

> [!NOTE]
> Mac users do not need to do anything like this because macOS is already UNIX (which some consider far superior to Linux in general).

1. [Open a terminal](open-a-terminal)
2. `wsl --update` (or `wsl --install` if required)

## Initialize podman virtual machine

This creates a tiny version of Linux in which to run containers like the `boost` container. [Open a terminal](open-a-terminal) and run the following:

```sh
podman machine init
podman machine start
```

## Build the boost container

Now we need to build the boost container from the source pulled from `github.com/rwxrob/boost`.

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
