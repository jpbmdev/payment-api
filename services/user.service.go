package services

import (
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

// -----------------------------------------------
// -- User service
// -----------------------------------------------
type UserService interface {
	CreateUser(user models.User) error
	GetUsers() (models.Users, error)
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
