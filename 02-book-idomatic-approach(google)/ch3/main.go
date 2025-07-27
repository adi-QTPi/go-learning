package main

import (
	"fmt"
	"slices"
)

// THIS FUNCTION WONT EVEN WORK !!! BECAUSE [10]int and [3]int are considered as different "types" and thus wont pass as arguments to the same function...
//
//	func printArr(arr int, length int) {
//		for i := 0; i < length; i++ {
//			fmt.Println(arr[i])
//		}
//	}

func printSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func hr() {
	fmt.Println("===break===")
}

func main() {
	fmt.Println("this is chapter 3")
	var arr [2]int //all index filled with 0 values
	for i := 0; i < 2; i++ {
		fmt.Println(arr[i])
	}

	hr()

	var arr2 = [3]int{1, 2, 3} // initialising list
	for i := 0; i < 3; i++ {
		fmt.Println(arr2[i])
	}

	hr()

	var arr3 = [9]int{2: 100, 4: 400, 7: 123} //partial initialising
	for i := 0; i < 9; i++ {
		fmt.Println(arr3[i])
	}

	hr()

	var arr4 = [...]int{1, 2, 3} // auto detect length
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	hr()
	hr()

	var sl = []int{1, 2, 3, 8}
	printSlice(sl)
	fmt.Println(sl == nil) //gives false

	hr()

	var sl2 []int
	printSlice(sl2)
	fmt.Println(sl2 == nil)

	hr()

	var sl3 = []int{1, 2}
	// var sl4 = [...]int{1, 2} // arrays can not be compared using slices.Equal
	var sl5 = []int{2, 1}
	fmt.Println(slices.Equal(sl3, sl5)) //gives false

	hr()

	var sl6 = []int{9, 8, 7}
	printSlice(sl6)
	sl6 = append(sl6, 12)
	printSlice(sl6)

	hr()

	sl7 := []int{5, 6, 11, 12, 12, 12, 14, 14}
	sl6 = append(sl7, sl6...)
	printSlice(sl6)

	hr()

	sl8 := make([]int, 4) //here the capacity set to default managed by Go runtime
	// sl8 := make([]int, 4, 10)
	fmt.Println(sl8, "length = ", len(sl8), "capacity = ", cap(sl8))
	sl8 = append(sl8, 2, 3, 5)
	fmt.Println(sl8, "length = ", len(sl8), "capacity = ", cap(sl8))

	hr()

	clear(sl8)
	fmt.Println(sl8)

	hr()
	var sl9 []int
	fmt.Println(sl9, "length = ", len(sl9), "capacity = ", cap(sl9), "is it null -> ", sl9 == nil)

	var sl10 = []int{}
	fmt.Println(sl10, "length = ", len(sl10), "capacity = ", cap(sl10), "is it null -> ", sl10 == nil)

	hr()

	sl11 := []string{"a", "b", "c", "d", "e"}
	a := sl11[:] //default values 0 and the last index...
	b := sl11[2:3]
	c := sl11[:2]
	d := sl11[3:]

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
}
