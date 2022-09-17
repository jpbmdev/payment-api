package services

import (
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- User service
// -----------------------------------------------
type UserService interface {
	CreateUser(user models.User) error
	GetUsers() (models.Users, error)
	FindUserById(id primitive.ObjectID, user *models.User) error
}

type userService struct {
	respository repositories.UserRepository
}

//Function to crete new user service
func NewUserService() UserService {
	return &userService{
		respository: repositories.NewUserRepository(),
	}
}

func (s *userService) CreateUser(user models.User) error {
	//Create user in database
	err := s.respository.InsertOne(user)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) GetUsers() (models.Users, error) {
	//Create query
	filter := bson.M{}

	users, err := s.respository.Find(filter)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) FindUserById(id primitive.ObjectID, user *models.User) error {
	//Create query
	filter := bson.M{"_id": id}

	err := s.respository.FindOne(filter, user)

	//Handle Errors
	if err != nil {
		return err
	}
	return nil
}
