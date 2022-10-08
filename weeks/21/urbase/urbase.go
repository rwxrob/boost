package urbase

import "fmt"

func Run() {
	fmt.Println("All your base ...")
	fmt.Println("Base: 2   8  10  16")
	for i := 0; i < 17; i++ {
		fmt.Printf("  %5b %3o %3d %3x\n", i, i, i, i)
	}
}
