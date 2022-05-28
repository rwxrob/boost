Week 2 focuses on UNIX/Linux origins and the terminal, command-line
interface and basic file system navigation. You should be comfortable
installing and removing and reinstalling Linux Ubuntu Server VMs and any
apt packages before next week.

* Primary responsibilities of any mentor
  1. Help you know what to learn
  1. Help you know where to learn it (curate resources)
  1. Help you when you get stuck

* What is a "shell" and why do I care?
* What is the difference between "terminal", "CLI", and "shell"?
* What is a REPL? (Read, Evaluate, Print, Loop)
* Install and configure terminal software (MS Terminal, iTerm2)
* What is UNIX and Linux?
* What is GNU? Who invented Linux?
* Basic system update and package management and installation
* When to use apt vs apt-get?
* What is a "package"?
* How do I install software?
* Really simple and essential shell commands
* What do the character's in the prompt mean?
* What is an "inode"? Why everything is "just a file"?
* What are UNIX/Linux permissions?
* What does `.` and `..` mean?
* How do I know that `apt` packages are safe?
* What is a shell "history"?

## Commands Used

* `man` - show manual information about a command
* `sudo` - do it as root (superuser)
* `apt` - use interactively only (use `apt-get` in scripts)
* `sudo apt update` - update all the sources for packages
* `sudo apt upgrade` - upgrade *all* packages to latest version
* `apt search ^neo` - search for all package starting with `neo`
* `sudo apt install neofetch` - install `neofetch` and dependencies
* `sudo apt remove neofetch` - remove `neofetch`
* `sudo apt autoremove` - automatically remove unused packages
* `ls` - list the files in the (current) directory
* `ls -al` - list all the files including hidden (begin with `.`)
* `hostname` - display name of host computer
* `pwd` - print working directory
* `cd foo` - change into the `foo` directory
* `cd`, `cd ~` - change back to the home directory
* `cd ..` - change into the relative parent directory
* `cd ../..` - change into the relative parent of the parent directory
* `cd -` - change to previous directory
* `cd /` change to the root directory

Related:

* https://github.com/microsoft/terminal/releases (no app store needed)
* https://youtu.be/crtZHNclndQ (cool retro terminal)
* https://www.linuxcommand.org
* https://replit.org
* https://medium.com/@rwxrob/problem-with-repl-it-415f885164f
* "AT&T Archives: The UNIX Operating System" https://youtu.be/tc4ROCJYbm0
* "It's a UNIX system" https://youtu.be/dxIPcbmo1_U 
* "Is it time to rewrite the OS in Rust" https://youtu.be/HgtRAbE1nBM
* https://youtu.be/ctGpRWCi8QU
* https://git-scm.org
