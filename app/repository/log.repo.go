package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type logRepository struct {
	LogCollection *mongo.Collection
}

type LogRepository interface {
	GetAll() ([]models.Log, error)
	GetOne(id primitive.ObjectID) (models.Log, error)
	InsertLog(log models.Log) (models.Log, error)
}

func NewLogRepository(dbClient *mongo.Collection) LogRepository {
	return &logRepository{LogCollection: dbClient}
}

func (mt *logRepository) GetAll() ([]models.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var logs []models.Log
	cursor, err := mt.LogCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}

func (mt *logRepository) GetOne(id primitive.ObjectID) (models.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var log models.Log
	err := mt.LogCollection.FindOne(ctx, bson.M{"id": id}).Decode(&log)
	fmt.Println(err)
	if err != nil {
		return log, err
	}
	return log, nil
}

func (mt *logRepository) InsertLog(log models.Log) (models.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Id = primitive.NewObjectID()
	log.CreatedAt = time.Now()
	_, err := mt.LogCollection.InsertOne(ctx, log)
	if err != nil {
		return log, err
	}

	return log, nil
}
