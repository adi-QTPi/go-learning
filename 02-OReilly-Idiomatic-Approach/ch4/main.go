package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Blocks, shadows and Control Function")

	{
		x := 10
		if x > 5 {
			fmt.Println(x) //10
			x := 23        //not reassigning the outer variable. Initialising new one.
			fmt.Println(x) //23 (outer variable gets shadowed)
		}
		fmt.Println(x) //10
	}

	hr()

	{
		z := 1
		if z < 4 {
			fmt.Println("z", z)
			z, y := 10, 2 //looks like re-assign, but actually SHADOWING happens.
			fmt.Println("y", y)
			fmt.Println("z", z)
		}
		fmt.Println("z", z)
	}

	hr()
	// {
	// 	fmt.Println("here we can use fmt package")
	// 	fmt := "hel"
	// 	fmt.Println("hey") //would cause compiler issue
	// }

	{
		//redifining true
		shiv := true
		true := "shiv"
		fmt.Println(true, shiv)
	}

	hr()
	hr()

	{
		n := rand.Intn(10)
		if n < 5 {
			fmt.Println("thas too low man")
		} else {
			fmt.Println("cool")
		}
	}

	hr()

	{
		fmt.Println("take 2")
		if n := rand.Intn(20); n < 10 {
			fmt.Println("thas too low man", n)
		} else {
			fmt.Println("cool", n)
		}
	}

	hr()

	{
		sl := []int{1, 2, 3, 4, 5}

		for k := range sl { //if nothing given then key is taken by default.
			fmt.Println(k)
		}
	}

	hr()

	{
		m := map[string]string{
			"one":   "hi",
			"two":   "hello",
			"three": "pranam",
			"four":  "bye",
		}
		for i := 0; i < 5; i++ {
			fmt.Println("Loop", i+1)

			//every loop map items printed in diff order
			for k, v := range m {
				fmt.Println("Key = ", k, "Value = ", v)
			}
		}
	}

	hr()
	hr()

	{
		fmt.Println("Exercises")

		var sl []int
		for range 100 {
			randomNumber := rand.Intn(100)
			sl = append(sl, randomNumber)
		}
		for _, v := range sl {
			switch {
			case v%2 == 0:
				fmt.Println("Two!")

			case v%3 == 0:
				fmt.Println("Three")

			case v%3 == 0 && v%2 == 0:
				fmt.Println("Six!")

			default:
				fmt.Println("Never Mind")
			}
		}
	}

	hr()

	{
		total := 10
		for i := 0; i < 100; i++ {
			total++
			fmt.Println(total)
		}

		fmt.Println(total)
	}
}

func hr() {
	fmt.Println("=====break=====")
}
