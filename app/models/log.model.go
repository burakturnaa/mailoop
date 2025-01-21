package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	Id             primitive.ObjectID `bson:"id,omitempty"`
	MailTemplateId primitive.ObjectID `bson:"mail_template_id,omitempty"`
	EmailList      []string           `bson:"email_list,omitempty"`
	CreatedAt      time.Time          `bson:"created_at,omitempty"`
}
