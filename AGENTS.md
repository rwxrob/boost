# Agent Instructions

Keep this file as agent-agnostic as possible.

## Rule: Commits

- Always use conventional commits (e.g. `feat:`, `fix:`, `docs:`, `chore:`)
- Never include any reference to AI agents or Copilot in commit messages, co-authorship, documentation, or any other files in this repository
- Committing directly to main is okay in this repo

## Rule: Context Maintenance

At the end of every significant task or session, summarize the current state, architectural decisions made, and pending todo items into AGENTS.md. Always ensure this file reflects the ground truth of the project so future sessions can resume without friction. Use the writeFile tool to overwrite it so the next session starts with current state.
Regularly check the `rwxrob/agentic-ai-template` repo's AGENTS.md and ensure any updated rules have been merged into this file.

## Rule: Environment

- Use `/usr/bin/open` (full path) to open files or URLs on macOS — never plain `open`
- Open markdown files using the Obsidian URL scheme: `/usr/bin/open "obsidian://open?path=URL_ENCODED_ABSOLUTE_PATH"`
- Always pipe URLs or other text to be pasted into `pbcopy` when suggesting them

## Rule: Secrets

- Never commit secrets, config files, or database files

## Rule: Writing tone

- Avoid hyperbolic and superlative language; state facts plainly
- Reference current events and concrete examples whenever possible

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

This repo is the Tech Career Boost — a comprehensive tech career guide for everyone: absolute beginners, mid-career professionals, and veterans.

**Scope:** Not just for first-time learners, but anyone who needs a boost in their tech career — including latest industry releases (e.g., Gravano), advanced patterns, agentic development, and cutting-edge tools.

**Structure:** `docs/` organized into three tracks:
- `docs/beginner/` — foundational path (career exploration, setup, core skills)
- `docs/advanced/` — next-level content (MCP, agents, orchestration, git workflows)
- `docs/shared/` — cross-track resources (infrastructure, community, contribution)
- `docs/shared/definitions/` — searchable, hotlinkable term definitions (one file per definition, "what-is-" naming)

**Navigation:** Main `docs/README.md` is a hub directing readers to Beginner, Advanced, or Shared resources.

## Current state

### Phase 1: Restructure complete

Reorganized 40+ existing markdown files into three tracks. Created:
- `docs/beginner/README.md` — guided path emphasizing "know thyself" first
- `docs/advanced/README.md` — MCP servers, multi-agent, CI/CD patterns
- `docs/shared/README.md` — tools, infrastructure, community
- `docs/shared/definitions-index.md` — index of searchable definitions
- `docs/shared/definitions/what-is-*.md` — individual definition files (pipeline, pipefail, POSIX, shebang, etc.)

### Recent changes

- Added `docs/mcp-servers-cloud-and-local.md` — practical MCP deployment patterns with Jira example
- Updated `docs/mandatory-commands.md` — added `column -t` flag
- Scope shift: dual-track for all career levels, not just beginners

## Pending todos

- Migrate remaining git-basics Q&A content into individual definition files with "what-is-" titles
- Link all definitions hotlinkably from reference docs
- Test cross-reference links work in Obsidian, on GitHub, and in rendered markdown
- Create advanced definitions (pipeline, serverless, container, IaC, AAIF, MCP, etc.)
