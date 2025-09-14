package main

import (
	"fmt"
	"rocketseat-golang-training/03-TypesAndVariables/booleans"
	deferr "rocketseat-golang-training/03-TypesAndVariables/defer"
	"rocketseat-golang-training/03-TypesAndVariables/forr"
	"rocketseat-golang-training/03-TypesAndVariables/numbers"
	Range "rocketseat-golang-training/03-TypesAndVariables/rangee"
	"rocketseat-golang-training/03-TypesAndVariables/strings"
	switches "rocketseat-golang-training/03-TypesAndVariables/switch-case"
)

func main() {
	// !! All variables always have a zero value.
	// !! [Const] are immutable variables.
	strings.StringsVar()
	fmt.Println("===============")
	booleans.BoolVar()
	fmt.Println("===============")
	numbers.ShowNumbers()
	fmt.Println("===============")
	forr.ForAndWhile()
	fmt.Println("===============")
	Range.Rangee()
	fmt.Println("===============")
	switches.SwitchCase()
	fmt.Println("===============")
	deferr.ShowDefer()
}
