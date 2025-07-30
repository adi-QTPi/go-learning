# Conditionals
- Don't put parenthisis.
```go
    n := rand.Intn(10)
    if n < 5 {
        fmt.Println("thas too low man")
    } else {
        fmt.Println("cool")
    }
```
## Scoping a varaible to conditional
- Sometimes you want some variables to be used only for the conditional.
- So Go allows spedial kind of decleration
```go
    fmt.Println("take 2")
    if n := rand.Intn(20); n < 10 {
        fmt.Println("thas too low man", n)
    } else {
        fmt.Println("cool", n)
    }
```
