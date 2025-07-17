# For loop
```
for i := 0; i<10 ; i++{
    fmt.Printf("%v", i)
}
```
- can skip declerations also, just put semicolons but without anything

## no While loop in Go !!!

```
for CONDITION {
    //do some stuff until condition is true
}
```
```
i := 0
for i<10 {
    fmt.Printf("%v", i)
    i++
}
```

- for can run **infinite** loop if we give it nothing!
```
for {
    //do something for infinite times!
}
```