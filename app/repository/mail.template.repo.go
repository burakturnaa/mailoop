package repository

import (
	"context"
	"time"

	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mailTemplateRepository struct {
	MailTemplateCollection *mongo.Collection
}

type MailTemplateRepository interface {
	InsertMailTemplate(mailTemplate models.MailTemplate) (models.MailTemplate, error)
	FindByMailTemplateID(mailTemplateId primitive.ObjectID) (models.MailTemplate, error)
}

func NewMailTemplateRepository(dbClient *mongo.Collection) MailTemplateRepository {
	return &mailTemplateRepository{MailTemplateCollection: dbClient}
}

func (mt *mailTemplateRepository) InsertMailTemplate(mailTemplate models.MailTemplate) (models.MailTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mailTemplate.Id = primitive.NewObjectID()
	mailTemplate.CreatedAt = time.Now()
	_, err := mt.MailTemplateCollection.InsertOne(ctx, mailTemplate)
	if err != nil {
		return mailTemplate, err
	}

	return mailTemplate, nil
}

func (mt *mailTemplateRepository) FindByMailTemplateID(mailTemplateId primitive.ObjectID) (models.MailTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var mailTemplate models.MailTemplate
	err := mt.MailTemplateCollection.FindOne(ctx, bson.M{"id": mailTemplateId}).Decode(&mailTemplate)
	if err != nil {
		return mailTemplate, err
	}
	return mailTemplate, nil
}
