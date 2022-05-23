This week is focused on getting connected with Secure Shell and talking
about the methods and considerations when connecting to a "remote"
computer even if that computer inside your host computer.

* START (20220522170124)

* Help us with timestamps (20220522171336)
* Boost sessions moved to Sunday (20220522171712)
* How to connect to our VM (Secure Shell)? (20220522173234)
* What is NAT and Bridged? (20220522174202)
* What is an IP Address? (20220522174927)
* What is `git-bash` and why would I need it? (20220522175639)

* BREAK (20220522180935)

* Why can't I use the 127 address? (20220522182241)
* How can you test ssh connection to localhost? (20220522182905)
* What is a "user" and "userid"? (20220522183516)
* How can see stuff scrolling off screen? (20220522184810)
* What package needs to be installed to connect with ssh?(20220522184900)
* What is the correct `ssh` connection address? (20220522190058)
* What about `cmd` (DOS)? (20220522191246)

* BREAK (20220522191428)

* What about `putty` for ssh access? (20220522192740)
* What are your favorite Windows Terminal settings? (20220522194907)
* Should I use VSCode and how? (20220522200053)
* What is JSON? (20220522201201)

* BREAK (20220522202837)

* Remote access considerations? (20220522204355)
* Hack for "good", honeypots are more fun? (20220522205233)
* How can I get a VM on the Internet? (20220522205617)
  * DigitalOcean.com
  * GCP
  * AWS
  * Vultr.com
  * linode.com
  * (be careful out there)
* How to start VM "headless"? (20220522211658)
  * From the GUI
  * From (DOS) command line
* What is minimum Vi/m that I need? (20220522212640)
* Why is `nano` the devil? (20220522214104)

## Commands

* `ip a` - show all IP addresses (`ipconfig /all`, on Windows)
* `clear` - clears the screen
* `which ssh` - display full path to ssh program
* `type ssh` - display what type of thing it is
* `who` - display who is logged in an how
* `id` - display user and group names and ids for self
* `w` - display logger version of who is logged in
* `whoami` - print effective user name/ID
* `users` - short name of all logged in users
* `last` - summary of last logged in users
* `exit` - exit the current program or login or shell
* `| less` or `| more` - see scrolled output in terminal (`q` to quit)
* `<Ctrl>-c` - interrupt whatever (exit)
* `<Ctrl>-d` - send "end of data/file"
* `sudo apt install openssh-server` - install ssh server (if not)
* `vboxmanage startvm <VMNAME> --type headless` - start vm headless
* `vboxmanage list runningvms` - list running vms
* `vboxmanage list vms` - list all vms
* `vboxmanage controlvm <VMNAME> poweroff` - power off headless vm
* `vim` - basic vim https://rwx.gg/visurvive
* `nano` - just so you understand nano editor is a thing

## Related

* https://www.virtualbox.org/manual/ch06.html
* https://try.digitalocean.com/freetrialoffer/
* https://cloud.google.com/compute/docs/create-linux-vm-instance
* https://www.how2shout.com/linux/use-command-line-to-create-list-start-stop-virtualbox-vms/
