# Agent Instructions

Keep this file as agent-agnostic as possible.

## Rule: Commits

- Always use conventional commits (e.g. `feat:`, `fix:`, `docs:`, `chore:`)
- Never include any reference to AI agents or Copilot in commit messages, co-authorship, documentation, or any other files in this repository
- Committing directly to main is okay in this repo

## Rule: Context Maintenance

At the end of every significant task or session, summarize the current state, architectural decisions made, and pending todo items into AGENTS.md. Always ensure this file reflects the ground truth of the project so future sessions can resume without friction. Use the writeFile tool to overwrite it so the next session starts with current state.

## Rule: Environment

- Use `/usr/bin/open` (full path) to open files or URLs on macOS — never plain `open`
- Open markdown files using the Obsidian URL scheme: `/usr/bin/open "obsidian://open?path=URL_ENCODED_ABSOLUTE_PATH"`
- Always pipe URLs or other text to be pasted into `pbcopy` when suggesting them

## Rule: Secrets

- Never commit secrets, config files, or database files

## Rule: Code style

- Single-line paragraphs in all markdown files — no multi-line wrapped paragraphs
- No underscores or spaces in filenames; use hyphens
- No extensions on executable scripts, ever
- Always use sentence case for all titles and headings
- Always use `----` for horizontal rules in markdown (not `---`) to avoid em-dash creation by pandoc and similar tools

## Rule: Containers

- Never refer to Docker specifically; always refer to OCI-compliant containers

## Rule: GitHub repos

- Always enable branch protection on `main` when creating a new GitHub repo so a PR is always required
- Exception: repos beginning with `notes-` are knowledge repos, not source code — commit directly to main with no branch protection and no PRs required
- Require zero reviews (PRs required but no approvals needed)
- Enable automatic branch deletion after a PR is merged
- After merging a PR, pull the latest changes into the current branch and delete any leftover worktrees

## Rule: Agent specific

- Always use `gh copilot` not `copilot`
- Use Claude Opus 4.6 (high) for all new projects
- When renaming a tmux window, always prefix the name with an emoji reflecting current state: `🚀` starting, `⏳` in progress, `🔍` researching, `🔨` building, `🧪` testing, `✅` done, `🔴` error, `💬` waiting for input, `🛑` blocked

## Current architecture

This repo is the rwxrob Beginner Boost — a beginner tech career guide published as a docs vault in `docs/` using Obsidian-compatible markdown. The outline lives in `docs/README.md`. Individual topic pages live as separate markdown files in `docs/`.

## Current state

The outline in `docs/README.md` covers sections 1–12. Recent additions include:

- `docs/build-mcp-servers.md` — article on building MCP servers as a modern tech skill, linked under section 10 (Harness the power of AI) as item 10.5
- `docs/voice-input-with-superwhisper.md` — article on using Superwhisper for live voice input with GitHub Copilot CLI, linked under section 10 as item 10.6
- `docs/tmux-for-multi-agent-development.md` — linked under section 6.7.4
- `docs/use-agentic-ai-template.md` and `docs/agentic-template-is-the-new-dotfiles.md` — linked under section 12

## Pending todos

- Fill in `## Current tags / versions` section once versioning is established
