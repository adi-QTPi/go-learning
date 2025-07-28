package main

import (
	"fmt"
	"maps"
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
	fmt.Println(slice)
}

func printSliceWithName(name string, slice []int) {
	fmt.Println(name, " = ", slice)
}

func hr() {
	fmt.Println("===break===")
}

func main() {
	fmt.Println("this is chapter 3")
	{
		var arr [2]int //all index filled with 0 values
		for i := 0; i < 2; i++ {
			fmt.Println(arr[i])
		}
	}

	hr()

	{
		var arr2 = [3]int{1, 2, 3} // initialising list
		for i := 0; i < 3; i++ {
			fmt.Println(arr2[i])
		}
	}

	hr()

	{
		var arr3 = [9]int{2: 100, 4: 400, 7: 123} //partial initialising
		for i := 0; i < 9; i++ {
			fmt.Println(arr3[i])
		}
	}

	hr()

	{
		var arr4 = [...]int{1, 2, 3} // auto detect length
		for i := 0; i < len(arr4); i++ {
			fmt.Println(arr4[i])
		}
	}

	hr()
	hr()

	{
		var sl = []int{1, 2, 3, 8}
		printSlice(sl)
		fmt.Println(sl == nil) //gives false
	}

	hr()

	{
		var sl2 []int
		printSlice(sl2)
		fmt.Println(sl2 == nil)
	}

	hr()

	{
		var sl3 = []int{1, 2}
		// var sl4 = [...]int{1, 2} // arrays can not be compared using slices.Equal
		var sl5 = []int{2, 1}
		fmt.Println(slices.Equal(sl3, sl5)) //gives false
	}

	hr()

	{
		var sl6 = []int{9, 8, 7}
		printSlice(sl6)
		sl6 = append(sl6, 12)
		printSlice(sl6)

		hr()

		sl7 := []int{5, 6, 11, 12, 12, 12, 14, 14}
		sl6 = append(sl7, sl6...)
		printSlice(sl6)
	}

	hr()

	{
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
	}

	hr()

	{
		sl11 := []string{"a", "b", "c", "d", "e"}
		a := sl11[:] //default values 0 and the last index...
		b := sl11[2:3]
		c := sl11[:2]
		d := sl11[3:]

		fmt.Println("a = ", a)
		fmt.Println("b = ", b)
		fmt.Println("c = ", c)
		fmt.Println("d = ", d)

		sl11 = append(sl11, "x", "y")
		fmt.Println(("a = "), a)
		fmt.Println(("sl11 = "), sl11)
	}

	hr()

	{
		x := make([]string, 0, 5)
		x = append(x, "a", "b", "c", "d")

		y := x[2:4:5]
		z := x[:2:5]
		fmt.Println("x = ", x, "cap = ", cap(x))
		fmt.Println("y = ", y, "cap = ", cap(y))
		fmt.Println("z = ", z, "cap = ", cap(z))

		z = append(z, "r") //changes made in z
		fmt.Println("z = ", z, "cap = ", cap(z))
		fmt.Println("x = ", x, "cap = ", cap(x)) //changes in x are observed
	}

	hr()

	{
		x := []int{1, 2, 3, 4, 5}
		y := make([]int, 3)
		z := make([]int, 2)

		num := copy(y, x)
		copy(z, x[2:])

		fmt.Println(num) //number of elements copied
		printSliceWithName("x", x)
		printSliceWithName("y", y)
		printSliceWithName("z", z)

		copy(x[:2], x[1:]) //works on subslices of same slice also !
		printSliceWithName("x", x)
	}

	hr()

	{
		array := [3]int{1, 2, 3}
		array2slice := array[:]
		printSliceWithName("array2slice", array2slice)

		slice2array := [3]int(array2slice)
		fmt.Println("converted array = ", slice2array)
	}

	hr()

	{
		slice := []int{1, 2, 3, 4, 5, 6}
		arrayPointer := (*[6]int)(slice)
		smallArrayPointer := (*[3]int)(slice)
		fmt.Println("arrayPointer = ", *arrayPointer) //dereferencing explored !!!
		fmt.Println("smallArrayPointer = ", smallArrayPointer)
	}

	hr()
	hr()

	{
		fmt.Println("starting maps")
		var nilMap map[string]int
		fmt.Println(nilMap)
		fmt.Println(len(nilMap))

		empMap := map[string]int{}
		fmt.Println(empMap)
	}

	hr()

	{
		teams := map[string][]string{
			"CSK": {"dhoni", "ashwin", "raina"},
			"MI":  {"rohit", "pandya", "pollard"},
		}

		fmt.Println(teams)
		fmt.Println(teams["CSK"])

		teams["CSK"] = append(teams["CSK"], "jadeja")
		newTeam := []string{"verma", "badoni", "vohra"}
		teams["MI"] = newTeam

		fmt.Println(teams["MI"])
		fmt.Println(teams["CSK"])

		delete(teams, "MI")
		fmt.Println(teams)
	}

	hr()

	{
		m := map[int]string{
			1: "hello",
			2: "there",
			3: "mister",
		}

		fmt.Println(m)
		clear(m)
		fmt.Println(m)
		x := m[2]
		fmt.Println(x)
	}

	hr()

	{
		m := map[string]int{
			"hello": 2,
			"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok)
		v, ok = m["world"]
		fmt.Println(v, ok)
		v, ok = m["brick"]
		fmt.Println(v, ok)
	}

	hr()

	{
		m := map[int]string{
			1: "hi",
			2: "there",
		}
		n := map[int]string{
			1: "hi",
			2: "there",
		}
		fmt.Println(maps.Equal(m, n))
	}

	hr()

	{
		map1 := map[int]bool{}
		sl := []int{1, 2, 3, 4, 2, 2, 3, 4, 1, 21, 34, 2, 2} //6 unique values.

		//collect unique values from sl (like make a set.)

		for k, v := range sl {
			fmt.Println("the ", k, "th value of slice is ", v)
		} //traverse through slice

		for _, v := range sl {
			map1[v] = true
		}

		fmt.Println(len(map1)) //gives 6

		//to check if 4 is in the set
		if map1[4] {
			fmt.Println("yes its there !")
		}

		for k, v := range map1 {
			fmt.Println(k, " -> ", v)
		}
	}

	hr()
	hr()

	{
		fmt.Println("Structs ...")

		type person struct {
			name    string
			age     int
			address string
			isMale  bool
		}

		var aditya person
		fmt.Println(aditya)
		amit := person{}
		fmt.Println(amit)

		bob := person{
			"bob",
			30,
			"delhi",
			true,
		}
		fmt.Println(bob)

		julie := person{
			name: "julie",
			age:  23,
		}
		fmt.Println(julie)
	}

	hr()

	{
		type gamer struct {
			name    string
			age     int
			xp      int
			favGame string
		}

		chintuGamer := gamer{
			age:     12,
			favGame: "minecraft",
		}
		chintuGamer.name = "chinteshwar"
		chintuGamer.xp = 234

		fmt.Println(chintuGamer)

		chintuGamer.age++

		chintuGamer.favGame = chintuGamer.favGame[4:]

		fmt.Println(chintuGamer)
	}

	hr()

	{
		johnny := struct {
			name     string
			age      int
			phoneNum int
		}{
			name:     "johnathon",
			age:      12,
			phoneNum: 12342231,
		}

		fmt.Println(johnny)
	}

	hr()

	{
		type manOne struct {
			name string
			age  int
		}
		type manTwo struct {
			name string
			age  int
		}
		piyush := manOne{
			name: "Piyush Garg",
			age:  45,
		}
		p_clone := manTwo(piyush)
		fmt.Println(p_clone)
	}
}
