# What is a pipeline?

A pipeline is an automated CI/CD (continuous integration, continuous deployment) sequence that runs on code changes.

Typical flow: build → test → lint → package → deploy. Triggered by pushes or pull requests, run by executors/runners, produce artifacts and logs.

Common tools: GitHub Actions, GitLab CI, Jenkins, CircleCI. Each job in a pipeline may depend on others; if any stage fails, the pipeline stops.

Why it matters: catch bugs early, enforce code quality standards, and deploy reliably without manual steps.

**Reference:** [Beginner track](../../beginner/) → Phase 3: Build core skills
