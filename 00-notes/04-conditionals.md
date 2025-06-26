# Conditionals

- operators(basic)
    * `==` equal to
    * `!=` not equal to
    * `<`
    * `>`
    * `>=`
    * `<=`
```
if height > 4 {
    fmt.Println("you are tall man!")
}

if height > 6 {
    fmt.Println("one")
} else if height > 4 {
     fmt.Println("okay, i see")
} else {
    fmt.Println("invisible boy")
}
```

## the other way

- Traditional
```
length := getLength(email);
if length <1 {
    fmt.Println("invalid")
}
```
- if Initial_statement; Condition{...blah blah...}
- better if we write
```
if length:= getLength(email); length <1 {
    fmt.Println("invalid email")
}
```

## advantage
- using "the other way" serves advantages
    - the `length` variable gets smaller scope **(concise Scope)**
    - the same variable is not accessible out of the if block
    - offcourse shortens the code.