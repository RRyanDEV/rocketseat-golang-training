package switchcase

import (
	"fmt"
	"time"
)

func SwitchCase() {
	do(1)
	doo(2)
	fmt.Println(isWeekend(time.Now()))
}

// !! The [break] is already posted automatically, without the need to type it
// !! The [fallthrough] is used to make the next statement happen, even if the condition does not reach it.

func do(x int) {
	switch x {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Default")
	}
}

func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}

func doo(x int) {
	switch {
	// The switch can happen without the conditions, as long as it happens in the case
	case x == 1:
		fmt.Println(1)
		fallthrough
	case x == 2:
		fmt.Println("Whithout Case")
	default:
		fmt.Println("Default")
	}
}
