package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) String() string {
	return fmt.Sprintf("%s %s, age %d ", p.first, p.last, p.age)
}

func main() {
	fmt.Println("Types, Methods and Interfaces")
	hr()
	{
		type myStructType struct {
			first string
			last  string
			age   int
		}
		someone := myStructType{
			first: "ambi",
			last:  "vert",
			age:   12,
		}
		fmt.Println(someone)
	}
	hr()
	{
		me := person{
			first: "adi",
			last:  "verma",
			age:   18,
		}
		aboutMe := me.String()
		fmt.Println(aboutMe)
	}
	hr()
	{
		type Employee struct {
			name   string
			age    int
			salary int
		}
		type Manager struct {
			Employee
			Reports []Employee
		} // the manager is basically an employee with additional field.

		oneManager := Manager{
			Employee: Employee{
				name:   "amiable foe",
				age:    23,
				salary: 123,
			},
			Reports: []Employee{},
		}

		fmt.Println(oneManager.name) //directly accessible
	}
	hr()
	{
		//interfaces
		type Stringer interface {
			String() string
		}

	}
}

func hr() {
	fmt.Println("===break===")
}
