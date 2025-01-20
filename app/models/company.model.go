package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id        primitive.ObjectID `bson:"id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Phone     string             `bson:"phone,omitempty"`
	Location  string             `bson:"location,omitempty"`
	Website   string             `bson:"website,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
}
