package main

import "fmt"

// user pointers when you need the ff:
// when we need to update state
// want to optimize the memory for large objects taht are getting called A LOT.

type User struct {
	email    string
	username string
	age      int
}
/////////////////
// this two function has the same functino

//x amouts of bytes => sizeOf(user)
func (u User) Email() string {
	return u.email
}

// 8 bytes
// func Email(user *User) string{
// 	return user.email
// }

/////////////////

func (u *User) updateEmail(newemail string){
	u.email = newemail
}

// func updateEmail(u *User,email string){
// 	u.email =email
// }

/////////////////
func main() {
	user := User{
		email: "agg@email.com",
	}

	fmt.Println(user.Email())


	user.updateEmail("new@gmail.com")
	fmt.Println(user.Email())

}

// credits
///src  = https://www.youtube.com/watch?v=3WsEDZRif6U