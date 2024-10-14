package models

type User struct {
	FirstName  string `json:"firstName" validate:"required,alpha"`
	SecondName string `json:"secondName" validate:"required,alpha"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password"`
}

// we need to be able to store a user
// Other functions we need to do things with the user we will determine later
