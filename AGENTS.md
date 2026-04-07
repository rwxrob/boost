# Agent Instructions

## Rule: Commits

- Always use conventional commits (e.g. `feat:`, `fix:`, `docs:`, `chore:`)
- Never include any reference to AI agents or Copilot in commit messages, co-authorship, documentation, or any other files in this repository
- Committing directly to main is okay in this repo

## Rule: Context Maintenance

At the end of every significant task or session, summarize the current state, architectural decisions made, and pending todo items into AGENTS.md. Always ensure this file reflects the ground truth of the project so future sessions can resume without friction.

## Rule: Environment

- Use `/usr/bin/open` (full path) to open files or URLs on macOS — never plain `open`
- Open markdown files using the Obsidian URL scheme: `/usr/bin/open "obsidian://open?path=URL_ENCODED_ABSOLUTE_PATH"`

## Rule: Secrets

- Never commit secrets, config files, or database files

## Rule: Code Style

- Single-line paragraphs in all markdown files — no multi-line wrapped paragraphs
- No underscores or spaces in filenames; use hyphens
- No extensions on executable scripts, ever

## Current architecture

This repo is the rwxrob Beginner Boost — a beginner tech career guide published as a docs vault in `docs/` using Obsidian-compatible markdown. The outline lives in `docs/README.md`. Individual topic pages live as separate markdown files in `docs/`.

## Current tags / versions

TODO
