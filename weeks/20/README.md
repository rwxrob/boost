# ⚡ SKILSTAK Beginner Boost 2022 - Week 20

"Magic eightball, should I never, ever use 'getopt' again?" This week we
make a simple tool to demonstrate command line argument parsing (args)
and get random with a magic eightball tool. We factor out stuff from
"greet" into its own package module so we can reuse it, and face
everything with that including use of "go work", "go get ..." with a
branch and more.

Covered:

* Basic idea of what functional programming paradigm
* Use of `args []string` vs `args ...string` and `foo(os.Args...)`
* Difference between "arrays" and "slices"
* Why I am not doing anything with "getopt"
* What about `bonzai`? (We'll get there.)
* Example-based testing vs "testing" package testing
* Interval notation
* Factoring out commonly used code into own library modules
* Using `go work` to avoid package dependency problems
* Using `go get` to update dependencies
* Learned to use `@2022` because of `boost` repo branch
* Finished magic `eightball`
