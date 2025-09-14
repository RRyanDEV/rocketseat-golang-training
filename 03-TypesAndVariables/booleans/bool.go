package booleans

import "fmt"

func BoolVar() {
	// By default it is already set to "False".
	bigger := 10 > 5
	smaller := 10 < 5

	fmt.Println("10 is greater than 5? ", bigger)
	fmt.Println("10 is greater than 5? ", smaller)
}
