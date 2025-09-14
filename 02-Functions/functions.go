package functions

import "fmt"

func MathsFunc() {
	a, b := swap(10, 20)
	fmt.Println("Result swap: ", a, b)
	fmt.Println("=============")
	rest, rem := split(15, 12)
	fmt.Println("Result split: ", rest, rem)
	fmt.Println("=============")
	x := toAdd(255)(1)
	fmt.Println("Result add: ", x)
}

func swap(a, b int) (int, int) {
	return b, a
}

func split(c, d int) (rest int, rem int) {
	rest = c / d
	rem = c % d
	return rest, rem
}

func toAdd(a int) func(int) int {
	// !! First Step: I declared a normal function! [**toAdd(a int)**]
	// !! Second Step: It returns another function! [**func(int)**]
	// !! Third Step: All this returns an integer! [**int**]
	return func(b int) int {
		// **Anonymous Function
		return a + b
	}
}
