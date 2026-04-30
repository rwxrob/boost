# Gas Town: multi-agent workspace manager worth knowing

Most people running multiple AI agents hit the same wall: agents lose context when sessions end, coordination becomes chaos beyond four or five agents, and there's no reliable way to hand work between them.

Gas Town (`gastownhall/gastown`) solves this by treating your multi-agent workspace like a small town — with persistent identities, mailboxes, and git-backed state that survives crashes and restarts.

## Why context loss is the real problem

Every time an agent session ends, it forgets everything. You can partially work around this with AGENTS.md files and careful prompting, but at scale — running ten, twenty, or thirty agents across multiple projects — manual context management breaks down fast.

Gas Town stores work state in a "Beads ledger" backed by git worktrees. Agents (called Polecats) can spin up, do work, end their session, and pick up exactly where they left off next time. Nothing is lost.

## The architecture in plain terms

- **The Mayor** — an AI coordinator instance that manages the whole town; you tell it what you want and it delegates
- **The Town** — your workspace directory (e.g. `~/gt/`) containing all projects and agents
- **Rigs** — one per project; wraps a git repo and manages its agents
- **Polecats** — worker agents with persistent identity but ephemeral sessions; spawned for tasks, but their history survives
- **Hooks** — git worktree-based storage; this is what makes state persist across restarts
- **Mailboxes** — how agents communicate and hand work to each other without stepping on each other

## Why this matters at the beginner stage

You don't need Gas Town on day one. But understanding that this kind of tooling exists changes how you think about agentic development. Instead of thinking "I run one agent at a time," you start thinking in terms of fleets of agents with roles, persistent memory, and structured handoffs.

That mental model is where professional-grade agentic development is heading, and Gas Town is one of the clearest working examples of it today.

## How it compares

| Tool | Strength | Weakness |
|------|----------|----------|
| Gas Town | Persistent state, scales to 30+ agents | Early stage, evolving API |
| tmux + AI agents | Simple, no dependencies | Manual coordination |
| CrewAI | Role-based structure | No git-native state |
| AutoGen | Mature framework | Complex setup |
| obra/superpowers | Hooks + ops management | Narrower scope |

## Getting started

```sh
gh repo clone gastownhall/gastown
```

Read the README carefully — the Mayor-first workflow is the key insight. Start by asking the Mayor what it can do before trying to wire anything manually.

## See also

- [Git worktrees for parallel agentic development](git-worktrees-for-agentic-development.md)
- [Scripting and automating agents with tmux send-keys](tmux-for-multi-agent-development.md)
