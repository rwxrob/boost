_Best read using an editor that allows collapsing the outline such as Obsidian, Vim, or Emacs._
## Get ready

Don't even bother with the Boost until you have achieved the following:

1. Get a basic computer (preferably laptop) ($250)
2. Get GitHub Copilot subscription ($10/month)
3. Plan an hour per day of AI-assisted learning
## What to do

Knowing what to learn is the hardest part of achieving any goal, tech job or otherwise. All the motivation and intelligence in the world won't help until you know this. So, this outline is primarily designed to outline exactly that, not so much teach it to you. You can learn it yourself once you know what to focus on.

>![NOTE]
Notice that almost everything in this outline starts with a verb, a call to action, just as would be listed on a job description or resume.

1. Join the Beginner Boost community
	1. Hello friends
	2. What's a Beginner Boost?
		1. Why was the Beginner Boost created?
		2. Who is it for and who should skip?
		3. How long will the Boost take?
		4. What will I get from the Boost?
		5. Why is there a Boost every year?
		6. What about previous Boosts?
		7. [Follow Rob for updates](https://link.tr/rwxrob)
	3. What do I need before starting the Boost?
		1. Mac or Windows Computer
			1. Why not a Linux Desktop?
		2. GitHub account with Copilot subscription
			1. Why GitHub?
			2. Pick a good username for work
			3. Manage public perception
		3. Time, an hour per day minimum
		4. Motivation and commitment
	4. Be supportive
		1. Follow The Four Agreements in life as best you can
			1. Be impeccable with your word
			2. Don't take anything personally
			3. Don't make assumptions
			4. Do your best
		2. Be less offensive, soften the message when you can
		3. Be less offended, give people benefit of the doubt
		4. Don't be afraid to ask anything
		5. Don't respond with conclusions, explain them
	5. Contribute if you like
		1. How is this free?
		2. Sponsor financially
		3. Make clips
		4. Create chapter comments on YouTube
		5. Post job fairs and opportunities
		6. Help others out
		7. Pay it forward
		8. [Say something nice](https://github.com/rwxrob/nice)
2. Start managing your knowledge
	1. Create your first GitHub repo: `notes-boost`
	2. Follow the `rwxrob/boost` repo
		1. Star the boost repo
		2. Follow the boost repo
		3. Submit issue when fix needed
		4. Clone if you already know how
		5. Join the boost discussion group
			1. Introduce yourself
			2. Follow each others socials
			3. Ask questions when you need to
	3. Take notes using just GitHub
	4. Install free version of [Obsidian](https://obsidian.md/) (optional)
3. Connect and practice working with others
	1. Manage your fear of people
	2. Aggressively seek out friendships and mentoring 
	3. Build your professional learning community
4. Learn how to learn
	1. Write in journal every day
		1. Proven to reduce stress, increase dopamine
		2. Promotes personal accountability
		3. Provides a place to vent that doesn't harm others
		4. Leads to further creativity and ideas 
	2. Know thyself
		1. What specific activities fill you with joy?
		2. What triggers you?
		3. What do people say when they compliment you?
		4. What comes naturally to you?
	3. Know what to know
	4. Create and control your own plan and progress
	5. Use real tools of the trade
	6. Gain experience building stuff you want
	7. Learn from those who have experience
	8. Focus on creativity, ideas, and solutions
5. Setup your technical learning lab
	1. Install free version of [VS Code](https://code.visualstudio.com/)
		1. Pick your favorite theme
		2. Add your first plugin extensions
			1. Favorite theme
			2. GitHub Copilot
		3. Learn to use the VS Code terminal
			1. Why not use another terminal instead?
		4. Integrate with GitHub
	2. Windows: Install WSLv2
	3. Mac: Install `brew` software package manager
	4. Install `git` Source Management Tool
	5. Install `gh` GitHub Command Line Tool
	6. Install `podman` Container Engine
	7. [Build and run `boost` container locally](https://github.com/rwxrob/boost)
6. Learn to move at _terminal velocity_
	1. Understand what a terminal is and why it matters
		1. What is a "user interface"?
			1. Terminal User Interface (TUI)
			2. Command Line Interface (CLI)
			3. Graphical User Interface (GUI)
			4. Conversational User Interface (CUI)
		2. Decipher the symbols in the CLI prompt
		3. Use terminal with modern agentic AI
	2. Look up help for commands
		1. Just use your AI
		2. `man`, `help`, `info`, `-h`
	3. Work with the UNIX/Linux filesystem from terminal
		1. What is a _file_? (inode, file descriptor)
		2. What is a _directory_?
		3. "Everything is a file"
		4. Lookup current working directory
		5. List everything in a directory
			1. Understand hidden files
			2. Current directory: .
			3. Parent directory: ..
		6. Create a file
		7. Work with permissions
		8. Renaming files/directories
		9. Move files and/or directories
		10. Link one file/directory to another
		11. Rename a file/directory
		12. Remove a file/directory
		13. `pwd`, `ls`, `cd`, `.`, `..`, `/`, `~`, `-`, `touch`, `mv`, `cp`, `chmod`, `chown`, `ln`, `rm`, `mkdir`, `rmdir`
	4. Work with UNIX/linux processes from terminal
		1. What is a _process_?
		2. Understand input and output (stdin, stdout, stderr)
			1. Scroll through output with a pager
			2. Simple stdin/out/err redirection
			3. `less`, `more`
	5. Work with text in UNIX/Linux terminal
		1. Work with text streams using pipes and filters
			1. Understand the _Unix Way_
				1. Build **small tools**
				2. Each tool does **one job really well**
				3. Then you **chain them together**
			2. What is map/filter/reduce?
		2. Edit text file interactively from the terminal
			1. Pick best method to edit text files for the job
				1. Are you just adding output to a file?
				2. Are you just filtering the output of a command?
				3. Is an interactive terminal editor available?
			2. Understand interactive terminal editor options
				1. ed/ex - the official editor of UNIX
				2. vi - visual mode of ed/ex
				3. vim - improved vi
				4. nvim - fatter version of vim in Lua
				5. emacs - not vim
			3. Learn to edit text files from terminal with nano
			4. Learn do edit text files from terminal with vim
				1. Understand [original keyboard layout](img/unix-keyboard.jpg)
		3. Work with structured data formats
			1. Delimited, CSV, JSON, YAML, TOML (INI)
	6. Work with terminal command shells
		1. What is a _shell_?
		2. Pick the best command shell for the job
			1. `bash`,`zsh`,`pwsh`, etc.
		3. Work with bash shell
		4. Write bash shell scripts
			1. What is a _script_?
			2. `shellcheck`,`shfmt`
	7. Use tmux/screen for terminal window management and safety
		1. screen compatibility and why C-a is the right prefix
		2. C-a w is the only window manager you need
		3. Why tmux beats GUI multi-agent tools like Conductor and T3 Chat
		4. [Scripting and automating agents with tmux send-keys](tmux-for-multi-agent-development.md)
		5. [Git worktrees for parallel agentic development](git-worktrees-for-agentic-development.md)
		6. [Gas Town: multi-agent workspace manager worth knowing](gas-town-multi-agent-orchestration.md)
7. Create your own `dot` GitHub repo with common configs and scripts
8. Create your first web site
	1. Create your special profile repo: `username`
	2. Learn GitHub Flavored Markdown
	3. Create your personal website repo: `username.github.io`
9. [Learn to read and sometimes write code](read-code-and-terminal.md)
10. [Harness the power of AI](harness-ai-power.md)
	1. Don't fear AI, use it
	2. Your ideas don't need permission — build them yourself
	3. Become impossible to ignore
	4. Make yourself valuable to companies
	5. [Build MCP servers to connect AI agents to real systems](build-mcp-servers.md)
	6. (Optional) [Add live voice input to GitHub Copilot with Superwhisper](voice-input-with-superwhisper.md)
11. [Become a startup](become-a-startup.md)
	1. Operate like a founder, hired or not
	2. Actually starting a startup
	3. Acting like a startup inside a company (intrapreneurship)
	4. Why this is the way to get a job in tech today
12. [Use rwxrob/agentic-ai-template for new projects](use-agentic-ai-template.md)
	1. What it includes and why it is minimal
	2. How to create a new repo from the template
	3. Keep AGENTS.md up to date so every session picks up where the last left off
	4. [Your agentic AI template is the new dotfiles](agentic-template-is-the-new-dotfiles.md)
