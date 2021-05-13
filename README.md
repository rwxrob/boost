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

1. Run Your First Linux Container
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
   1. Workspace Means *Not* Optimized
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
1. History and Value of Command Line Interfaces
1. Every Command Line Entered is a Line of Shell Code
1. Use Bash, Don't Play the Shell Game (For Now)
   1. Interactive Shell is Not Same as Shell Scripting
   1. Why Bash and Not Zsh, Fish, or Whatever?
   1. Use Consistent Supported Bash Tab Completion
1. Get Your Bearings and Recognize Value of CLI
   1. Use `man` Command to Get Help
   1. Use `ls` Command to List Files and Directories
   1. Use `cd` Command to Change Directories
   1. Use `pwd` Command to Show Working Directory
   1. Use `clear` to Clear Screen (and Alternatives)
   1. Use `cal` to View Calendar
   1. Use `date` to View Dates and Times
   1. Use `bc` to Do Math
   1. Use `ps` and/or `top` to See Running Processes
1. Create, Move, Rename, Remove, and Link to Stuff
   1. Understand the Linux File System
   1. "I do not think that word means what you think it means."
   1. Use `mkdir` to Make a New Directory
   1. Use `rmdir` to Remove an Empty Directory
   1. Use `touch` to Make a New File (or Update Timestamp)
   1. Use `mv` to Move or Rename a File or Directory
   1. Use `cp` to Copy a File or Directory
   1. Use `rm` to Remove a File or Directory
   1. Use `ln` to Link to a File or Directory
      1. What is the difference between hard and soft link?
      1. Used `ln` for Multicall Executables (BusyBox)
1. View and Edit Files Without a Text Editor
   1. Understand *Basic* UNIX Pipes and Redirection
      1. Use `|` Pipe Operator
      1. Use `>` Write File Redirect Operator
      1. Use `>>` Append to File Redirect Operator
      1. Use `<` File Input Redirect Operator
      1. Use `<<` "Here" Document Operator
      1. But Wait, There's More
      1. Understand Standard Input, Output, and Error
   1. Use `echo` to Print Stuff
   1. Use `cat` to View or Concatenate Stuff
      1. Never `cat` to a Pipe!
1. Find Stuff In and Out of Files
   1. Use `grep` to Find Lines of a File
   1. Use `find` to Find Anything
1. There's So Much More
   1. Use `ssh` to Make Secure Remote Connections
   1. Use `scp` to Copy Remote Files Securely
   1. Use `sftp` When You *Must* Use FTP

## Day 8: Basic Linux System Administration

1. Create and Manage Users and Groups
   1. Use `adduser` and `useradd` to Create User
1. Understand UNIX File and Directory Ownership and Permissions
   1. Use `chmod` to Change Permissions
   1. Use `chown` to Change Owner (and Group)
   1. Use `chgrp` to Change Group
1. Manage Software Packages from Command Line
   1. Which Package Manager Should I Use?
   1. Use `apt update`
   1. Use `apt search`
   1. Use `apt install`
   1. Use `apt remove`

## Day 9: Edit Files *Simply* with Raw Vi and Vim

1. Why Vim and Not NeoVim/Emacs/Nano/VSCode?
1. Restore Esc Key to Its Original Keyboard Home
1. Do the Vim Tutorial (`vimtutor`)
1. Appreciate the Difference Between Vi and Vim
1. Start with a Solid and Safe Vim Configuration (`.vimrc`)
   1. Complex `.vimrc` Coverage is *Not* for Beginners
1. Extend Vim as God Intended (With the Shell)
1. Extend Vim with Plug for Syntax and Specifics
1. Don't Use Vim Pane Splitting (Use TMUX Instead)
1. Set Vi Mode in Shell (`set -o vi`)

## Day 10: Manage Screen and State with Screen and TMUX

1. Any Window Manager Will Do
1. Screen, the One that Started It All
1. TMUX, a Modern Multiplexer
1. Cut and Paste with TMUX Buffers
1. Customize and Extend TMUX (`.tmuxrc`)

## Day 11: Use Terminal for Fastest Possible Web Research

1. Why Browse the Web from the Terminal?
1. Pick a Terminal Web Browser for You
1. Use and Configure `w3m` Text Web Browser
1. Use and Configure `lynx` Text Web Browser
1. Turn Shell History Into Search History
1. Use `curl` to Fetch Files and More

## Day 12: Your Commands *Are* Code, *Basic* Shell Scripting (Part I)

1. Learn Basic POSIX Shell Scripting (It's Not Hard)

## Day 13: *Basic* Shell Scripting (Part II)

## Day 14: *Basic* Shell Scripting (Part III)

1. Write Your Own Embedded Bash Tab Completion

## Day 15: Customize Your Interactive Shell

1. Customize Bash (`.bashrc`,`.profile`)
   1. Customize Your Bash Command Prompt
   1. Colorize Your Pager and `man` Pages
   1. Prefer Commands Over Aliases (Cuz Subprocs)

## Day 14: Get On Git and GitHub

1. Minimal Secure Shell Setup
1. See What's Listening with `netstat -tulpn` and `nmap`
1. Setup Git and GitHub CLI
1. Preserve Your Dotfiles on GitHub
1. Create a Dotfiles Docker Image

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
