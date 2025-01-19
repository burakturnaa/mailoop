package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailTemplate struct {
	Id        primitive.ObjectID `bson:"id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Subject   string             `bson:"subject,omitempty"`
	Content   string             `bson:"content,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
}
