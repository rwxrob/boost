# ⚡ SKILSTAK Beginner Boost 2022 - Week 13

Bash configuration day.

## Covered

* Why are they called dot files?
* Use initial `.` to "hide" a file or directory
* `ls -la` to see all hidden
* Be careful who you show your dotfiles to?
* How should I manage my dotfiles?
* Difference between "login/interactive shell"
* Old school home dir organization vs modern standards
* Adding to `/etc/skel` for all new users
* When to use an "alias", "function", or "script"?
* Why use a Bonzai monolith instead of scripts?
* Use "\" to disable aliases and functions
* Use `uname` to see info about system?
* When to use `[[ -t 0 ]]` vs `case $-`
* Use `shopt -o` to see the value of a shell option
* Use `set -o ` to change the value of a shell option
* Grok the difference between `shopt` and `set`
* Use `stty stop undef` to disable accidental terminal suspend
* Grok backgrounding (suspending) processes

## Standard (expected) files and directories

* `.profile` - original login script
* `.bashrc` - main bash configuration file
* `.config` - user (local) configurations (`os.UserConfigDir`)
* `.cache` - user (local) cache (`os.UserCacheDir`)
* `.local` - user (local) localized tree similar to `/`
* `.local/bin` - modern place to put localized executables
* `bin` - old-school place to put localized executables
* `tmp` - old-school to put temporary files
* `cp -L` - follow symbolic links

## Next Week

* ex/vi(m)
* tmux
* scripts
* lynx/w3m

* ssh
* gpg
* git
* gh
* GitHub
* ln
* custom installers
