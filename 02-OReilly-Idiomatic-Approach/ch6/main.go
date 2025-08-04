package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	hr()
	{
		variable := 10
		pointer := &variable

		fmt.Println(pointer) //prints hex value (the address)
		fmt.Println(variable)
		fmt.Println(*pointer)
	}
	hr()
	{
		var p *int
		fmt.Println(p == nil) //prints true
		// fmt.Println(*p)       //panic
	}
	hr()
	{
		pt := new(int)
		fmt.Println(pt == nil) //prints false
		fmt.Println(*pt)       //doesn't panic returns 0
		fmt.Println(pt)
	}
}

func hr() {
	fmt.Println("===break===")
}
