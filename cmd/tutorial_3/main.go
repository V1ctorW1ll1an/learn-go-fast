package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	// measure time
	var n int = 1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("total time without preallocation: %v\n", timeLoop(testSlice1, n))
	fmt.Printf("total time with preallocation: %v\n", timeLoop(testSlice2, n))
	// []arrays
	/*
	* Fixed length
	* same type
	* indexable
	* contiguous in memory
	 */
	var intArr [3]int32 // 12 bytes 3 * 4bytes (int32 size)
	intArr[1] = 1
	intArr[2] = 20
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0]) // >> 0x14000104020
	fmt.Println(&intArr[1]) // >> 0x14000104024
	fmt.Println(&intArr[2]) // >> 0x14000104028

	elementSize := unsafe.Sizeof(intArr[0])
	// calc the index (using a offset)
	offset := 2
	index := int(uintptr(offset) * elementSize / elementSize)
	fmt.Println(intArr[index]) // Output: 1

	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7) // [4, 5, 6, 7, *, *]
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	// spread operator
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	// especify the length and the capacity
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("the length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))
	fmt.Println(intSlice3)

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8) // keys are string and values are uint8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Victor": 100, "Willian": 99}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Victor"]) // case sensitive
	fmt.Println(myMap2["sousa"])  // default value of uint8 = 0

	var age, ok = myMap2["Foo"]
	fmt.Println(age, ok) // >> 0 false

	age, ok = myMap2["Victor"]
	fmt.Println(age, ok) // >> 100 true

	if ok {
		fmt.Printf("the age is %v\n", age)
	} else {
		fmt.Println("invalid name")
	}

	delete(myMap2, "Willian")
	fmt.Println(myMap2)

	myMap2["Sousa"] = 200

	// important, order can be change in map
	for name := range myMap2 {
		fmt.Printf("name: %v | age: %v\n", name, myMap2[name])
	}
	// or
	for name, age := range myMap2 {
		fmt.Printf("name: %v | age: %v\n", name, age)
	}

	// for loops
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// infinit
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
