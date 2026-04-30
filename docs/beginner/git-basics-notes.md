# Git basics — a quick reference

A git branch is a movable pointer to a commit used to work on a separate line of development.

A commit is a recorded snapshot of the repository at a point in time and is identified by a SHA hash.

A repository is the version-controlled project that contains the working tree, commits, branches, and refs.

Version control records changes over time so you can view history, revert mistakes, and collaborate safely.

Why use branches: isolate features, experiment safely, and collaborate without blocking the main history.

Why commit often: small, focused commits make review and debugging easier.

Common commands:
- git status — shows staged and unstaged changes, untracked files, and branch state. Run it frequently to see what will be committed.
- git add <files> — stage changes to be included in the next commit. Use 'git add -p' to interactively stage hunks.
- git commit -m "message" — record staged changes as a commit with a descriptive message. Make small, focused commits.
- git log — view commit history; use --oneline or --graph for compact views.
- git branch — list and manage branches; use 'git branch <name>' to create.
- git switch <branch> — change branches safely; use 'git switch -c <name>' to create and switch.
- git push -u origin <branch> — publish a local branch to remote and set upstream for future pushes/pulls.
- git pull --ff-only — update from upstream only if fast-forward is possible; prevents unexpected merge commits.

Safety tips:
- Use descriptive, short branch names (feature/desc, fix/desc).
- Commit small, focused changes with clear messages.
- Push branches to the remote for backup and collaboration.
- Use pull --ff-only to avoid unexpected merge commits when you want a linear history.

Further reading:
- git-scm.com/book  
- the Pro Git book  

Note: keep this file in your notes so you can revisit these basics as you learn more.
