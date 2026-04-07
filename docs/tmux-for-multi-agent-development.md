# tmux is the best multi-agent development environment

There are shiny new tools designed for running multiple AI agents at once. Conductor gives you a dashboard. T3 Chat gives you a polished interface. They look great in demos.

tmux is better. Here is why.

## screen compatibility is not an accident

The first thing in [rwxrob/dot tmux config](https://github.com/rwxrob/dot/tree/main/tmux) is rebinding the prefix from `C-b` to `C-a` — the same default used by GNU screen. This is not nostalgia. It is intentional compatibility with decades of muscle memory and the fact that screen is still installed on nearly every Linux system you will ever SSH into. Learn `C-a` once and it works everywhere, on everything, forever.

## C-a w is the only window manager you need

The single most important keybinding for multi-agent development is `C-a w` — the interactive window selector. It drops you into a live list of every window in your session. You navigate with vim keys (`j`/`k`), you see each window's name and status at a glance, and you jump to any agent instantly. No mouse. No clicking around a dashboard. No switching browser tabs. One keystroke gives you a real-time map of everything running.

When you have six agents running in parallel — one per feature branch, one watching tests, one tailing logs — `C-a w` is how you stay oriented. Nothing in Conductor or T3 Chat comes close to the speed and clarity of that single interaction.

## The fundamental problem with GUI-based multi-agent tools

When you use a GUI to orchestrate multiple agents, you are putting a layer of abstraction between you and your agents that costs you in every direction — visibility, control, speed, and portability.

GUI tools make decisions for you. They decide how much context each agent gets. They decide how agents communicate. They decide what output you see and when. They decide what is possible.

When something goes wrong — and it will — you are debugging through glass.

## tmux gives you the terminal, which means everything

With tmux, each agent runs in a real terminal pane or window. You see exactly what it sees. You type exactly what you want. You pipe output from one agent directly into another. You script the whole thing in bash if you want to.

This is not a workaround. This is the Unix philosophy applied to multi-agent development: small tools, full control, chain them together however you need.

## What tmux multi-agent actually looks like

Name each window after the task it owns. One window runs the backend agent, one runs the frontend agent, one runs tests in a watch loop, one tails logs. Use `C-a w` to jump between them. Use `C-a |` to split a window horizontally when you want to watch two things at once. Use `C-a y` to sync input across all panes when you want to send the same command to every agent simultaneously.

You can also drive agents programmatically with `tmux send-keys`:

```sh
tmux send-keys -t agents:backend "review the auth module for security issues" Enter
```

That one line from a script can drive an agent without any GUI, any API wrapper, or any third-party service.

## Persistence that actually works

tmux sessions survive disconnects, closed laptops, and SSH timeouts. Detach with `C-a d`. SSH in from another machine. Reattach with `tmux attach`. Your agents are exactly where you left them, still running, still thinking.

GUI tools live in a browser tab. Close the tab and the context is gone. T3 Chat does not survive a network blip. Conductor does not run on a headless server.

## Portability

tmux runs everywhere. Your laptop, a remote server, a container, a VPS. If it has a terminal, it has tmux — or at minimum, screen, which works the same way with the same prefix. The skills transfer completely.

## Scripting and automation

Because tmux is just a terminal multiplexer, every agent interaction is scriptable. Write a bash script that spins up a named session, opens a window per agent, assigns each one a task, and tears it all down when done. Trigger it from a webhook. Integrate it into CI. Schedule it with cron.

Try doing any of that with a dashboard.

## Get the config

The tmux config at [rwxrob/dot](https://github.com/rwxrob/dot/tree/main/tmux) is public domain. Clone it, study it, make it yours. It uses `C-a` as prefix, vi keys for navigation and copy mode, minimal status bar, screen-compatible splits, and synchronized panes on demand. It is the result of years of real use and is optimized for exactly this kind of work.

## The career signal

Developers who orchestrate agents from a tmux session are operating at a different level than developers who click around a dashboard. The tmux developer understands what is actually happening. They can script it, reproduce it, automate it, and explain it. The dashboard developer is dependent on the dashboard.

Learn tmux. Learn `C-a w`. Run your agents in named windows. Own your workflow.
