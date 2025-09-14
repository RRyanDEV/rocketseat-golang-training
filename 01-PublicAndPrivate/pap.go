package publicandprivate

import (
	"fmt"
	internal "rocketseat-golang-training/01-PublicAndPrivate/internal/internalVar"
)

func PublicAndPrivate() {
	PublicFunc()
	privateFunc()
	PrintInternalVar()
}

func PublicFunc() {
	fmt.Println("Public Function starts with upper letter case")
}

func privateFunc() {
	fmt.Println("private functions start with a lowercase letter")
}

func PrintInternalVar() {
	fmt.Println(internal.HiddenByInternal)
	//!! If I remove this import, the variable will be unavailable
}
