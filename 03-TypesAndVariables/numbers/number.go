package numbers

import "fmt"

// **Int8:** Integer of 8 Bytes
// **Int16:**: Integer of 16 Bytes 
// **Int32:** Integer of 32 Bytes 
// **Int64:** Integer of 64 Bytes

func ShowNumbers() {
	Integers()
	Floats()
}

func Integers() {
	var yearsOld int = 21
	var accountant int32 = 160
	var speedOfLight int64 = 299793458

	fmt.Println("yearsOld:", yearsOld)
	fmt.Println("Accountant:", accountant)
	fmt.Println("Speed Of Light:", speedOfLight)
	fmt.Println("------")
}

func Floats() {
	// var floatNumber float32 = 4.444
	var pi float64 = 3.14
	var raio float64 = 2.5
	var area = pi * raio * raio

	// fmt.Println("Float Number:", floatNumber)
	fmt.Println("Circle Area:", area)
}
