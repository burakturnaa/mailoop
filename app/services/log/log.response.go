package _log

import (
	"time"

	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogResponse struct {
	Id             *primitive.ObjectID `json:"id,omitempty"`
	MailTemplateId primitive.ObjectID  `json:"mail_template_id,omitempty"`
	EmailList      []string            `json:"email_list,omitempty"`
	CreatedAt      time.Time           `json:"created_at,omitempty"`
}

func NewLogResponse(log models.Log) LogResponse {
	return LogResponse{
		Id:             &log.Id,
		MailTemplateId: log.MailTemplateId,
		EmailList:      log.EmailList,
		CreatedAt:      log.CreatedAt,
	}
}

func NewLogArrayResponse(logs []models.Log) []LogResponse {
	logRes := []LogResponse{}
	for _, v := range logs {
		p := LogResponse{
			Id:             &v.Id,
			MailTemplateId: v.MailTemplateId,
			EmailList:      v.EmailList,
			CreatedAt:      v.CreatedAt,
		}
		logRes = append(logRes, p)
	}
	return logRes
}
