// !! There can't be any code before the package mention, just comment!
// !! Package main cannot be imported!
// !! The main package will always be used to build a Go project!
package main

import (
	"fmt"
	PAP "rocketseat-golang-training/01-PublicAndPrivate"
	functions "rocketseat-golang-training/02-Functions"
	
	/*
		!! The "internal" package can only be accessed by other files above its hierarchy in the same package.
		internal "rocketseat-golang-training/01-PublicAndPrivate/internal/internalVar"
	*/)

func main() {
	fmt.Println("==========================")
	PAP.PublicAndPrivate()
	// internal.HiddenByInternal
	fmt.Println("==========================")
	functions.MathsFunc()
	fmt.Println("==========================")
}
