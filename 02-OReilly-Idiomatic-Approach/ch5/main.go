package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Chapter5 -> functions")
	hr()
	fmt.Println("Variadic functions")
	fmt.Println(addTo(2, 1, 2, 3, 4))
	fmt.Println(addTo(2, []int{1, 2, 3}...))
	hr()
	ans, remainder, error := divideWithReturn(2, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("the answer of division = ", ans)
		fmt.Println("the remainder = ", remainder)
	}

	hr()

	{
		ans, err := performAFunc(divideWithReturn, 10, 3)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ans)
	}
}

func addTo(base int, vals ...int) int {
	num := base
	for _, v := range vals {
		num += v
	}
	return num
}

func divideWithReturn(num, denom int) (quotient int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("can't divide by 0")
		return quotient, remainder, err
	}
	quotient, remainder = num/denom, num%denom
	return quotient, remainder, err
}

type aFuncType func(int, int) (int, int, error)

func performAFunc(theFunction aFuncType, a, b int) (quotient int, err error) {
	quotient, _, err = theFunction(a, b)
	return quotient, err
}

func hr() {
	fmt.Println("===break===")
}
