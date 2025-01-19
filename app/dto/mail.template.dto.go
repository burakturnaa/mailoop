package dto

type MailTemplateBody struct {
	Title   string `json:"title" form:"title" validate:"required,min=2,trim"`
	Subject string `json:"subject" form:"subject" validate:"required,min=2,trim"`
	Content string `json:"content" form:"content" validate:"required,min=2,trim"`
}
