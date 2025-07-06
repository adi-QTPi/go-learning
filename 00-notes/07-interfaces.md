# Interface
```
type shape interface {
    area() float64
    perimeter() float64
}
```
```
type rect struct{
    widht, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64{
    return 2*(r.width+r.height)
}
```
```
type circle struct{
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64{
    return 2 * math.Pi * c.radius
}
```

- both the structs rectangle and circle belong to the shape interface
- both have the same methods defined to them
- but the methods have different calculation techniques.
- both structs `implement` the shape interface.
- if a function takes an interface as argument, then while using the function u can use any struct type which follows that interface format !

```
func givearea(s shape) float64{
    return s.area()
}

onerect := rect{
    height : 10
    width : 20
}
onecircle := circle {
    radius : 12
}

givearea (onerect)
givearea (onecircle)
```
- both the givearea funciton work and utilises their respective area methods.


## Type assertion

```
type shape interface {
    area() float64
}

type circle struct {
    radius float64
}

c, ok := s.(circle)  
// Casting the interface to a struct
//s is a shape interface object
// we and to check if it was a circle struct
// `ok` is a boolean that stores true if s was a circle 
// `c` stores the interface as a circle struct if the shape was from a struct in the first place

```

## Type Switches
- to do different things for different structs which implement the `same` interface.
```
func printNumericValue (num interface{}){

    switch v := num.(type) {
    case int:
        fmt.Printf("%T\n", v)
    case string:
        fmt.Printf("%T\n", v)
    default:
        fmt.Printf("%T\n", v)
    }
}
```
- `v := num.(type)` v stores the type of num
- `Printf("%T\n", v)` prints the type v

## Clean interfaces and Good Practises
- Keep the interfaces SMALL (with less methods for universal behaviour and thus useful)
- Interfaces are `not` Classes !
- Interface dont have constructor, destructors etc. No inheritance and heirarchical
- Interfaces store the function signature, `not` the function definition.