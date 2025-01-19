package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatalln(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	databaseName := "mailoop"
	db := client.Database(databaseName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections, err := db.ListCollectionNames(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatalf("Error listing collection names: %v", err)
	}

	collectionExists := false
	for _, existingCollection := range collections {
		if existingCollection == collectionName {
			collectionExists = true
			break
		}
	}

	if !collectionExists {
		err := db.CreateCollection(ctx, collectionName)
		if err != nil {
			log.Fatalf("Error creating a collection: %v", err)
		}
		log.Printf("'%s' collection is created", collectionName)
	} else {
		log.Printf("'%s' collection already exists.", collectionName)
	}

	return db.Collection(collectionName)
}
