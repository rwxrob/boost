package main

/* WRONG: creates wrong number of arguments
func ExampleMain() {
	main()
	// Output:
	// Hi there!
}
*/

func ExampleSayHi() {
	SayHi()
	SayHi("Rob")
	SayHi("Doris")
	// Output:
	// Hi there!
	// Hi Rob!
	// Hi Doris!
}
