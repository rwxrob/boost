# Boost tools then and now

The Beginner Boost has been running since 2020, and SKILSTAK has been teaching tech skills since 2013. Over the years, the tools and skills we taught evolved, but the arrival of agentic AI has forced the sharpest reassessment yet. Some tools that were once central have become optional. Others that seemed like power-user luxuries are now essential. And some entirely new categories have appeared that did not exist in any previous Boost.

Here is an honest evaluation of every major tool and skill the Boost and SKILSTAK have traditionally covered, how important each one is today, and what you should prioritize instead when something has shifted.

## Tech is now like medicine

Before anything else, understand what has changed about the career itself. AI has eliminated most junior and mid-level tech jobs. The entry-level positions where you could learn on the job — writing CRUD endpoints, building simple UIs, maintaining straightforward infrastructure — are being handled by AI agents now. What remains are senior roles that require deep understanding across multiple technologies and the ability to architect systems, not just implement them.

This means approaching a career in tech is now more like a decision to become a doctor. You are committing to mastering multiple disciplines at a meaningful depth — not because it is fun to collect skills, but because the work that still needs a human requires someone who can guide AI across complex systems with confidence. An AI agent can generate a Kubernetes deployment manifest, but it takes a person who understands networking, storage, security, and application architecture to know whether that manifest will actually work in production. An AI agent can write a database migration, but it takes someone who understands data modeling, consistency guarantees, and the specific business domain to know whether that migration is safe.

The good news is that AI itself is the biggest boost to your learning. AI *is* the Boost. You are not studying from textbooks alone anymore. You can use AI to teach yourself interactively, build real things as you learn, and reach working competence in a new technology faster than any previous generation of engineers. The depth that used to take years of on-the-job experience can now be reached in months of focused, AI-assisted study and building. But you have to actually do the learning. There are no shortcuts to understanding — only faster paths to it.

One critical tip: tell the AI that you are learning. Put it in your AGENTS.md file for every project. Something like "The person working on this project is actively learning these technologies. Explain decisions thoroughly, call out potential pitfalls, and prefer safe, well-understood approaches over clever shortcuts." When the AI knows you are a learner, it explains *why* it made a choice instead of just making it. It warns you about things a senior engineer would catch on their own. It teaches you as it builds. Without that context, the AI optimizes for speed and assumes you already understand the tradeoffs. The difference in what you learn from each session is enormous.

## The most important skill now

The single most important skill you can develop right now is learning how to use AI to teach yourself what you want to learn, and then using AI to deliver the things you have learned to build.

This is the meta-skill that makes everything else on this page accessible. Want to learn Terraform? Use AI to teach it to you interactively, in context, based on the specific thing you are trying to build. Want to set up a Proxmox server? Have AI walk you through it step by step while you do it. Want to understand how containers work? Have an AI agent explain it while you follow along in a terminal.

One of the most important skills to develop is learning to trust AI exactly as much as you would trust a stranger who posts an answer on Stack Exchange — which is to say, not at all until you verify it. Interact with AI the same way you would interact with a human you have never worked with before: assume competence, but confirm everything. Read its output. Question its choices. Test its code. The people who get burned by AI are the ones who either refuse to use it or trust it blindly. The people who get the most out of it treat it like a confident colleague who is sometimes wrong — useful, fast, and worth double-checking every time.

The Boost has always been about knowing what to learn. AI has now made the learning itself dramatically faster — but only if you know how to use it. That means learning to ask good questions, iterate on prompts, evaluate AI output critically, and direct agents toward what you actually need. It also means demonstrating what you know by shipping real things. Employers do not care how you learned something. They care that you can use AI to deliver working software, functioning infrastructure, and real solutions to real problems.

Every tool and skill below is something you can learn with AI assistance. The question is which ones are worth your time, and how deep you need to go.

## More important than ever

### Git and GitHub (`git`, `gh`)

Git has been in every Boost since the beginning. It was always important. Now it is the foundation that most agentic AI workflows depend on — agents commit, branch, diff, and revert constantly. GitHub is where your proof-of-work lives, where employers look, and where AI agents operate natively through `gh` and GitHub Actions. Git worktrees, which were barely mentioned in previous Boosts, are now important for running multiple agents in parallel on the same repo.

But git is no longer just a code tool. Knowledge itself is now managed as source code. Your notes, your documentation, your learning journal, your personal knowledge base — all of it belongs in git repos with full revision history. Every change is tracked, every version is recoverable, and collaboration happens through pull requests and issues. This is how professionals manage knowledge now, and it is how you should manage yours from day one. A GitHub repo full of well-organized Markdown files beats any proprietary note-taking app because you own it, you can search it, you can version it, and no company can take it away from you or change the format out from under you.

A note on GitHub specifically: the Boost has always recommended GitHub over alternatives, and the AI era has made that recommendation stronger. GitHub Copilot, the `gh` CLI, GitHub Actions, and the agentic AI ecosystem are built on GitHub. GitLab has no equivalent command-line tool, no native AI agent integration, and no clear path to the agentic workflows that define modern development. If you are starting fresh, GitHub is the right choice.

**Then:** Important tool you had to learn for code.
**Now:** Foundation for code, knowledge management, documentation, and agentic workflows. And it should be GitHub.

### Markdown

Every Boost taught Markdown as part of web development and note-taking. It was always useful. Now it is something more fundamental: Markdown is the language of knowledge.

Every README, every AGENTS.md, every Obsidian vault, every GitHub issue, every AI prompt, every piece of documentation that matters is written in Markdown. It is plain text, which means it works with git, it works with grep, it works with AI, and it will still be readable in fifty years. No proprietary format can make that promise.

Learn GitHub Flavored Markdown specifically. Do not waste time learning Microsoft Word, Google Docs, or any other proprietary document format for your own knowledge work. Those tools lock your thinking inside someone else's software. Markdown files are yours. They live in git repos with full revision history. They render beautifully in Obsidian, on GitHub, in any static site generator, and in every terminal. They are the universal format for both human and AI communication.

If you take one thing from the Boost, make it this: write everything in Markdown, store it in git, and never depend on a proprietary tool for your own knowledge again.

**Then:** Basic web skill and note-taking format.
**Now:** The standard format for AI-assisted development, knowledge management, and professional communication.

### Terminal and CLI fluency

Every Boost spent significant time on terminal basics — filesystem navigation, processes, stdin/stdout/stderr, permissions, the whole `pwd`/`ls`/`cd`/`chmod` vocabulary. That investment pays off more now than ever. Agentic AI lives in the terminal. GitHub Copilot CLI, Claude Code, and every serious AI coding tool runs as a terminal application. You need to read output, understand exit codes, spot failures buried in log output, and not panic when something goes wrong silently.

One skill that has become especially important: being able to read and understand single-line shell commands quickly. When an AI agent proposes a command before executing it, you get one chance to review it. That command might chain together several operations with pipes and semicolons. It might use flags you have not seen before. It might redirect output somewhere unexpected. You need to be comfortable enough with shell syntax to scan that line and decide whether it is safe to run. This is not about memorizing every flag for every command — it is about being fluent enough to spot when something looks wrong or dangerous.

**Then:** Core curriculum, the "terminal velocity" section.
**Now:** Even more central. The terminal is the primary interface for AI-assisted work, and reading shell commands quickly is a daily safety skill.

### Secrets hygiene in agentic workflows

Previous Boosts covered environment variables as a way to pass configuration and secrets into programs. That advice needs updating. When an AI agent has access to your terminal, it can read environment variables. Every secret stored as an environment variable is visible to the agent session. This changes the threat model significantly.

Prefer writing secrets to files with restrictive permissions (`chmod 600`) and reading them from disk rather than exporting them into the shell environment. Use credential helpers and secret managers that do not expose values in the process environment. Keep API keys, tokens, and passwords out of your shell history and out of any context an agent can access. This is not paranoia — it is basic hygiene for a workflow where an automated system is reading and executing commands in your terminal.

**Then:** Environment variables were the standard way to handle secrets.
**Now:** Keep secrets out of the environment entirely. Files with restrictive permissions, credential helpers, and secret managers are the safer path when agents share your terminal.

### tmux

Previous Boosts taught tmux as a terminal multiplexer — handy for keeping sessions alive and splitting panes. The SKILSTAK linux.md even showed `synchronize-panes` for configuring multiple machines at once. That was useful then and the original use case is still just as relevant today: keeping a running process alive on a remote system even after you disconnect. SSH into a server, start a long-running job inside a tmux or screen session, and walk away. If your connection drops, the process keeps running. Reattach later and pick up where you left off. This is still one of the most common and practical reasons to use tmux or screen, and it is a skill every person working with remote systems needs.

What tmux has also become is something on top of that: a multi-agent orchestration layer. Running multiple AI agents in parallel tmux panes, scripting them with `send-keys`, managing worktrees across sessions — this is a practical multi-agent workflow that works better than most GUI tools trying to solve the same problem.

**Then:** Nice-to-have power-user tool for session persistence and pane splitting.
**Now:** Essential for both its original purpose and as an effective multi-agent orchestration tool.

### Reading code

The Boost always taught reading code, but it was bundled with learning to write code. Those two skills have diverged. AI writes most of the code now. Your job is to direct it, verify what it produces, and catch mistakes before they ship. The ability to scan a diff, trace logic, and spot when something is off is a core part of working with AI agents effectively.

This does not replace creative thinking or systems reasoning — those are still the skills that make someone genuinely valuable. But reading code is the mechanical prerequisite that lets you apply that thinking to what AI produces.

**Then:** Part of learning to code.
**Now:** A critical verification skill that sits alongside creative problem-solving.

### Containers and virtual machines

SKILSTAK taught LXC containers early on, and the Boost moved to Podman. Containers are how modern software ships — every CI/CD pipeline, every cloud deployment, every reproducible dev environment runs in containers.

Agentic AI has added a new dimension to this. When you give an AI agent access to your terminal, it can run arbitrary commands on your system. Containers and virtual machines provide the isolation layer that makes this safe. Running agents inside a container or VM means they can install packages, modify files, execute code, and make mistakes without any risk to your host system. This is not theoretical — it is how responsible agentic workflows operate in practice. Understanding containers and VMs is no longer just about deployment. It is about safely running powerful automated AI workflows that would be reckless to run directly on your machine.

**Then:** Important for ops-track learners and deployment.
**Now:** Essential for everyone — both for shipping software and for safely running agentic AI workflows without risking your host system.

### Home lab with Proxmox and Terraform

SKILSTAK's start.md had detailed advice on buying laptops, mini servers, and tower servers for a home lab. That advice is worth revisiting — not because home labs are less relevant, but because what you do with one has changed.

A Proxmox server with a library of virtual machines is now a powerful agentic AI lab. Combined with Terraform, you can use AI to automate the creation and deployment of entire environments — spin up a fresh VM, provision it, run an agentic workflow inside it, tear it down, and start over. This gives you a safe, repeatable sandbox for testing AI-driven automation without risking anything in production or on your personal machine.

Even more valuable: a centralized jump host on your Proxmox server gives you a single point from which to manage and test all your agentic interactions. SSH into the jump host, run your agents in tmux sessions against isolated VMs, and keep everything off your laptop. This is the home lab use case that matters now — not just learning Linux administration, but building infrastructure that lets you experiment with AI safely and at scale.

This is also a perfect example of using AI to learn. You do not need to already know Terraform or Proxmox to get started. Use AI to teach yourself as you build it. The infrastructure knowledge you gain along the way is directly employable, and the lab itself becomes your platform for learning everything else.

For beginners who are not ready for that investment yet, a laptop with a container engine is enough to start. But if you want to go deeper, a Proxmox server is one of the best investments you can make.

**Then:** Recommended early purchase with specific hardware advice for learning Linux.
**Now:** A powerful agentic AI lab when combined with Terraform and used as an isolated testing environment for AI-driven automation.

## Still worth learning but the depth has changed

### Programming languages (Python, Go, JavaScript)

SKILSTAK taught Python, JavaScript, and Go simultaneously. The Boost covered coding fundamentals. Writing code from scratch is still a useful skill, but AI has reduced the advantage of being a fast coder. What matters more now is understanding code well enough to direct AI effectively and verify what it generates.

Focus your reading fluency on the languages that AI models know best. Language models were trained on massive codebases of Python, Go, JavaScript, and TypeScript. These are the languages where AI output is most reliable and most abundant. If you can read these languages confidently, you can review the vast majority of what AI agents produce.

This also means choosing the right language for a task matters differently now. Consider writing a multiprocess daemon: in bash, that means fragile background jobs, trap handlers, PID files, and signal gymnastics that are hard to write and harder to review. In Python or Go, the same daemon is a straightforward program with clear concurrency primitives that both you and your AI agent can reason about. The language that is easiest for AI to generate correctly and easiest for you to scan and verify is the right choice.

Languages with simple, consistent syntax have a real advantage in agentic development. Go's minimal syntax — no exceptions, no inheritance, no operator overloading — means there is less to misread during a fast code review. Python's readability has always been a strength, and it pays off even more when you are scanning AI-generated output rather than writing it yourself. Complex languages with heavy syntax and many implicit behaviors produce code that takes longer to verify, which slows down the agentic loop.

**Then:** Learning three languages was a major curriculum investment.
**Now:** Read fluency across languages matters more than write fluency in any single one. Prioritize the languages AI knows best and that are easiest to scan.

### Bash scripting

Every Boost covered bash — from basic commands through shell scripting with `shellcheck` and `shfmt`. Knowing enough bash to write a quick script, chain commands, and understand what AI-generated shell scripts do is still useful. But deep bash mastery — complex traps, parameter expansion tricks, associative arrays — matters less when AI can generate and maintain those scripts for you. And for anything involving concurrency, networking, or system daemons, you are better off in Python or Go where the code is easier for both you and the AI to get right.

**Then:** Deep curriculum item with its own section.
**Now:** Learn the fundamentals. Let AI write the complex stuff. Use a real language for anything beyond simple scripts.

### Vim

The Boost taught vim in the context of terminal text editing — `ed`, `vi`, `vim`, `nvim`, `nano`. Knowing how to open a file, make a quick edit, save, and quit is still a survival skill for machines where vim is the only editor. But investing serious time in vim motions, custom NeoVim configs, and Lua plugin ecosystems is no longer a good use of learning hours.

**Then:** Significant curriculum time on editor options and vim basics.
**Now:** Learn `i`, `Esc`, `:wq`, `/search`, `dd`, `p`. That is enough unless you enjoy it as a hobby.

### Pipes, filters, and the Unix Way

The Boost always taught the Unix philosophy — small tools, each doing one job, chained together. Understanding stdin/stdout/stderr, basic pipe chains with `grep`, `sort`, `uniq`, `wc` is still worth your time. But memorizing dozens of text processing commands matters less when you can ask AI to construct the pipeline.

**Then:** Core philosophy with deep command coverage.
**Now:** Understand the concept and the common commands. Let AI handle the exotic `awk` and `sed` one-liners.

### Data formats (JSON, YAML, TOML)

The SKILSTAK web module and Boost both covered structured data formats. You will still encounter these constantly — config files, API responses, automation definitions. You do not need to be an expert, but you need to recognize the syntax and spot errors when AI produces broken YAML or malformed JSON.

**Then:** Part of web development and coding sections.
**Now:** Read fluency matters. Deep knowledge of the specs does not.

### HTML and CSS

The SKILSTAK web module was a significant part of previous Boosts — learning HTML elements, CSS styling, creating a basic web site. Today, AI generates HTML and CSS faster than you ever will. The skill that matters is understanding what it produced, not writing it from scratch. Basic web literacy is still useful but it has moved from a core skill to background knowledge.

**Then:** Major curriculum section with its own module.
**Now:** Useful context, but not where your learning time should go.

### Linux system administration

SKILSTAK's linux.md covered buying hardware, choosing distros, setting up VMs, and home lab infrastructure. The Boost taught Linux as a user-first skill. System administration still matters for certain career paths (SRE, platform engineering) and becomes directly relevant when you are managing a Proxmox server or provisioning VMs with Terraform. Use AI to learn it as you need it — the depth you need depends on your career track, and AI makes it possible to learn operations infrastructure on demand rather than front-loading months of study.

**Then:** Central to the Boost identity. "Linux focused" was a defining principle.
**Now:** Still valuable, especially when combined with infrastructure automation. Learn it with AI as you build real things.

## Less relevant than before

### Typing speed

The SKILSTAK basic.md included "Learn how to *really* type" as a fundamental skill. Previous Boosts treated keyboard proficiency as a prerequisite. AI writes most of the text now, autocomplete fills in the rest, and voice input handles long prompts. Getting to 40-50 WPM with touch typing is enough. Chasing speed beyond that is not a good investment.

**Then:** Prerequisite skill, drilled early.
**Now:** (Optional) [Typing speed matters much less than you think](typing-matters-less-now.md).

### Deep editor customization

Previous Boosts spent time on choosing editors, understanding the vim family tree (`ed` → `vi` → `vim` → `nvim`), and setting up your environment. The editor configuration rabbit hole — NeoVim Lua plugins, custom keybindings, dotfiles repos — was always a risk, but now the opportunity cost is clear. Time spent on editor config is time not spent learning to direct AI agents or ship something.

**Then:** Part of the learning journey.
**Now:** Pick one editor, learn it enough to not fight it, move on.

### GitLab

Some previous community members used GitLab, and it came up occasionally in discussions. GitLab has no equivalent to the `gh` CLI, no native AI agent integration, and no clear path to the agentic workflows that define modern development. The AI tooling ecosystem — GitHub Copilot, Copilot CLI, Claude Code's GitHub integration — is built around GitHub. If you are still on GitLab, it is worth considering a move. If you are starting fresh, go with GitHub.

**Then:** Viable alternative to GitHub for some workflows.
**Now:** Missing the integrations that matter most for agentic AI development.

## New skills that did not exist in previous Boosts

### Agentic AI tools (Copilot CLI, Claude Code)

No previous Boost covered this because it did not exist. Learning to direct an AI agent effectively — giving clear instructions, reviewing output, iterating on prompts, knowing when to accept and when to reject — is one of the highest-leverage skills you can develop right now. But just as important: learning to use these tools to teach yourself new skills on demand. AI is both the tool you ship with and the teacher that gets you there.

### Voice input with [Superwhisper](https://superwhisper.com)

Speaking a prompt is faster than typing it. Over a full day of agentic work, voice input saves meaningful time. This is not a gimmick — it removes a real bottleneck between your thinking and the machine. Grab the [lifetime license](https://superwhisper.com) while it is still available.

### AGENTS.md and agentic project templates

Structuring a project so AI agents can pick up where the last session left off — with clear AGENTS.md files, small repos, and good context — is a new kind of engineering literacy that no previous Boost could have anticipated.

### The SkilStak method, revisited

SKILSTAK was originally built around the idea of managing your personal skill stack — tracking what you know, scoring your mastery, identifying gaps, and prioritizing what to learn next. That concept is more relevant than ever, but the skills on the stack have changed. The evaluation framework still works. The items on the list need updating. That is what this article is for.

## The bottom line

The Boost has always been about learning to think creatively, solve real problems, and build things that matter. The most important shift is this: AI is now both the tool you deliver with and the way you learn what you need to deliver. Use it to teach yourself. Use it to build. Use it to demonstrate what you can do.

The tools in service of that have shifted. Git, terminal, containers, tmux, and Markdown are as important as ever. AI fluency — directing agents, reading their output, using voice input — is new and essential.

Write everything in Markdown. Store everything in git. Learn to read the languages AI knows best. Keep your secrets out of the environment. Develop your ability to think across systems and spot opportunities for efficiency. Those habits will serve you better than any amount of editor configuration or typing speed drills.
