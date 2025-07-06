# Functions 

## function signature
- tells what kind of args expected.
- tells what data type would be returned.

## function definition

```
func sub ( x int, y int) int {
    return x-y
}
```

```
func FUNCTION_NAME ( ARG1 <type of arg1>, ARG2 <type of arg2>, ...) <type of what the function returns> {
    //function content
    return SOMETHING
}
```
*type comes AFTER the argument!*  similar to how we speak in english...

if multiple arguments have the same type, then can use the following method to define the function.
```
func add (x,y int) int{
    return x+y
}
```

## Return properties
```
func getCoord()(x,y int){
    //initialises x and y with their 0 values

    return //returns both x and y (less recommended to use naked return)
}

//same as 
 func getCoord(){
    var x int
    var y int
    return x,y
 }

//Explicit and impicit retur
func hello() (x,y int){
    return 5,6 //this overides returning x and y...
}
```

## Early Return
- you can return early the variables if some condition breaks...

- following show the implementation of `guard clauses`, also known as early return.
```
finc divide (dividend, divisor int)(int, error){
    if divisor == 0{
        return 0, errors.New("dont divide by 0)

        //returns 0 as int and the error object 
    }

    return dividend/divisor, nil
}
```