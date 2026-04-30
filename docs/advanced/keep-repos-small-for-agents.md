# Keep repos small so you can fire bad agent sessions fast

One of the most underappreciated aspects of working with agentic AI is what happens when a session goes wrong — and sessions do go wrong. An agent misunderstands the goal, goes down a bad path, makes a series of compounding mistakes, or simply accumulates so much confused context that its output becomes unreliable. When that happens, the right move is fast: kill the session and start a fresh one.

How fast you can do that depends almost entirely on how large your codebase is.

## The cost of a large repo in the agentic era

When you start a new agent session on a large repository — hundreds of files, deep directory trees, interleaved concerns — the agent has to orient itself. It reads file structures, loads relevant files into context, builds a mental model of the codebase. This takes time and tokens, and the result is a context window that's already partially full before the agent has done any useful work.

When that session goes bad and you need to restart, you pay that cost again. And again. Every bad hire costs you the full onboarding fee.

Large repos also make it harder to write a good AGENTS.md. When a codebase does many things, the instructions needed to guide an agent reliably become long, complex, and fragile. A new session has more to absorb before it can work safely.

## Small repos let you fire and hire instantly

A small, focused repo — one that does one thing, has a clear purpose, and a concise AGENTS.md — can be handed to a fresh agent in seconds. The agent reads a handful of files, understands the scope immediately, and gets to work. No long onboarding. No confusion about what's in scope.

When that agent session goes sideways, you discard it without grief. The replacement agent comes up to speed just as fast. The turnaround from bad session to productive session is measured in seconds, not minutes.

This is the agentic equivalent of the Unix philosophy: small tools that do one thing well, composed together. Repos that do one thing hand off cleanly to agents. Repos that do everything create agents that know a little about a lot and are confused about most of it.

## Monorepos are particularly risky

Monorepos — single repositories containing many services, packages, or applications — were already a tradeoff before agentic AI. With agents they become genuinely hazardous. An agent working in a monorepo can wander into unrelated code, make changes in the wrong package, or receive contradictory instructions from files in different parts of the tree.

The larger the repo, the more likely an agent is to get lost, and the more expensive it is to restart when it does.

## What to do instead

Keep each repo focused on a single responsibility. If a project grows to the point where different parts of it require different context, different instructions, or different expertise — split it. The overhead of managing multiple small repos is trivial compared to the cost of nursing a confused agent through a codebase it can barely hold in mind.

Use AGENTS.md to define scope explicitly. Tell the agent what the repo does, what it doesn't do, and where the boundaries are. A well-scoped AGENTS.md is a force multiplier. A vague one covering a sprawling codebase is noise.

## The deeper principle

Agents are cheap to start and cheap to discard. That's their superpower — you can run many of them, in parallel, each focused on a narrow task. The repos that take best advantage of this are the ones built to be handed off: small, clear, well-documented, and easy to pick up cold.

Build your repos for the agent that's never seen them before. That agent might be starting in five seconds.
