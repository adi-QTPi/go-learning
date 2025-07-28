# Structs

- Go is Not an object oriented language, because it does not exhibit inheritance.
- That is the biggest difference between Go and any other object oriented language like C++.
- Though there are ways here to implement the OOPS things.
<br><br>
- Structs declared like -> `type structName struct` 
```go
    type person struct {
        name    string
        age     int
        address int
        isMale  bool
    }
```
- Struct defined (done)
<br><br>
- Defining variable of type that struct
```go
    var aditya person //0 value struct

    amit := person{} //empty struct literal
```
- unlike slices and maps, there is `no` difference between the 0 value and empty struct literal.

- Initialising struct type variables -> 2 ways
```go
    //struct type : Comma seperated values
    bob := person{
        "bob",
        30,
        "delhi",
        true,
    } //partial filling possible if put the respective 0 value of the required data type...

    //map type : key value pair! (partial filling possible)
    julie := person{
        name: "julie",
        age:  23,
    }
```

## Accessing Fields in a Struct
- fields and methods are accessed using the `dot`
- assign, change values using the dot itself.
```go
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
```

## Anonymous Structs
- one time use structs, which are not intended to be used regularly.
- effective in decoding and storing external data (like json etc.)
```go
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
```

## Comparing and Converting Structs
- a struct if has all fields comparable then it is `comparable`.
- if a field is a map or slice, then struct not comparable.
- Go disallows comparisons of different fields.
<br><br>
- Go `allows type conversion` from one struct to another iff they have fields of `same fields, names and order`.
- The only case where conversion happens :-
```go
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
```
- The `order` in which fields written also matters. (would not allow to convert).
- Different field names are an obvious indication of `non-convertibility`.