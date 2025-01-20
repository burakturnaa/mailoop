package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CompanyBody struct {
	Name     string `json:"name" form:"name" conform:"trim" validate:"required,min=2"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Phone    string `json:"phone" form:"phone" validate:"required,phone"`
	Location string `json:"location" form:"location" validate:"required,min=2"`
	Website  string `json:"website" form:"website" validate:"required,min=2,url"`
}

type UpdateCompanyBody struct {
	Id       primitive.ObjectID `json:"id,omitempty" form:"id"`
	Name     string             `json:"name" form:"name" validate:"required,min=2"`
	Email    string             `json:"email" form:"email" validate:"required,email"`
	Phone    string             `json:"phone" form:"phone" validate:"required,min=2"`
	Location string             `json:"location" form:"location" validate:"required,min=2"`
	Website  string             `json:"website" form:"website" validate:"required,min=2,url"`
}
