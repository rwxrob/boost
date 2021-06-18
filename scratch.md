*This file contains stuff that needs to make it into the README.md file
eventually.*

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
