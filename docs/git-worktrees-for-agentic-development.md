# Git worktrees for parallel agentic development

When you run multiple AI agents simultaneously in tmux windows, you need each agent to have its own working directory. Git worktrees make that possible without cloning the repository multiple times.

## What a worktree is

A git worktree is a separate checkout of a branch from the same repository. Each worktree has its own directory, its own working tree, and its own branch — but they all share the same git history and object store. No duplication, no separate clones, no confusion about which remote to push to.

```sh
git worktree add ../project-feature-auth feature-auth
git worktree add ../project-feature-api feature-api
git worktree add ../project-bugfix-login bugfix-login
```

Now you have three directories, each on a different branch, all backed by the same repository.

## Why this matters for agents in tmux

Each tmux window gets its own worktree directory. The agent in that window has a clean, isolated branch to work on. It can make commits, run tests, and modify files without touching anything another agent is working on. No stashing, no branch switching, no stepping on each other.

Without worktrees, running multiple agents on the same checkout is asking for chaos — one agent's uncommitted changes collide with another's, files get overwritten mid-task, and debugging what went wrong is painful.

## The isolation problem and why rebase is the answer

Worktree isolation is the feature, but it creates a problem: branches diverge. If the auth agent and the api agent both modify shared utility code, they each have a version the other does not know about. Left unchecked, this compounds into a merge conflict that wipes out the time you saved running them in parallel.

The solution is to rebase frequently — bring each worktree's branch up to date with main (or the integration branch) before divergence accumulates:

```sh
cd ../project-feature-auth
git fetch origin
git rebase origin/main
```

Rebase replays your branch's commits on top of the latest main, keeping history linear and conflicts small. Do this before starting a long agent task and again when you are ready to merge. A rebase on a branch that has been diverging for an hour is much easier than one that has been diverging for a day.

## A practical multi-agent workflow

1. Create a worktree per task before assigning work to an agent
2. Name each tmux window after its worktree and task — `⏳ feature-auth`, `⏳ feature-api`
3. Start each agent in its own worktree directory
4. Rebase against main at the start of each session and before merging
5. When an agent finishes, merge or rebase its branch into main, then delete the worktree

```sh
# Clean up a finished worktree
git worktree remove ../project-feature-auth
git branch -d feature-auth
```

## Keeping agents synchronized mid-task

If one agent's work is a dependency for another — say the api agent needs a type the auth agent just defined — fetch and rebase the dependent branch before the second agent continues:

```sh
cd ../project-feature-api
git fetch . refs/heads/feature-auth:refs/heads/feature-auth
git rebase feature-auth
```

This pulls the auth branch's commits directly into the api worktree without going through a remote. Fast and local.

## This is less complex than it looks — just ask your agent

The commands above look involved but you never have to type most of them. Ask your agent and it handles the whole setup. A prompt like this is enough:

> Create a new git worktree and branch called `feature-auth`, open a new tmux window for it, and start working on issue #42 — create a draft PR when you have something to show.

The agent will run `git worktree add`, create the branch, open the tmux window, change into the directory, start work, and open the PR with `gh pr create --draft`. You stay in your current window while a new agent spins up in a new one, fully isolated, already working.

## List all worktrees at any time

```sh
git worktree list
```

This shows every active worktree, its path, and which branch it is on — a useful sanity check before spinning up a new agent session.

## The rule

One agent, one worktree, one branch. Rebase early and rebase often. Merge when done, delete the worktree, move on. This keeps parallel agent work fast, clean, and recoverable.

**See also:** [tmux is the best multi-agent development environment](tmux-for-multi-agent-development.md)
