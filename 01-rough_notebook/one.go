package main

import (
	"fmt"
)

func hello() (x, y int) {
	fmt.Println("hi there")
	x = 3
	y = 2
	return
}

func main() {
	a, b := hello()
	fmt.Printf("the numbers are %v and %v. \n", a, b)
}
