package _mailTemplate

import (
	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailTemplateResponse struct {
	Id      *primitive.ObjectID `json:"id,omitempty"`
	Title   string              `json:"title,omitempty"`
	Subject string              `json:"subject,omitempty"`
	Content string              `json:"content,omitempty"`
	Token   string              `json:"token,omitempty"`
}

func NewMailTemplateResponse(mailTemplate models.MailTemplate) MailTemplateResponse {
	return MailTemplateResponse{
		Id:      &mailTemplate.Id,
		Title:   mailTemplate.Title,
		Subject: mailTemplate.Subject,
		Content: mailTemplate.Content,
	}
}

func NewMailTemplateArrayResponse(mailTemplates []models.MailTemplate) []MailTemplateResponse {
	mailTemplateRes := []MailTemplateResponse{}
	for _, v := range mailTemplates {
		Id := v.Id
		p := MailTemplateResponse{
			Id:      &Id,
			Title:   v.Title,
			Subject: v.Subject,
			Content: v.Content,
		}
		mailTemplateRes = append(mailTemplateRes, p)
	}
	return mailTemplateRes
}
