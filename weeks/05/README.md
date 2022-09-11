This day is almost entirely about finding files matching criteria
including specific content in those files, one of the most important
skills of any ops engineer, hacker, or developer. Next time we will
start editing files.

* How do I find a file from the UNIX/Linux command line? (20220612172124)
* Difference between `locate`, `whereis`, and `find`? (20220612172510)
* Find all the files in a directory and subdirectories? (20220612174044)
* Find only the files recursively? (20220612174830)
* Find only the "dot" files beginning with dot? (20220612174911)
* Find only directories beginning with dot? (20220612175454)
* Find all file created or modified in last hour? (20220612175715)
* Find anything modified in last five minutes (20220612180122)
* Diff between `/tmp` and `/dev/shm`? (20220612182646)
* How do I silence errors? (20220612183206)
* How do I save output to a file? (20220612183423)
* How do I only show top five lines? (20220612184214)
* How to avoid using `cat`? (20220612185426)
* What is a UNIX filter / the UNIX philosophy? (20220612185541)
* How to append lines to a file? (20220612190759)
* How to use `tee` command?
* How do I filter lines with keyword from file? (20220612193202)
* What's a regular expression and why do I care? (20220612194251)
* Why is it better to allow rather than deny? (20220612194542)
* How to play regular expression golf? (20220612200534)
* What's a glob/globstar? (20220612204614)

## Commands Used

* `find . ` - sort of the same as `ls -l1`
* `find . -ls` - sort of the same as `ls -l1`
* `find . -name '???'` - find files/dirs with three letter name
* `find . -path '.git'` - find files/dirs with `.git` anywhere in path
* `head -5 foo` - show top five lines of `foo` file
* `tail -5 foo` - show bottom five lines of `foo` file
* `tac foo` - reverse lines of `foo` file
* `wc -l` - print count of lines
* `nl` - add line numbers to output
* `tee /tmp/foo` - both writes to a file `foo` and to stdout
* `>` - (over) write to file
* `>>` - append to file
* `|` - connect stdout to stdin
* `<` - send file to stdin
* `lolcat` - colorize things
* `cowsay` - fun way to show output

## Related

* "Missing Semester" on YouTube (Regular Expressions)
* Regex Golf https://alf.nu/RegexGolf
* O'Reilly Mastering Regular Expressions
* Geek Code (but a little rated M)
