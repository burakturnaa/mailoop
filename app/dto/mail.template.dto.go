package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type MailTemplateBody struct {
	Title   string `json:"title" form:"title" validate:"required,min=2"`
	Subject string `json:"subject" form:"subject" validate:"required,min=2"`
	Content string `json:"content" form:"content" validate:"required,min=2"`
}

type UpdateMailTemplateBody struct {
	Id      primitive.ObjectID `json:"id,omitempty" form:"id"`
	Title   string             `json:"title" form:"title" validate:"required,min=2"`
	Subject string             `json:"subject" form:"subject" validate:"required,min=2"`
	Content string             `json:"content" form:"content" validate:"required,min=2"`
}
