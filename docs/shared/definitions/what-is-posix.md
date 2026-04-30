# What is POSIX?

POSIX (Portable Operating System Interface) is a family of IEEE standards defining APIs, utilities, and shell behavior for Unix-like systems.

It specifies how software and scripts should work across compliant operating systems (Linux, macOS, BSD, etc.) so they're portable. Many features, like `pipefail`, are bash-specific and not POSIX.

Practical notes:
- `/bin/sh` aims for POSIX compliance
- If you need portability, write scripts to POSIX shell or declare a bash shebang (`#!/usr/bin/env bash`) when using bash-only features
- Most Linux distributions include GNU extensions beyond POSIX

Why it matters: understanding POSIX helps you write portable code that runs reliably across different Unix-like systems and containers.

**Reference:** [Beginner track](../../beginner/) → Phase 2: Setup
