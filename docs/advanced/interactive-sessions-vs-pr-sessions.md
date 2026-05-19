# Interactive AI sessions vs. long-running PR sessions

Two fundamentally different ways to work with AI agents: live interactive conversations where you steer in real-time, versus submitting a comprehensive prompt that runs unattended and opens a PR. Each has tradeoffs—cost, control, iteration speed—and knowing when to use each is critical for effective AI-assisted development.

[Watch on YouTube](https://youtu.be/7WgHJ5eoWkU)

## The two models

### Long-running PR sessions

Submit a well-constructed prompt with a detailed plan. The AI creates an issue, opens a PR, works through it (sometimes taking 30 minutes to an hour), and returns with completed work. Historically, this counted as a single API request—a huge cost advantage.

**Advantages:**
- Lower cost per session (historically—this is changing)
- Requires careful upfront planning, which leads to better prompts
- You don't need to babysit the work
- The issue and PR description get refined by the AI into clarified prompts
- Good for fire-and-forget tasks where steering isn't needed

**Disadvantages:**
- Can't steer mid-process if the AI goes wrong
- Requires very thorough initial prompting
- No human checkpoints to catch errors early
- Cost advantage disappearing with token-based pricing (June 2026+)

### Interactive sessions

Open a terminal session with an AI (Claude/Haiku via GitHub Copilot), use voice input, and interact in real-time. The AI makes changes, you see them immediately, you can correct course on the fly.

**Advantages:**
- Full control and steering capability
- Immediate feedback—see what's happening
- Much cheaper if you use Haiku (0.33x the cost of other models)
- Explicit memory control via AGENTS.md—you decide what the AI remembers
- Great for iterative work: planning, budgeting, live edits, personal repos
- Better UX with voice input (Superwhisper integration)
- You control all the context—no black-box memory like ChatGPT
- Perfect for operational tasks where immediate feedback matters

**Disadvantages:**
- Requires your active attention
- Can drift if the session runs too long
- Needs good terminal skills and workflow setup

## When to use each

### Use interactive sessions for:
- Anything you want to save and iterate on (health tracking, budget planning, travel planning)
- Personal projects and repos where you benefit from live steering
- Making small, precise edits to existing code or docs
- Building your own infrastructure or tools
- Learning—the AI can explain as it goes
- Operational tasks requiring immediate feedback: managing Kubernetes clusters, backing up systems, root cause determination, incident response
- Live troubleshooting where you need to see output between commands and adjust

### Use PR sessions for:
- Large features or refactors that don't need mid-process steering
- Work that's complex enough that you'd struggle to explain it live
- Scheduled/automated work (webhooks, CI triggers)
- Code you don't need to deeply understand yet (the PR lets you review it)
- Tasks where a single comprehensive prompt is genuinely better than iteration

## The cost math (changing)

**Before June 2026:** PR sessions cost way less because they counted as one request. This made the upfront planning investment worthwhile.

**After June 2026:** Tokens-based pricing changes everything. A long interactive session with Haiku might actually be cheaper than a massive PR session, especially if you're not paying per-request anymore. GitHub's token model "definitely needed" to happen, but the old single-request advantage is gone.

## Memory and control

ChatGPT remembers you—but it decides what to remember based on proprietary algorithms you can't control. Terminal-based interactive sessions give you explicit memory:

```
Store memory in AGENTS.md for persistence across sessions

Update AGENTS.md at the start of each session so your agents 
catch up and regain full context before doing real work.
```

You can explicitly tell Haiku: "Make sure you remember that this is primarily influenced by somebody who's seeking to become a fully remote digital nomad." The AI will confirm it stored this. It goes into the repo as documented context. Every future agent session reads it and knows your perspective.

With ChatGPT, you have no visibility into what it remembers or how. This is why terminal-based interactive sessions with explicit memory (AGENTS.md) are more reliable for work that matters.

## How the PR session flow works

1. Write a comprehensive prompt (or speak one via Superwhisper)
2. AI clarifies/restates your request (this is the prompt refinement)
3. AI creates an issue with the clarified prompt
4. Another AI (or the same one) reads the issue and breaks it into a detailed plan
5. The plan shows up in the PR before you approve it
6. You review the PR description and files, then approve
7. Regression testing validates it works (this is mandatory for agentic work)

The restatement at each step—prompt → issue → PR → plan—means even a casual, caveman-style initial prompt gets fleshed out into something precise.

## Trial and error is fundamental

Large-scale software and infrastructure work eventually requires deploy-and-test validation. You can't calculate your way to the right answer when systems are complex enough. This is how humans have always worked: a cam breaks in a factory, engineers try incremental tweaks until it works. DNA is "full of errors" because nature uses trial-and-error to decide what's ready.

Agentic AI work follows the same pattern. Incremental, interactive changes—whether through live sessions or tested PR deployments—are not a weakness. They're how complexity actually gets built. Embrace the iteration.

## See also

- [Keep repos small for agents](keep-repos-small-for-agents.md)
- [AI mistakes and safe usage](ai-mistakes-and-safe-usage.md)
- [Git worktrees for agentic development](git-worktrees-for-agentic-development.md)
- [Tmux for multi-agent development](tmux-for-multi-agent-development.md)
