package forr

import "fmt"

func ForAndWhile() {
	LoopFor()
	While()
}

func LoopFor() {

	for range 10 {
		// !! Short declaration of For!
		fmt.Println("dentro")
	}

	// **FOR** initialization; expression; end iteration;
	sum := 0
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		sum += i
	}
}

func While() {
	//While "does not exist" in Go
	// To instantiate a While loop, use FOR
	num := 0
	for num < 20 {
		fmt.Println(`Loop do "While"`)
		num += 1
	}
	fmt.Println(num)
}
