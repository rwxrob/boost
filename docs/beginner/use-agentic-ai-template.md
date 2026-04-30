# Use rwxrob/agentic-ai-template for new projects

Every new project you start should have a solid foundation for working with agentic AI from day one. The [rwxrob/agentic-ai-template](https://github.com/rwxrob/agentic-ai-template) is a minimal, agent-agnostic starting point designed exactly for this.

## What it includes

- A ready-to-use `AGENTS.md` file structured for AI agents to read at the start of every session
- A placeholder `README.md` that reminds you what to fill in
- Nothing else — by design

It is intentionally minimal. The goal is not to give you a framework or a pile of boilerplate. It is to give you the one file that matters most when working with agentic AI: a clear, maintained description of your project that an agent can pick up and run with.

## How to use it

1. Go to [github.com/rwxrob/agentic-ai-template](https://github.com/rwxrob/agentic-ai-template)
2. Click **Use this template** → **Create a new repository**
3. Name your repo and set visibility
4. Clone it locally and start working

```sh
gh repo create your-project-name --template rwxrob/agentic-ai-template --clone
cd your-project-name
```

## The most important file: AGENTS.md

`AGENTS.md` is how you communicate your project to any AI agent — GitHub Copilot, Claude, or anything else — without having to re-explain everything at the start of each session.

Keep it up to date with:

- What the project does and why it exists
- The current architecture and tech stack
- Key conventions, rules, and things the agent must never do
- Current versions of major dependencies
- What is in progress right now

A well-maintained `AGENTS.md` means every session picks up exactly where the last one left off. A stale one means repeating yourself constantly and getting inconsistent results.

## Why this matters

When you operate like a startup — building fast, shipping real things, iterating on feedback — you cannot afford to spend the first ten minutes of every session re-onboarding your tools. The template makes the context persistent so you can stay in motion.

Start every project this way. It takes two minutes and saves you hours.
