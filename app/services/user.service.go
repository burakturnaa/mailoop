package services

import (
	"errors"

	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/repository"
	_user "github.com/burakturnaa/mailoop.git/app/services/user"

	"github.com/mashingan/smapping"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	userRepo repository.UserRepository
}

type UserService interface {
	CreateUser(registerRequest dto.RegisterBody) (*_user.UserResponse, error)
	FindUserByID(userId primitive.ObjectID) (*_user.UserResponse, error)
	FindUserByEmail(userEmail string) (*_user.UserResponse, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateUser(registerRequest dto.RegisterBody) (*_user.UserResponse, error) {
	user, err := u.userRepo.FindByUserEmail(registerRequest.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))
	if err != nil {
		return nil, err
	}

	user, _ = u.userRepo.InsertUser(user)

	res := _user.NewUserResponse(user)
	return &res, nil
}

func (u *userService) FindUserByID(userId primitive.ObjectID) (*_user.UserResponse, error) {
	user, err := u.userRepo.FindByUserID(userId)
	if err != nil {
		return nil, err
	}

	userResponse := _user.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

func (u *userService) FindUserByEmail(email string) (*_user.UserResponse, error) {
	user, err := u.userRepo.FindByUserEmail(email)
	if err != nil {
		return nil, err
	}

	userResponse := _user.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}
