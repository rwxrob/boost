# Beginner Boost

![WIP](https://img.shields.io/badge/status-wip-red.svg)

Every [May the 4th] I hold a free series of live-recorded *Beginner
Boost* sessions on [Twitch] and [YouTube] that run through the end of July
(or so). Below is the outline of current and planned content. This is a
living document; it changes from day to day --- even hour to hour ---
[like software].

[Twitch]: <https://twitch.tv/rwxrob>
[YouTube]: <https://youtube.com/rwxrob>
[May the 4th]: 20210503061604
[like software]: <https://github.com/rwxrob/zet/search?q=knowledge%20as%20source>

> 🌟 You may want to *Watch* this GitHub repo to be notified of updates and
changes. A *Star* would be nice as well. If you want to contribute open
an *Issue* to discuss it. Much of the content will be drawn from my
[personal Zettelkasten repo] and my [learning labs] (which you might
also want to watch for changes, but they are not always Boost related).
You can also help fund this project by clicking *Sponsor*. Thanks.

[personal Zettelkasten repo]: <https://github.com/rwxrob/zet>
[learning labs]: <https://github.com/rwxrob/lab>

## Day 0: Get Set Up, May the 4th Be With You

[[Full Live-Recorded Video]](https://youtu.be/UkE2KMsVzjQ)

1. [What is the purpose of Beginner Boost?](20210504161937)
1. [Beginner Boost, Day 0, Logistics](20210505025834)
1. [Who is the Beginner Boost for?](20210505031231)
1. [Setup Essential Services][services]

[services]: <https://github.com/rwxrob/zet/tree/main/20210505023938>

## Day 1: Get Motivated to Learn ¡Viva la Revolución!

[[Full Live-Recorded Video]](https://youtu.be/U7BNE-AKLeI)

1. [Problems with Traditional Education][prob]
1. [Help Others Learn, A Social Responsibility][social]
1. [Become an Autodidact][autodidact]
1. [Tribute to Aaron Swartz][aaron]

[prob]: <https://github.com/rwxrob/zet/tree/main/20210506155713>
[social]: <https://github.com/rwxrob/zet/tree/main/20210506160012>
[autodidact]: <https://github.com/rwxrob/zet/tree/main/20210506163935>
[aaron]: <https://github.com/rwxrob/zet/tree/main/20210506164209>

## Day 2: The Scientific Method, Our Greatest Learning Tool

[[Full Live-Recorded Video]](https://youtu.be/mlecHo2p7aU)

1. What did you learn about the Scientific Method?
1. The RWX Method, Refining the Scientific Method
   1. Read, Research, Reach Out, Repeat
   1. Write Who, What (If), When, Where, Why, and How
   1. Execute, Explore, Experiment, and Exercise
1. Fail Faster, Fail Better, Fail Often
1. Meritocracy in the Workplace
1. Zettelkasten (Slips in a Box) Knowledge Management System
1. Capture Thoughts as Markdown, The Syntax of Knowledge Source

## Day 3: Understand Linux Approach

[[Full Live-Recorded Video]](https://youtu.be/8myY7Ydf64g)

1. What's Linux? History of UNIX and GNU/Linux
1. Do Macs Use Linux?
1. Does Windows Use Linux?
1. The UNIX Philosophy
1. How Does Linux Work?
1. When and How Should I Install Linux?

## Day 4: Get on the Terminal

[[Full Live-Recorded Video]](https://youtu.be/GlpSi2LW4_Y)

**Mac:** Open and Configure a Mac Terminal  
**Windows:** Install and Configure Modern Windows Terminal  
**Linux:** Open and Configure a Linux Terminal  

1. What's the Difference Between TUI and CLI?
1. Terminal Emulators Replaced Teletype Machines
1. The Terminal is Different Than Command Line Shell
1. Should I install Git-Bash on Windows?
1. What about Alacritty/Kitty or ...?
1. My Favorite Theme: GruvBox
1. Can't I just use VSCode on Windows?
1. Why not just use Vim plugins with VSCode?
1. What about WSL2 on Windows?

## Day 5: Learn About Containers and Get Docker

1. What's a Container and Why Should You Care?
   1. The Difference Between "Container" and "Image"?
   1. Kubernetes is the New Operating System
   1. Containers are the New Programs

1. Install Docker
   1. Install Docker on Windows
      1. Why do I still need WSL2?
   1. Install Docker on a Mac
   1. Install Docker on Linux

## Day 6: Run Linux Inside a Container

1. Linux Distro Matters Less in Container World
   1. `docker run -it --rm ubuntu`
   1. `docker run -it --rm fedora`
   1. `docker run -it --rm archlinux`
   1. `docker run -it --rm opensuse`
   1. `docker run -it --rm nixos/nix`
   1. `docker run -it --rm kalilinux/kali-rolling`
   1. `docker run -it --rm blackarch/blackarch`

1. Run Rob's Workspace Linux Container
   1. `docker run -it --rm rwxrob/workspace`
   1. Workspace Means *Not* Optimized (3GB)

1. Use Docker Basic Commands
   1. Use Docker Pull
   1. Use Docker Run
   1. Use Docker Start
   1. Use Docker Attach and Detach `Ctrl-pq`
   1. Use Docker Stop
   1. Use Docker Remove

1. Don't Worry About Dockerfile and Compose (For Now)

## Day 7: Start Using the Linux Terminal Command Line

1. Book (free): The Linux Command Line, William Shotts
1. The Command Line, Fastest Human-Computer Interface
1. Every Command Line Entered is a Line of Shell Code
1. Use Bash, Don't Play the Shell Game (For Now)
   1. Interactive Shell is Not Same as Shell Scripting
   1. Why Bash and Not Zsh, Fish, or Whatever?
   1. Use Consistent Supported Bash Tab Completion
1. Understand the Parts of the Command Prompt
   1. Username
   1. Hostname
   1. Interesting Origin of Email Addresses
   1. Current Directory
   1. Git Branch
   1. Long Versus Short Prompts
1. Navigate Your Command History Efficiently
   1. Use Up and Down Arrow Keys (For Now)
   1. Use Initial Spaces to Stop Adding to History
   1. Use Hashtag to Make One Line Comment Notes
   1. Don't Learn More for Now (Learn `set -o vi` Later)
1. Send Special Terminal Escapes
   1. Remember, a Terminal is Just a Teletype Emulator
   1. `Ctrl-c` to Send Interrupt Signal (*Not* Copy)
   1. `Ctrl-[` to Send Exact Same as `Esc` Key
   1. `Ctrl-s`/`Ctrl-q` to Buffer and Suspend/Unsuspend
      1. "Help! My Terminal is Frozen!"
   1. `Ctrl-z` to Background Current Process (Not Quit)
   1. `Ctrl-d` to Send EOF (End of File)
   1. `Ctrl-v`, `Ctrl-x` Do Nothing (From Shell)
   1. Silencing the Terminal Bell
1. Cut and Paste from Host Operating System
   1. Depends on Terminal
   1. First Ever Mouse Had Three-Button, Middle Was Paste
   1. Windows Right Click to Paste
   1. `Ctrl-x|c|v` Don't *Usually* Work (And Shouldn't)
1. Learn Useful and Essential Beginning User Commands
   1. Use `reset` to Fix Bork Terminal
   1. Use `clear` to Clear Screen (and Alternatives)
   1. Use `watch` to Repeat Things and See Output
   1. Use `man` Command to Get Help
   1. Use `less`/`more` for Paging Output
   1. Use `ls` Command to List Files and Directories
   1. Use `type` to See What Type of Thing It Is
   1. Use `which` to See Where Command Lives
   1. Use `pwd` Command to Show Working Directory
   1. Use `cd` Command to Change Directories
      1. Use `cd foo` to Change Into `foo` Child Directory
      1. Use `cd ..` to Change Into Parent Directory
      1. Use `cd -` to Toggle Change to Last Directory
      1. Use `cd ~` (or Just `cd`) to Change to Home Directory
      1. Don't Fall for Silly `cd` Replacements (Use `CDPATH`)
   1. Use `cal` to View Calendar
   1. Use `date` to View Dates and Times
      1. Use `date -u +%Y%m%d%H%M%S` to Get ISO Second
      1. Use `date -d 'last week'` to Exact Time Last Week
   1. Use `bc` for Floating Point Precision Math Calculations
      1. Don't Forget To Set Scale (`scale=2`)
      1. Use Semicolons for Same Line
   1. Use `top` to See Running Processes
      1. Often Replaced with `htop`

## Day 8: Build Up Up Container with Software

1. Create and Name a Workspace Container
   1. `docker run -it --name boost -h boost ubuntu`
   1. Difference Between Detach and Exit
      1. Detach with `Ctrl-pq`
      1. Exit with `exit`
1. Understand Container States and Storage
1. Manage Software Packages from Command Line
   1. Restoring Documentation and More with `unminimize`
   1. Understand Linux Package Management
   1. Use `apt update`
   1. Use `apt search`
   1. Use `apt install`
   1. Use `apt remove`

## Day 9: Work with File System from Command Line

1. [OverTheWire], Fun Hacker Practice
1. Understand the Linux File System
1. Use `mkdir` to Make a New Directory
1. Use `rmdir` to Remove an Empty Directory
1. Use `touch` to Make a New File (or Update Timestamp)
1. Use `mv` to Move or Rename a File or Directory
1. Use `cp` to Copy a File or Directory
1. Use `rm` to Remove a File or Directory
1. Use `ln` to Link to a File or Directory
   1. What is the difference between hard and soft link?
   1. Used `ln` for Multicall Executables (BusyBox)
1. Use `file` to See What Details About File
1. Use `stat` to See Exhaustive Details About File
1. Use `cat` to Display Lines of a File, First to Last
1. Use `tac` to Display Lines of a File, Last to First
1. User `head` to Display Number of Lines at Top
1. User `tail` to Display Number of Lines at Bottom
1. Use `grep` to Find Lines of a File
1. Use `uniq` to Find Unique Lines
1. Use `sort` to Sort Lines
   1. Use `wc` to Count Lines, Words, Bytes, and Runes

[OverTheWire]: <https://overthewire.org/wargames/bandit/>

## Day 10: Understand Streams, Pipes, and Redirection

1. Understand Standard Output and Error
   1. Use `echo` and `printf` to Print Stuff
      1. Difference Between `echo` and `printf`
      1. Never Use `print` (for Printers)
   1. Use `>` File Redirect Operator
      1. Fix `find` Command Errors with Redirection 
   1. Use `>>` File Append Operator
   1. Use `>|` Force File Overwrite Operator

1. Understand Pipes and Pipeline
   1. Use `|` Pipe Operator to Connect Out with In
   1. Power of Pipelines and Shell Integration
      1. Watch Ken Thompson Describe UNIX Pipes
      1. Pipes are at the Core of UNIX Philosophy
      1. Use Shell and Pipes from Within Applications
   1. Transform and Filter Lines
      1. Use `nl` to Number Lines
      1. Use `rev` to Reverse Line
      1. Use `tee` to Pipe *and* Redirect Lines
      1. Use `echo` for Arguments and `cat` for Lines
      1. Use `xargs` to Transform Lines Into Arguments
      1. Use `cut` to Remove Stuff from Lines of Stream
      1. Use `tr` to Translate Stuff in Lines of Stream
      1. Use `sed` to Edit Lines of Stream (Streamed `ed`)
      1. Use `jq` to Select from JSON Input
      1. Use `yq` to Select from YAML Input

1. Get Standard Input Into Your Programs
   1. Input Most Useful After You Know How to Code
   1. Use `<` File Input Redirect Operator
      1. Always Use `<` Instead of `cat foo | ...`
   1. Use `<<` Here Document Operator
   1. Use `read` to Read Input
      1. Combine `printf` and `read` to Create Prompts
      1. Always Use `-r` and Understand Why
      1. Use `for` or `while` with Read for Line Loops
   1. Use `curl` to Read Input from Internet
      1. Usually Combined with `jq` or `yq`

## Day 11: Edit Files with Vi (Then Vim)

1. Why Vi/m and Not NeoVim/Emacs/Nano/VSCode?
1. Restore Esc Key to Its Original Keyboard Home
1. Do the Vim Tutorial (`vimtutor`), But Beware
1. Appreciate the Difference Between `vi` and `vim`
1. Start with Defaults and *Zero* Configuration
   1. Complex `.vimrc` is *Not* for Beginners
   1. Eventually, Learn a Little Vimscript
1. Extend Vim With the Shell, Not Plugins (At First)
1. Extend Vim with `Plug` for Syntax and Specifics
1. Avoid Vim Pane Splitting (Use TMUX Instead)
1. Use `set -o vi` for Efficient Shell History Handling

## Day 12: Manage Users, Groups, and Permissions 

1. Create and Manage Users and Groups
   1. Use `adduser` and `useradd` to Create User
   1. Use `usermod` to Modify User Settings
   1. Use `groupmod` to Modify User Settings
   1. Use `passwd` to Change Passwords
1. Understand UNIX File and Directory Ownership and Permissions
   1. Use `ls -l` to See Permissions
   1. Use `stat` to See Even More About File
   1. Read, Write, and Execute Permission
      1. On Files
      1. On Directories
1. Modify Ownership and Permissions
   1. Use `chmod` to Change Permissions
   1. Use `chown` to Change Owner (and Group)
   1. Use `chgrp` to Change Group
1. Know About Setuid, Setgid, and Such, But Don't Use 

## Day 13: Commands *Are* Code, Start POSIX Shell Scripting 

1. You're Already Coding, Every Command *is* Code
   1. Commands are Really Just Functions with Arguments
1. POSIX Shell is a Universal Command Interpreter
   1. Shell Started as Bourne Shell, Now Ash and Dash
   1. Korn Shell Led to Bash and Zsh Interpreters
   1. Awk, Perl, Python, Ruby, Node Also Shell Scripting
1. Create First POSIX Shell Script 
   1. Come Up with a Few Commands to Run
   1. Put Those Commands in a File
      1. Use History and `echo '...' >>` to Save Typing
   1. Use `sh` Interpreter to Run Commands in File
      1. Use `bash` to Run Same Commands
      1. Use `perl` to Attempt Run and Note Errors
   1. Use `chmod +x` to Make File Executable
   1. Add `#!/bin/sh` Shebang Line to Specify Interpreter
      1. Assumes Current Shell Interpreter When No Shebang
   1. Use `shellcheck` to Be Sure It's POSIX
1. Difference Between *Running* and *Sourcing* Scripts
   1. Use `.` or `source` to Source a Script
   1. Some Stuff Can't Be Done Any Other Way
      1. Change Current Working Directory
      1. Modify Current Environment
1. Put Executables in Your `PATH` to Run From Anywhere
   1. Use `which` to See Which Executable Wins
      1. Understand Difference Between `which` and `type`

## Day 14: Manage Jobs, Processes, and Signals

1. A Running Program is a *Process*
1. A Backgrounded Program a *Job*
1. `Ctrl-z` to Background Running Process
   1. Use `jobs` to See All Background Processes
   1. Use `fg` to Bring Background Job Forward
1. Use ` &` to Start Program in Background
1. Use `pgrep`, `ps`, `/proc` to See Processes
1. Use `kill`,`pkill` to Send a Signal
1. Don't Bother with `nohup` (Use `docker`, `screen`/`tmux`)
1. Use `crontab` to Schedule Jobs

----

*Note that stuff below this line is still being worked out based on what
we have covered so far. Stuff above is baked a little more but still
always in flux. This is all very organic (and always has been).*

## Day 15: More POSIX Shell Scripting Basics

1. Commands *are* Functions with Arguments and Return Values
   1. Understand UNIX Return Values
   1. Distinguish Return Value from Output
1. Difference Between Subroutines, Procedures, and Functions
1. Variables Save State
   1. Use `=` With No Surrounding Spaces to Assign Variable
   1. Like Putting Something Into Box or on List
   1. All Variables are Global in POSIX Shell
   1. Variables Contain Data of Different Types
      1. All Types are Ultimately the Same Binary Bits
1. Conditions and Altering Flow Based on State
   1. Use `set -e` to Exit on Any Error
1. What About Object-Oriented Programming?

## Day 15: Customize Your Interactive Shell

1. We Call Them "Dotfiles" For a Reason
1. What Are All These Files?
1. Edit `~/.profile` to Point to `~/.bashrc`
1. Customize Bash (`.bashrc`,`.profile`)
   1. Customize Your Bash Command Prompt
   1. Colorize Your Pager and `man` Pages
   1. Prefer Commands Over Aliases (Cuz Subprocs)
1. Use `alias` to Add Interactive Shell Aliases
   1. Avoid Abusing Aliases, Commands are Better
1. Exporting Functions
1. Use `eval` to Source Other Code Inline
   1. The Case For and Against
1. Use Backslash to Disable Aliases and Functions

## Day 16: Basic Networking and Remote Access

1. How Does the Internet Work?
   1. Short History of the Internet
   1. Nerds 2.0.1, Glory of the Geeks
   1. Use `ip`/`ifconfig` to See IP Information
      1. BTW, It's `ipconfig` on Windows
   1. Use `dig`/`nslookup` to Lookup Domain Names
   1. Use `ping` to Explore Hosts (But Don't Depend On)
   1. Use `netstat -tupl` to See Local Network Stuff
   1. Use `nmap` to Scan for Hosts and Ports
1. Use Secure Shell for Remote Access and File Transfer
   1. Setting Up Secure Shell on Local System
      1. Use `ssh-keygen -t ed25519`
   1. Use `ssh` to Make Secure Remote Connections
   1. Use `scp` to Copy Remote Files Securely
   1. Use `sftp` When You *Must* Use FTP
1. Use `rsync` for Customized Backups

## Day 17: Get On Git and GitHub

1. Minimal Secure Shell Setup
1. See What's Listening with `netstat -tulpn` and `nmap`
1. Setup Git and GitHub CLI
1. Preserve Your Dotfiles on GitHub
1. Create a Dotfiles Docker Image


## Day 15: Manage Screen and State with Screen and TMUX

1. Any Window Manager Will Do
1. Screen, the One that Started It All
1. TMUX, a Modern Multiplexer
1. Cut and Paste with TMUX Buffers
1. Customize and Extend TMUX (`.tmuxrc`)

## Day 11: Use Terminal for Fastest Possible Web Research

1. Leverage the Power of `find`
1. Why Browse the Web from the Terminal?
1. Pick a Terminal Web Browser for You
1. Use and Configure `w3m` Text Web Browser
1. Use and Configure `lynx` Text Web Browser
1. Turn Shell History Into Search History
1. Use `curl` to Fetch Files and More


1. Write Your Own Embedded Bash Tab Completion

## Day 15: Create and Manage Knowledge Content

1. Structure Data with YAML and JSON
1. Query YAML and JSON with `yq` and `jq`
1. Create and Publish Web Documents
   1. Write Basic Pandoc Markdown
   1. Code a Static Site Generator in 10 Lines of Bash
   1. Host Web Content on GitHub and Netlify for Free
   1. Code Basic HTML
   1. Code Basic CSS
   1. Code Basic (DOM) JavaScript
1. KEG, The Knowledge Exchange Grid
   1. What is KEG and Why Should I Care?
   1. Capture Knowledge with `kn`
   1. Share Knowledge with `keg`
1. PEGN, Parsing Expression Grammar Notation
   1. Everything's a Grammar
   1. PEGN AST, Consistent Abstract Syntax Tree Notation
   1. Create and Parse Your Own Language
   1. Better Than Regular Expressions

## Code C for Understanding

1. Head First C

## Code Python for Applications and Automation

1. ??

## Code Go for Utilities and Services

1. Code Go Templates

## Get Good to Gig

1. The Prescient Technology Professional
1. Use Twitter Productively
1. Follow Usenet News Groups
1. Use IRC and WeeChat from Terminal

## Create and Manage Docker Containers

1. Code Dockerfiles

## Cloud-Native with Kubernetes

1. Explore Kubernetes Locally with Kind
1. Develop Cloud-Native Applications
1. Deploy Kubernetes Applications with Helm
1. Setup and Administer Kubernetes
