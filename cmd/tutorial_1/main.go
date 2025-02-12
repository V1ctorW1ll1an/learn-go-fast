package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum) // >> 32768

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum) // >> 1.23456789e+07

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result) // >> 12.1

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // >> 1
	fmt.Println(intNum1 % intNum2) // >> 1

	var myString string = "Hello" + " " + "world" + "!"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("Y")) // >> 1

	var myRune rune = 'a'
	fmt.Println(myRune)

	/*
	* Default value (0) to => uint, uint8, uint, uint16, uint32, uint64, int, int8,
	* int16, int32, int64, float32, float64, rune
	* Default value ('') to => string32, string64
	* Default value (false)  to => bool
	* Default value (nil)  to => error
	 */
	var intNum3 rune
	fmt.Println(intNum3)

	// implicit type inference
	var myVar = "text"
	fmt.Printf("%T\n", myVar) // >> string

	myVar2 := "text2"
	fmt.Printf("%T\n", myVar2) // >> string

	// mult initilize variables
	var1, var2 := 1, 2
	fmt.Println(var1, var2) // >> 1 2

	// good practice is to explicitly inform the type of the
	// variable when we call a function
	var myString1 string = foo()
	fmt.Println(myString1)

	// constants
	const myConst string = "const value"
	fmt.Println(myConst) // >> const value

	const pi float32 = 3.1415
	fmt.Println(pi) // >> 3.1415

}

func foo() string {
	return "foo bar"
}
