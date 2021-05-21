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

[[Full Live-Recorded Video]](https://youtu.be/YcBIUOSOAeg)

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

[[Full Live-Recorded Video]](https://youtu.be/_jJWOgZwBBs)


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

[[Full Live-Recorded Video]](https://youtu.be/EBtWx5m7pds)

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
   1. Use `man`/`help`/`info` Commands to Get Help
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

## Day 8: Build Up Container with Software

[[Full Live-Recorded Video]](https://youtu.be/6EDT-Vc4PCY)

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

[[Full Live-Recorded Video]](https://youtu.be/eZV14xpFAlA)

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

[[Full Live-Recorded Video]](https://youtu.be/rhknj8h0zU8)

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

[[Full Live-Recorded Video]](https://youtu.be/RJ3EVB5-Emw)

1. Vi ("Visual Mode") History and Legacy
1. Why Vi/m and Not NeoVim/Emacs/Nano/VSCode?
1. Appreciate the Difference Between `vi` and `vim`
1. Restore `Esc` Key to Its Original Keyboard Home
1. Use `Ctrl-[` Instead of `Esc` Key (Never `Ctrl-x|z`)
1. Do the Vim Tutorial (`vimtutor`), But Beware
1. Other Recommended Learning Resources
   1. <https://openvim.com>
   1. <http://vimgenius.com> (no `s`)
   1. Vim Adventures is Strongly Discouraged
1. Start with Defaults and *Zero* Configuration
   1. Complex `.vimrc` is *Not* for Beginners
   1. Customizing `.vimrc` Requires Scripting Skills
   1. Eventually, Learn a Little Vimscript
1. Avoid Vim Pane Splitting (Use TMUX Instead)

## Day 12: Manage Users, Groups, and Permissions 

1. How Much User Stuff Do I Need to Know?
1. Create and Manage Users and Groups
   1. Use `adduser` and `useradd` to Create User
   1. Use `su - <user>` to "Login" as User
   1. Use `deluser` and `userdel` to Delete User
   1. Use `usermod` to Modify User Settings
   1. Use `addgroup` and `groupadd` to Create Group 
   1. Use `delgroup` and `groupdel` to Delete Group 
   1. Use `groupmod` to Modify User Settings
   1. Use `passwd` to Change Passwords
   1. Know the Files Involved
      1. `/etc/passwd`
      1. `/etc/shadow`
      1. `/etc/group`
   1. Use `id` to Get the User and Group Information
   1. Use `login` to Login
      1. Does Not Register as Logged in User
   1. Use `su - <user>` to Simulate a Login as Root
   1. Understand `root` Access
      1. Use `sudo` to Grant `root` (SuperUser) Access
      1. Use `sudo su -` to Get Root Shell
      1. use `doas -s` to Do Something As Root
   1. Use `who`,`w`,`whoami`,`who am i`,`last` to See Users
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

## Day 13: Commands *Are* Code, Creating Shell Scripts

1. You're Already Coding, Every Command *is* Code
   1. Commands are Really Just Functions with Arguments
1. POSIX Shell is a Universal Command Interpreter
   1. Shell Started as Bourne Shell, Now Ash and Dash
   1. Korn Shell Led to Bash and Zsh Interpreters
   1. Awk, Perl, Python, Ruby, Node Also Shell Scripting
1. Create First Script 
   1. Create a File Containing Some Commands
   1. Know What *Interpreted* Means
   1. Use `sh` Interpreter to Run Commands in File
   1. Use `bash` Interpreter to Run Same Commands
   1. Use `perl` to Attempt Same and Note Errors
   1. Use `chmod +x` to Make File Executable
   1. Add `#!/bin/sh` Shebang Line to Specify Interpreter
1. Difference Between *Running* and *Sourcing* Scripts
   1. Use Dot (`.`) or `source` to Source a Script
   1. Most Stuff Should Be in Script (Subprocesses)
   1. Some Stuff Can't Be Done Any Other Way
      1. Change Current Working Directory
      1. Modify Current Environment
1. Put Executables in Your `PATH` to Run From Anywhere
   1. Use `which` to See Which Executable Wins
   1. Understand Difference Between `which` and `type`
1. Manage Jobs and Processes
   1. A Running Program is a *Process*
   1. A Backgrounded Program a *Job*
   1. `Ctrl-z` to Background Running Process
   1. Use `jobs` to See All Background Processes
   1. Use `fg` to Bring Background Job Forward
   1. Use ` &` to Start Program in Background
   1. Use `pgrep`, `ps`, `/proc` to See Processes
   1. Use `crontab` to Schedule Jobs

## Day 14: POSIX Shell Programming Syntax

1. Know What *Syntax* Means
1. Know What *POSIX* Means
1. Know What *Script* Means
1. Know What *Program* and *Programming* Mean
1. Use `shellcheck` Throughout to Ensure POSIX
1. Functions Group Commands
   1. Know What a *Function* Is
   1. Know What an *Argument* Is
   1. Know What a *Return Value* Is
   1. Commands *are* Functions
   1. Functions *are* Commands
   1. Use `foo () {}` to Create a Function
   1. Know What a *Block* Is
   1. Functions Have Arguments
      1. Arguments Do *Not* Go in `()`
      1. Use `$1` for First Argument (And So On) 
   1. Functions Have Return Values
      1. Understand UNIX Return Values
      1. Function Return Values Same as Commands
   1. Distinguish Return Value from Output
   1. Use `return` to Set Return Value
      1. Sometime Omitting `return` is Fine
   1. Use `set -e` to Exit Program if *Anything* Fails
   1. Use `exit` to Exit Program Immediately
   1. Write Clean Functions
      1. Some Functions are Not Functions At All 
      1. Know What a *Subroutine* or *Procedure* Is
      1. Know How *True* Function Differs from Procedure
      1. Know What *Side Effects* Are
      1. Know What *Functional Programming* Is
      1. Separate Functions from *Procedure* "Functions"
1. Variables Save State
   1. Use `=` With No Surrounding Spaces to Assign Variable
   1. Like Putting Something Into Box or on List
   1. Know What *Scope* Means
   1. Know What *Global* and *Local* Mean
   1. All Variables are Global in POSIX Shell
   1. Know What *Type* Means When Programming
   1. Variables Contain Data of Different Types
      1. All Types are Ultimately the Same Binary Bits
   1. Usually Wrap Shell Variables in Quotes (`"$foo"`)
   1. Use `${}` to Disambiguate
1. Conditions Alter Flow
   1. Use `test` to Check Condition
   1. Avoid Problematic `[]` to Check Conditions
   1. Use `if` to Group Commands If True
   1. Avoid `else` Whenever Possible
   1. Use `case` to Branch Multiple Conditions
   1. Use `set -e` to Add *Exit on Any Error* Condition
   1. Know and Use *Short-Circuit Logic*
1. Loops Repeat Commands Until Condition Met
   1. Use `break` to Break Out of Any Loop
   1. Use `continue` to Start Next Iteration Early
1. Signals Communicate Between Programs
   1. Know What *Interprocess Communication (IPC)* Is
   1. Signals are Just One IPC Method
   1. Use `kill`, `pkill` to Send a Signal
   1. Use `trap` to Handle a Signal
   1. Don't Bother with `nohup` (Use `docker`, `screen`/`tmux`)
1. Some Fun Challenges to Get You Started
   1. <https://rwx.gg/lang/cha>

*More to come (boost is at least 100 days)...*
