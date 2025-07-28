# Slice - Array Conversions

## Array to Slice
- Even arrays can be sliced using the slice expression
```go
    array := [3]int{1, 2, 3}
    array2slice := array[:]
```
- Same memory sharing as was in slices, thus memory overwrite threat remains...

## Slice to Array
- Slices can be turned into array using type casting.
- New copy of the array is made somewhere else.
- So no memory overlap.
```go
    slice2array := [3]int(array2slice)
	fmt.Println("converted array = ", slice2array)
```
- Here the initialisation of array plays key role.
- the size of array here cant be defined as `[...]`. (compile time error)
- Size of array can be smaller than that of slice, but NEVER larger. (Panic : Runtime error).

## Slice to ArrayPointer
- Slice can be used to make an array pointer.
- This one will have the memory shared so beware.
```go
    slice := []int{1, 2, 3, 4, 5, 6}
    arrayPointer := (*[6]int)(slice)
    smallArrayPointer := (*[3]int)(slice)
    fmt.Println("arrayPointer = ", *arrayPointer) //dereferencing explored !!!
    fmt.Println("smallArrayPointer = ", smallArrayPointer)
```