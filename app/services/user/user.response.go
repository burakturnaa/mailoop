package _user

import (
	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Id        *primitive.ObjectID `json:"id,omitempty"`
	FirstName string              `json:"fname,omitempty"`
	LastName  string              `json:"lname,omitempty"`
	Email     string              `json:"email,omitempty"`
	Token     string              `json:"token,omitempty"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		Id:        &user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
