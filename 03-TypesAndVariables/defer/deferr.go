// !!The [defer] always happens like this: last one in, first one out.
// !!The scope of [defer] is tied to a function

package deferr

import "fmt"

func ShowDefer() {
	doDefer()
}

func doDefer() {
	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)
}
