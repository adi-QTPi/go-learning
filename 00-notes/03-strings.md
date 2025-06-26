# Strings 

- `fmt.Printf` ->  prints a formatted string to standard out
- `fmt.Sprintf() -> returns the formatted string

## %v
- used when u are not sure what else to use
- prints the Go syntax presentation of the value
- But better if you use type specific variant.
```
fmt.Printf("i am %v years old", 10)
//i am 10 years old

fmt.Printf("i am %v years old", "many")
//i am many years old
```

## %s
- interpolate a string value
```
fmt.Printf("i am %s years old", "many")
//i am many years old
```

## %d
- interpolate integer to its decimal representation.

## %f
- interpolate float
- also possible to round of the value to appt spaces.
```
fmt.Printf("this is %f years old", 10.523)
//this is 10.523 years old

//rounding of to the required decimal place
fmt.Printf("this is %.2f years old", 10.523)
// this is 10.5 years old
```

## Use
- `Sprintf` is used because we want the returned value to be stored in msg.
- later `Println` is used to print whole line.
```
const name = "hola"
const rate = 1.2344
msg := fmt.Sprintf(
    "hi %s, your open rate is %.1f dollars",
    name,
    rate
)

fmt.Println(msg)
```