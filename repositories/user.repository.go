package repositories

import (
	"context"

	"github.com/jpbmdev/payment-api/database"
	"github.com/jpbmdev/payment-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// -----------------------------------------------
// -- User Repository
// -----------------------------------------------
type UserRepository interface {
	InsertOne(user models.User) error
	Find(filter bson.M) (models.Users, error)
	FindOne(filter bson.M, user *models.User) error
}

type userRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

//Function to crete new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		collection: database.GetCollection("users"),
		ctx:        context.Background(),
	}
}

func (r *userRepository) InsertOne(user models.User) error {
	var err error

	//Insert user into the database
	_, err = r.collection.InsertOne(r.ctx, user)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Find(filter bson.M) (models.Users, error) {
	users := []models.User{}

	//Get cursor of database
	cur, err := r.collection.Find(r.ctx, filter)

	//Handle errors
	if err != nil {
		return nil, err
	}

	//Iterate over the cursor to get the users
	for cur.Next(r.ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) FindOne(filter bson.M, user *models.User) error {

	//Get the targetSchema
	err := r.collection.FindOne(r.ctx, filter).Decode(user)

	//Handle Errors
	if err != nil {
		return err
	}
	return nil
}
