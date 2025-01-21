package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type MailSenderBody struct {
	Id             primitive.ObjectID `json:"id,omitempty" form:"id"`
	MailTemplateId primitive.ObjectID `json:"mail_template_id" form:"mail_template_id" validate:"required"`
	EmailList      []string           `json:"email_list" form:"email_list" validate:"required,dive,email"`
}
