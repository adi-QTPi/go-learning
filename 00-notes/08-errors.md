# Errors in Go

```
    user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	profile, err := getuserProfile(user.id)
	if err != nil {
		fmt.Println(err)
		return
	}

    func getuser() (User, error) {
        // do some logic here
        //the function returns User struct
    }
```

- Error is just an interface

## Errors Packags
```
import "errors"
```
```
var err error = errors.New("this is an error, without the cumbersome structs and all!")
```