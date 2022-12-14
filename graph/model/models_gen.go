// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteUser struct {
	ID string `json:"id"`
}

type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
}

type UpdateUser struct {
	ID        string  `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Dob       *string `json:"dob"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
}
