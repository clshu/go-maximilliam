package user

import (
	"errors"
	"fmt"
	"time"
)

// User contains user info
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// OutputUserDetails display User details
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName)
	fmt.Println(u.lastName)
	fmt.Println(u.birthdate)
	// fmt.Println(u.createdAt)
}

// ClearUserName clears first name and last name
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// New is contructor of User
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
