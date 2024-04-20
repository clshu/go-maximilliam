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

// Admin is the type for administrator
type Admin struct {
	email    string
	password string
	User     // Anonymous embedding
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

// New is the onstructor of User
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

// NewAdmin is the constructor of Admin
func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "01/01/1980",
			createdAt: time.Now(),
		},
	}
}
