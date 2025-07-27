package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is the 2nd chapter...")
	//exercise 1
	var i int = 20
	fmt.Println("i = ", i)
	var f float64 = float64(i)
	fmt.Println("f = ", f)

	//exercise 2
	const value = 32
	var i2 int = value
	var f2 float64 = value
	fmt.Println("i2 = ", i2)
	fmt.Println("f2 = ", f2)

	//exercise 3
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807

	b++
	smallI++
	bigI++
	//you surpass the max value, then you cycle back to the min value for that data type

	fmt.Printf("b = %v ; smallI = %v ; bigI = %v \n", b, smallI, bigI)
}
