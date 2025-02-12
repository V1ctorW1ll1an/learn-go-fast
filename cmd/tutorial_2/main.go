package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var printValue string = "Hello world!"
	printMe(printValue) // >> Hello world!

	var numerator int = 11
	// var denominator int = 0
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		log.Fatal(err) // >> 2024/11/24 17:53:18 Cannot Divide by zero exit status 1
	} else if remainder == 0 {
		fmt.Println(result) // >> 5
	} else {
		fmt.Println(result, remainder) // >> 5 1
	}

	switch { // switch dont need break in go
	case err != nil:
		log.Fatal(err) // >> 2024/11/24 17:53:18 Cannot Divide by zero exit status 1
	case remainder == 0:
		fmt.Println(result) // >> 5
	default:
		fmt.Println(result, remainder) // >> 5 1
	}

	switch remainder {
	case 0:
		fmt.Println("the division was exact")
	case 1, 2:
		fmt.Println("the division was close")
	default:
		fmt.Println("the division was not close")
	}

	// for loops
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// for continued
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// forever
	for {
		println("forever")
		break
	}

	// stack defer
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
