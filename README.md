# SKILSTAK Beginner Boost Series

## "What's a 'Beginner Boost'?"

These days the Boost is broken into three parts:

* 18 Weeks (May/4-Aug/27): Tech Tools, Practices, and Programming (I)
* 18 Weeks (Sep/3-Dec/31): Infrastructure Engineering (II)
* 10 Weeks (Jan/7-Apr/22): Containers and Kubernetes (III)

The first part focuses on upskilling a basic computer user into
a entry-level engineer proficient and terminal native with strong
command-line skills and basic programming ability. 

The second part focuses on the regular tasks asked of any professional
infrastructure engineer supporting an on-prem cloud of Linux servers.

The third part specifically focuses on the cloud-native approach to
infrastructure engineering and applications development. Note that the
focus is always on building *on-prem* infrastructure, not cloud
providers. Understanding cloud providers is easy after building and
maintaining your own cloud at home, plus it will get you a job faster if
you do.

Here's a sample of some specific topics:

* Use a Mac iTerm2 terminal with brew
* Use Windows terminal with Git Bash on Windows
* Setup Ubuntu Linux virtual machine with VirtualBox
* Use SSH to create terminal connections and tunnel
* Use containers (Docker) from terminal on Linux servers
* Set up basic networking for virtual machines and containers
* Script powerful "one-liner" commands with filters
* Code utilities in Bash and Go
* Manage and share knowledge and source code
* Connect with professional tech communities to seek opportunities

## "When are they?"

Sessions take place every Saturday 12:00-17:00 (New York time) with
"office hours" Sundays 12:00-15:00. Sometimes they will be moved or
cancelled depending on personal conflicts and holidays. Sometimes we'll
go slightly longer depending on the content. Check the
<https://twitch.tv/rwxrob/schedule> for when.

## "Five hours? Do you take breaks?"

* 12 minutes off
* 54 minutes on

Yes, we pretty strictly take a 12 minute break ever 54 minutes (which
happens to be the exact times required by Twitch to ensure we are not
interrupted with ad breaks by taking them intentionally instead).

## Before You Begin

Please make sure you have a computer with at least the following minimum
specs:

* Windows 10 or Mac x86
* 4 Core (x86)
* 8 GB Memory (RAM)
* 500 GB Drive (100 GB Free)

Also make sure you meet the SKILSTAK minimum
[requirements](requirements) — especially installing VirtualBox onto a
computer on which you have admin permissions. We are hear to help if you
have questions about any of these requirements. Check the [FAQ](faq) for
answers to your questions or ask them in the
[Discord](https://discord.gg/9wydZXY) server or live during Sunday [AMA
sessions](https://twitch.tv/rwxrob/schedule). Videos of sessions are
always saved to [YouTube](https://youtube.com/rwxrob) and posted to
[Twitter](https://twitter.com/rwxrob).

## What changed since 2021?

* Focus on cloud-native infrastructure engineer and security jobs
* Increased duration to 18 weeks (5 weeks before)
* No longer daily, now 5-hours on Saturday (usually)
* Single, long-form video with full coverage of material
* Asking to repeat during week what was done to master
* More project based, even less theory
* More "hand holding" since we are "pairing" with beginners
* Multiple programming languages covered 
* More networking, but still just "How does Internet work?"
* Certain topics moved to Infrastructure Engineering Boost
* Focus on Linux Ubuntu server VM on local machine
* Lighter on containers (Docker) and no Kubernetes (moved to own Boost)

## Week 0: Wednesday, May 4, 11:11:11 AM, 2022

*🎉 Let's get this party started! 🎉*

Hour 1:

* Hello, friends. May the fourth be with you! (5m)
* Dedication to Aaron Swartz (5m)
  * "The Internet's Own Boy" (watch it this week) (1m)
* Start downloading VirtualBox and Ubuntu Server (5m)
* Prepare your computer for command-line terminal work
  * Windows
    * Install and customize Windows Terminal (5m)
    * Install Git Bash (5m)
  * Mac
    * Install and customize iTerm2 (5m)
    * Install `brew`, coreutils, etc. (5m)
* Discuss stuff while things are installing (1m)
* Know who we are and decide to join us (or not) (1m)
  * "Why are you doing this for free?" (2m)
  * Know the history of rwxrob and #SKILSTAK (2m)
  * Decide if this community is for you (2m)
  * Consider other communities as well (2m)

TOTAL MINUTES: 46

--- 

Hour 2:

* Known what we offer and why you are here (1m)
  * Learn to love UNIX command line ("this is the way") (1m)
  * Learn to live UNIX/Linux as your daily driver (1m)
  * Learn to code in eight languages:
    * Learn to identify the best tool/language for a give task (1m)
    * Markdown, JSON, YAML (1m)
    * Bash, Go (1m),
    * HTML, CSS, JavaScript for Web documents (not apps) (1m)
  * Learn to manage knowledge and source code with Git (1m)
  * Learn to learn like a hacker(1m)
  * Learn to self-assess and know your value (1m)
  * Learn prescience and intuition (1m)
  * Learn to connect and find work (1m)
* Accept what we require in return
  * No Sith Lords (2m)
  * Community over competition (1m)
  * Take responsibility (1m)
  * Give credit and benefit of the doubt(1m)
  * Forgive and forget (1m)
  * Pay it forward (1m)
  * Save the world (1m)
  * Subscribe, sponsor, and share (1m)
* Setup essential accounts and know why
  * Get Twitch account (1m)
    * Do not change your ID/name for duration of Boost (1m)
    * Earn rwxrob community "rep" points (2m)
  * Get YouTube account (1m)
  * Get optional Twitter account (1m)
  * Get optional Discord account (1m)
  * Get optional ProtonMail account (5m)
  * Get GitHub account (5m) 
    * Why not GitLab or other? (1m)
    * "Don't you hate Microsoft? (1m)
* Help us know who *you* are (if you want)
  * <https://boosted.skilstak.io> (1m)
  * No, it's not a Google form (1m)
  * Never used for solicitation, just stats (1m)

TOTAL MINUTES: 41

----

Hour 3:

* ¡Viva la revolución! Defend your right to learn! (2m)
  * Paulo Friere, Pedagogy of the Oppressed (2m)
  * Sir Ken Robinson, *Out of Our Minds* (2m)
* Activate your autodidact, learn like the best
   * Remember how to learn
     * You already know everything you need, just remember it (1m)
   * Who helped you learn and love learning? (1m)
     * Mom (4m)
     * Mr. Poulson (2m)
     * Mr. Lundquist (2m)
     * Mr. Bevans (2m)
   * What obstacles have been in your way? (2m)
   * How do *you* learn best? (2m)
   * When do you remember learning the best? (1m)
   * Start a personal/professional learning network
     * GitHub (1m)
     * Discord (1m)
     * Slack (1m)
     * Twitter (1m)
     * Twitch (1m)
     * What about the others? (1m)
   * Organize your learning
     * Get GitHub (1m)
     * Learn CommonMark Markdown (1m)
     * Start a Zettelkasten (5m)
   * Use effective learning techniques
     * Keep simple, silly, and fun so you remember (2m)
     * "Repetition is the mother of learning" (1m)
       * Repeat what we did together during the week (1m)
     * Distinguish side quests from main mission (1m)
     * Consider the RWX Method (1m)
       * Give yourself `rwx` permissions to learn (1m)
       * Read, research, and repeat things (2m)
       * Write about your learning (2m)
       * Explore, experiment, and experience (2m)

TOTAL MINUTES: 49

---

Hour 4:

---

Hour 5:

## Week 1: Saturday, May  7, 2022

*Become a shell native. This is the way, the UNIX way.*

## Week 2: Saturday, May 14, 2022

*Editing text from Ed/Vi/Vim command line.*

## Week 3: Saturday, May 21, 2022

*Boosting your command line skills with bash.*

## Week 4: Saturday, May 28, 2022

*Working with structured data: JSON, YAML, TOML, etc.*

## Week 5: Saturday, June  4, 2022

*Accessing the Internet from the command line*

## Week 6: Saturday, June 11, 2022

*Networking, just the essentials: How does the Internet work?*

## Week 7: Saturday, June 18, 2022

*Secure connections with Secure Shell and GPG*

## Week 8: Saturday, June 25, 2022

*Getting started with git version management.*

## Week 9: Saturday, July  2, 2022

*Using containers (Docker)*

## Week 10: Saturday, July  9, 2022

*Programming fundamentals* 

## Week 11: Saturday, July 16, 2022

*Programming fundamentals.*

## Week 12: Saturday, July 23, 2022

*Programming fundamentals.*

## Week 13: Saturday, July 30, 2022

*Programming fundamentals.*

## Week 14: Saturday, August  6, 2022

*Programming fundamentals.*

## Week 15: Saturday, August 13, 2022

*Programming fundamentals.*

## Week 16: Saturday, August 20, 2022

*Web doc/site development: create a portfolio page*

## Week 17: Saturday, August 27, 2022

*Connecting with professional community for fun and profit*

* Find online and offline communities

