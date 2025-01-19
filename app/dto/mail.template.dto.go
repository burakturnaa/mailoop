package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type MailTemplateBody struct {
	Title   string `json:"title" form:"title" validate:"required,min=2,trim"`
	Subject string `json:"subject" form:"subject" validate:"required,min=2,trim"`
	Content string `json:"content" form:"content" validate:"required,min=2,trim"`
}

type UpdateMailTemplateBody struct {
	Id      primitive.ObjectID `json:"id" form:"id" validate:"required"`
	Title   string             `json:"title" form:"title" validate:"required,min=2,trim"`
	Subject string             `json:"subject" form:"subject" validate:"required,min=2,trim"`
	Content string             `json:"content" form:"content" validate:"required,min=2,trim"`
}
