# Structs

```
type car struct {
    make string
    model string
    height int
    width int
}
```
- access the values using `<name of car variable>.make` etc.

- Struct within structs
```
type Wheel struct {
    radius int
    material string
}
type car struct {
    make string
    model string
    front_wheel Wheel
    back_wheel Wheel
}
```
## struct instantiation
```
myCar := car{}

myCar.front_wheel.radius = 4
```

## Anonymous struct
`avoid` anonymous structs , but still
- like a struct but immediately instantiating a struct directly.
- for structs that would have only one instances. (other wise ignor)
```
myCar := struct {
    make string
    model string
    //struct defined 
}{
    make : "honda"
    model : "city"
    //struct instantiated
}
```

- or this can be used to avoid nested structs
```
type car struct {
    make string
    model string

    wheel struct {
        radius int
        material string
    }
}
```

## Embedded Struct
- like children struct (like in cpp)
- making a struct truck that has car embedded in it.
- classes and inheritance are not a part of Go
- go is `not` an object oriented language.

```
type car struct {
    make string
    model string
}
type truck struct {
    car
    bedsize int
}
```
- different from nested because you access the fields.
```
onetruck := truck {
    bedsize : 10,
    car : car{
        make: "mecedes"
        model : "ek truck"
    }
}

fmt.Println(myTruck.model)
//not myTruck.car.model ...
```

## Struct Methods
- say you have a rectangle struct
```
type rect struct {
    height int
    widht int
}
```

- now we want a method (some function that takes the obejct of that struct type only and does some struct specific work)

```
func (r rect) area() int{
    return r.height * r.width
}
// method for rect struct defined
```
```
onerect := {
    width : 5,
    height: 10,
}

fmt.Println(onerect.area())
```