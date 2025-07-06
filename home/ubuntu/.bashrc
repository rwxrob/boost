#!/bin/bash
# shellcheck disable=SC1090,SC1091

case $- in
*i*) ;; # interactive
*) return ;;
esac

export black="[30m"
export red="[31m"
export green="[32m"
export yellow="[33m"
export blue="[34m"
export magenta="[35m"
export cyan="[36m"
export white="[37m"
export blink="[5m"
export reset="[0m"
export clear="[2J"
export curoff="[?25h"
export curon="[?25h"
export top="[H"

# ---------------------- local utility functions ---------------------

_have() { type "$1" &>/dev/null; }
_source_if() { [[ -r "$1" ]] && source "$1"; }

# ----------------------- environment variables ----------------------

export LANG=en_US.UTF-8
export HRULEWIDTH=73
export TZ=America/New_York
export TERMINAL_BROWSER=w3m # lynx
export TERM=xterm-256color
export PYTHONDONTWRITEBYTECODE=2
export LC_COLLATE=C
export CFLAGS="-Wall -Wextra -Werror -O0 -g -fsanitize=address -fno-omit-frame-pointer -finstrument-functions"
export BAT_THEME=gruvbox-dark

# explicit to fix wierd defaults (like macOS)
export XDG_STATE_HOME="$HOME/.local/state"
export XDG_CACHE_HOME="$HOME/.cache"
export XDG_CONFIG_HOME="$HOME/.config"
export XDG_DATA_HOME="${XDG_DATA_HOME:-$HOME/.local/share}"

# gruvbox-material
export LS_COLORS="di=38;5;245:fi=38;5;223:ln=38;5;179:ex=38;5;108:*.txt=38;5;223"
export LESS="-FXR"
export LESS_TERMCAP_md=$'\e[1;33m'       # start bold (yellow)
export LESS_TERMCAP_mb=$'\e[1;35m'       # start blinking (magenta)
export LESS_TERMCAP_me=$'\e[0m'          # end bold/blinking
export LESS_TERMCAP_so=$'\e[38;5;108;1m' # start standout (green bold)
export LESS_TERMCAP_se=$'\e[0m'          # end standout
export LESS_TERMCAP_us=$'\e[4m'          # start underline
export LESS_TERMCAP_ue=$'\e[0m'          # end underline

[[ -d /.vim/spell ]] && export VIMSPELL=("$HOME/.vim/spell/*.add")

# ----------------------------- go stuff  ----------------------------

case "$(uname -m)" in
x86_64) GOARCH="amd64" ;;
aarch64) GOARCH="arm64" ;;
*)
	echo "Unsupported architecture"
	exit 1
	;;
esac && export GOARCH
export GOBIN="$HOME/.local/bin"
export GOPATH="$XDG_DATA_HOME/go"
export GOPROXY=direct
export CGO_ENABLED=0

# -------------------------------- gpg -------------------------------

GPG_TTY=$(tty)
export GPG_TTY

# ------------------------------- pager ------------------------------

if [[ -x /usr/bin/lesspipe ]]; then
	export LESSOPEN="| /usr/bin/lesspipe %s"
	export LESSCLOSE="/usr/bin/lesspipe %s %s"
fi

# ----------------------------- dircolors ----------------------------

if _have dircolors; then
	if [[ -r "$HOME/.dircolors" ]]; then
		eval "$(dircolors -b "$HOME/.dircolors")"
	else
		eval "$(dircolors -b)"
	fi
fi

# ------------------------------- path -------------------------------

pathappend() {
	declare arg
	for arg in "$@"; do
		test -d "$arg" || continue
		PATH=${PATH//":$arg:"/:}
		PATH=${PATH/#"$arg:"/}
		PATH=${PATH/%":$arg"/}
		export PATH="${PATH:+"$PATH:"}$arg"
	done
} && export -f pathappend

pathprepend() {
	for arg in "$@"; do
		test -d "$arg" || continue
		PATH=${PATH//:"$arg:"/:}
		PATH=${PATH/#"$arg:"/}
		PATH=${PATH/%":$arg"/}
		export PATH="$arg${PATH:+":${PATH}"}"
	done
} && export -f pathprepend

# remember last arg will be first in path
pathprepend \
	"$HOME/.local/bin" \
	"$HOME/.local/go/bin" \
	/usr/local/go/bin \
	/usr/local/opt/openjdk/bin \
	/usr/local/bin \
	/opt/homebrew/bin \
	"$SCRIPTS"

pathappend \
	/usr/local/opt/coreutils/libexec/gnubin \
	'/mnt/c/Windows' \
	'/mnt/c/Program Files (x86)/VMware/VMware Workstation' \
	/mingw64/bin \
	/usr/local/bin \
	/usr/local/sbin \
	/usr/local/games \
	/usr/games \
	/usr/sbin \
	/usr/bin \
	/snap/bin \
	/sbin \
	/bin

# ------------------------ bash shell options ------------------------

# shopt is for BASHOPTS, set is for SHELLOPTS

shopt -s checkwinsize # enables $COLUMNS and $ROWS
shopt -s expand_aliases
shopt -s globstar
shopt -s dotglob
shopt -s extglob

# -------------------------- stty annoyances -------------------------

stty -ixon # disable control-s/control-q tty flow control

# ------------------------------ history -----------------------------

export HISTCONTROL=ignoreboth
export HISTSIZE=5000
export HISTFILESIZE=10000

set -o vi
shopt -s histappend

# --------------------------- smart prompt ---------------------------
#                 (keeping in bashrc for portability)

PROMPT_LONG=20
PROMPT_MAX=95
PROMPT_AT=@

__ps1() {
	local P='$' dir="${PWD##*/}" B countme short long double \
		r='\[\e[31m\]' h='\[\e[34m\]' \
		u='\[\e[33m\]' p='\[\e[34m\]' w='\[\e[35m\]' \
		b='\[\e[36m\]' x='\[\e[0m\]' \
		g="\[\033[38;2;90;82;76m\]"

	[[ $EUID == 0 ]] && P='#' && u=$r && p=$u # root
	[[ $PWD = / ]] && dir=/
	[[ $PWD = "$HOME" ]] && dir='~'

	B=$(git branch --show-current 2>/dev/null)
	[[ $dir = "$B" ]] && B=.
	countme="$USER$PROMPT_AT$(hostname):$dir($B)\$ "

	[[ $B == master || $B == main ]] && b="$r"
	[[ -n "$B" ]] && B="$g($b$B$g)"

	short="$u\u$g$PROMPT_AT$h\h$g:$w$dir$B$p$P$x "
	long="${g}╔$u\u$g$PROMPT_AT$h\h$g:$w$dir$B\n${g}╚$p$P$x "
	double="${g}╔$u\u$g$PROMPT_AT$h\h$g:$w$dir\n${g}║$B\n${g}╚$p$P$x "

	if ((${#countme} > PROMPT_MAX)); then
		PS1="$double"
	elif ((${#countme} > PROMPT_LONG)); then
		PS1="$long"
	else
		PS1="$short"
	fi

	if _have tmux && [[ -n "$TMUX" ]]; then
		tmux rename-window "$(wd)"
	fi
}

wd() {
	dir="${PWD##*/}"
	parent="${PWD%"/${dir}"}"
	parent="${parent##*/}"
	echo "$parent/$dir"
} && export wd

PROMPT_COMMAND="__ps1"

# ------------------------------ aliases -----------------------------
#      (use exec scripts instead, which work from vim and subprocs)

unalias -a
alias ip='ip -c'
alias '?'=duck
alias '??'=google
alias ls='ls -h --color=auto'
alias free='free -h'
alias tree='tree -a'
alias df='df -h'
alias neo="neo -D -c gold"
alias diff='diff --color'
alias grep='grep --color=auto'
alias view='vi -R' # which is usually linked to vim
alias clear='printf "\e[H\e[2J"'
alias c='printf "\e[H\e[2J"'
alias env='env -u LESS_TERMCAP_mb -u LESS_TERMCAP_md -u LESS_TERMCAP_me -u LESS_TERMCAP_so -u LESS_TERMCAP_se -u LESS_TERMCAP_us -u LESS_TERMCAP_ue'
alias neo="neo -D -c gold"
alias more="less -R"
alias gitl="git log -n 5 --graph --decorate --oneline"

set-editor() {
	export EDITOR="$1"
	export VISUAL="$1"
	export GH_EDITOR="$1"
	export GIT_EDITOR="$1"
	alias vi="\$EDITOR"
}
_have "vim" && set-editor vi
_have "nvim" && set-editor nvim

# ------------- source external dependencies / completion ------------

_have gh && . <(gh completion -s bash)
_have pandoc && . <(pandoc --bash-completion)
_have yq && . <(yq completion bash)
