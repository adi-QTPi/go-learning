# Assigning variables

`var empty string` the standard way 
    * can be used to declare variables with "specific" data type

The other method **Short Variable Decleration** automatically gives the data type based on what u are assigning it to. \
`:=` used to initialise and `=` used to assign (or re-assign)
```
empty := ""
num_cars := 10
temp := 0.4
var isFunny = true
```

## multiple assignment
```
mileage, company := 802, "tata"
```

## type casting
- temp_int := 88
- temp_float := float64(temp_int)

# constants
- not support short assignment
- every const should have a vlaue while compiling
```
const name_user = "aditya verma"

const first = "aditya"
const last = "verma"
const fullname = first +" "+ last 
```
- the const are **not calculated while runtime**, they are replaced with the values while compiling itself.