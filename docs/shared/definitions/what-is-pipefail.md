# What is pipefail?

Pipefail is a shell option (`set -o pipefail`) that makes a pipeline fail if any command in it fails.

By default, a pipeline (cmd1 | cmd2 | cmd3) returns only the exit status of the rightmost command. Failures in cmd1 or cmd2 are hidden. With pipefail enabled, the pipeline's exit status reflects the first command that fails.

Example:
- Without pipefail: `false | true` returns 0 (success) because the last command succeeded
- With pipefail: `set -o pipefail; false | true` returns 1 (failure) because false failed

Why use it: makes scripts detect errors reliably when chaining commands with pipes. Use in bash/zsh scripts; not POSIX in `/bin/sh`.

**Reference:** [Beginner track](../../beginner/) → Phase 2: Setup
