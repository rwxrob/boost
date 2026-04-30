# What is a shebang?

A shebang is the first line of a script that tells the OS which interpreter to use. It starts with `#!` followed by the interpreter path.

Examples:
- `#!/usr/bin/env bash` — portable; finds bash in PATH
- `#!/bin/sh` — POSIX shell
- `#!/usr/bin/env python` — Python

When you run a script directly (`./script`), the kernel uses the shebang to launch the specified interpreter. If you invoke the interpreter explicitly (`bash script`), the shebang is ignored.

Steps:
1. Add shebang as first line
2. Make the file executable: `chmod +x script`
3. Run it: `./script`

Use `env` for portability across systems (finds interpreters in PATH rather than hard-coding paths).

**Reference:** [Beginner track](../../beginner/) → Phase 2: Setup
