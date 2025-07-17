package main

import "fmt"

func main() {
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
}

func getuser() (User, error) {
	// do some logic here
	//the function returns User struct
}
