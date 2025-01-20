package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mailTemplateRepository struct {
	MailTemplateCollection *mongo.Collection
}

type MailTemplateRepository interface {
	GetAll() ([]models.MailTemplate, error)
	GetOne(id primitive.ObjectID) (models.MailTemplate, error)
	InsertMailTemplate(mailTemplate models.MailTemplate) (models.MailTemplate, error)
	UpdateMailTemplate(mailTemplate models.MailTemplate) (models.MailTemplate, error)
	DeleteMailTemplate(id primitive.ObjectID) (bool, error)
	FindByMailTemplateID(mailTemplateId primitive.ObjectID) (models.MailTemplate, error)
}

func NewMailTemplateRepository(dbClient *mongo.Collection) MailTemplateRepository {
	return &mailTemplateRepository{MailTemplateCollection: dbClient}
}

func (mt *mailTemplateRepository) GetAll() ([]models.MailTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var mailTemplates []models.MailTemplate
	cursor, err := mt.MailTemplateCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &mailTemplates); err != nil {
		return nil, err
	}
	return mailTemplates, nil
}

func (mt *mailTemplateRepository) GetOne(id primitive.ObjectID) (models.MailTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var mailTemplates models.MailTemplate
	err := mt.MailTemplateCollection.FindOne(ctx, bson.M{"id": id}).Decode(&mailTemplates)
	fmt.Println(err)
	if err != nil {
		return mailTemplates, err
	}
	return mailTemplates, nil
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

func (mt *mailTemplateRepository) UpdateMailTemplate(mailTemplate models.MailTemplate) (models.MailTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", mailTemplate.Id}}
	update := bson.D{{"$set", bson.D{{"title", mailTemplate.Title}, {"subject", mailTemplate.Subject}, {"content", mailTemplate.Content}, {"updated_at", time.Now()}}}}
	_, err := mt.MailTemplateCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return mailTemplate, err
	}
	return mailTemplate, nil
}

func (mt *mailTemplateRepository) DeleteMailTemplate(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := mt.MailTemplateCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || result.DeletedCount < 1 {
		return false, err
	}
	return true, nil
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
