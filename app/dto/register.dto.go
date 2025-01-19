package dto

type RegisterBody struct {
	FirstName string `json:"fname" form:"fname" validate:"required,min=2"`
	LastName  string `json:"lname" form:"lname" validate:"required,min=2"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" validate:"required,min=6"`
}
