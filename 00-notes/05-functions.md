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
