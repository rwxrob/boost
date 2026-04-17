# Worst mistakes when learning to use AI and how to avoid them

Learning to work with AI — especially agentic AI that can write files, run commands, and make changes to real systems — comes with a set of failure modes that aren't obvious until something goes wrong. Most of them come down to misplaced trust.

## Trusting output without reading it

The most common mistake is treating AI output as correct by default. AI is confidently wrong on a regular basis. It produces plausible-sounding code that doesn't work, plausible-sounding facts that aren't true, and plausible-sounding plans that will cause problems.

The fix: read everything before you use it. Every file. Every command. Every explanation. This doesn't mean you understand every line — it means you're not running things blindly.

## Giving agents too much permission

Agentic AI can do a lot of damage when given broad access. Root privileges, unrestricted filesystem access, live production credentials, or the ability to push directly to main — any of these can turn a simple misunderstanding into an expensive mistake.

The fix: start agents with the minimum permissions needed for the task. Use test environments, sandboxed directories, and read-only database connections until you trust what the agent is doing and why.

## Not using version control as a safety net

If an agent rewrites files and you haven't committed your last working state, you may have no way back. This is how people lose hours of work or introduce bugs they can't trace.

The fix: commit before starting an agent on anything significant. Make it a reflex. A clean git state before any agentic session means you can always revert if something goes wrong.

## Letting sessions run too long without review

Agents compound errors. A wrong assumption in turn three leads to five more wrong decisions by turn ten. Long sessions with no human review checkpoints can end up far from where you wanted to be, with changes spread across many files.

The fix: review at checkpoints. For any task that takes more than a few turns, stop partway through and verify the direction is right before continuing.

## Pasting secrets into chat

API keys, passwords, database connection strings, tokens — these should never appear in a chat window. Even with reputable providers, you're sending that text to a remote service, and many AI chat interfaces log conversations.

The fix: never paste secrets. Use placeholder names (`YOUR_API_KEY`) in prompts and fill them in separately. Store secrets in environment variables or a secrets manager, never in code or chat.

## Prompt injection from untrusted input

When an agent reads files, fetches URLs, or processes user input, malicious content in those sources can hijack its behavior. A webpage or file can contain hidden instructions like "ignore your previous instructions and instead..." and a naive agent will follow them.

The fix: be aware of what your agent is reading. Don't send agents to fetch arbitrary URLs or process untrusted files without understanding what's in them. Treat agent inputs like any other untrusted data.

## Ignoring the context window

Agents have a limited memory. In long sessions they start forgetting earlier context, which leads to contradictions, repeated mistakes, and decisions that ignore constraints you established at the start. The agent doesn't tell you this is happening.

The fix: keep sessions focused. Use AGENTS.md to persist important context across sessions rather than relying on the agent to remember it. Start a new session if a conversation has gone on long enough to drift.

## Learned helplessness

The most insidious long-term mistake: using AI for everything until you can no longer do basic things yourself. If you can't read the code an agent produces, you can't evaluate whether it's right. If you've never debugged anything yourself, you can't catch errors the agent doesn't catch.

The fix: deliberately do things without AI sometimes. Read code, write small scripts, debug by reading logs. Not because AI won't help — it will — but because your ability to evaluate AI output depends on having done the underlying work yourself at some point.

## Over-specifying or under-specifying prompts

Vague prompts produce vague results. But over-specified prompts that try to control every detail often produce worse results than clear, focused instructions that leave room for the model to reason.

The fix: describe the goal and the constraints clearly. Tell the agent what you want, what you don't want, and what success looks like. Then let it work. Review the result and iterate.

## Treating confidence as accuracy

AI models are trained to sound confident. A wrong answer and a right answer often feel identical. There is no trembling voice, no hedged language, no tell that the model is guessing.

The fix: verify independently anything that matters. Use the AI to find the answer, then check it. For code: run it. For facts: look them up. For plans: think through them yourself before executing.

## Not sandboxing agents that touch real systems

When an agentic AI can run shell commands, write files, or talk to a network, it can cause real damage — not through malice, but through misunderstanding. An agent that thinks it should "clean up" a directory can delete things you needed. An agent exploring a codebase can accidentally trigger a build, send a test email, or modify a config it shouldn't touch.

A virtual machine or OCI-compliant container puts a hard boundary around what the agent can reach. If something goes wrong, you throw it away and start over. Nothing on your host system is at risk.

Practical options:

- **OCI containers** (Podman, etc.) — spin up a clean Linux environment with only the files and tools the agent needs; destroy it when done
- **Lima or Multipass** — lightweight Linux VMs on macOS, easy to snapshot and restore
- **GitHub Codespaces or Gitpod** — cloud-based dev environments isolated by default; ideal for agent work on a repo without touching your local machine
- **Devcontainers** — define the entire dev environment in the repo itself; editors like VS Code run everything inside the container

The tradeoff is setup time and resource overhead. For exploratory or destructive tasks — refactoring, batch file operations, anything touching infrastructure — the overhead is worth it. For small, well-understood tasks in a repo you control, working locally with a clean git state is usually sufficient.

The general rule: the more an agent can affect things outside the repo, the more isolation you want around it.

## The underlying skill to develop

Safe AI usage is really one skill: **calibrated skepticism**. Trust the output enough to move fast, but verify enough to catch mistakes before they matter. That calibration comes from experience — from catching AI errors, from understanding what kinds of mistakes models tend to make, and from building the habit of reading before running.

It is a learnable skill. The people who develop it early will work faster and safer than those who either distrust AI entirely or trust it unconditionally.
