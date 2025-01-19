package repository

import (
	"context"
	"log"
	"time"

	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	UserCollection *mongo.Collection
}

type UserRepository interface {
	InsertUser(user models.User) (models.User, error)
	FindByUserID(userId primitive.ObjectID) (models.User, error)
	FindByUserEmail(email string) (models.User, error)
}

func NewUserRepository(dbClient *mongo.Collection) UserRepository {
	return &userRepository{UserCollection: dbClient}
}

func (u *userRepository) InsertUser(user models.User) (models.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	_, err := u.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) FindByUserID(userId primitive.ObjectID) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := u.UserCollection.FindOne(ctx, bson.M{"id": userId}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) FindByUserEmail(userEmail string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	err := u.UserCollection.FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		// panic("Failed to hash a password")
	}
	return string(hash)
}
