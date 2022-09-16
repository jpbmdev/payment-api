package repositories

import (
	"context"

	"github.com/jpbmdev/payment-api/database"
	"github.com/jpbmdev/payment-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// -----------------------------------------------
// -- Payment Repository
// -----------------------------------------------
type PaymentRepository interface {
	InsertOne(payment models.Payment) error
	Find(filter bson.M) (models.Payments, error)
}

type paymentRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

//Function to crete new payment repository
func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{
		collection: database.GetCollection("payments"),
		ctx:        context.Background(),
	}
}

func (r *paymentRepository) InsertOne(payment models.Payment) error {
	var err error

	//Insert payment into the database
	_, err = r.collection.InsertOne(r.ctx, payment)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) Find(filter bson.M) (models.Payments, error) {
	payments := []models.Payment{}

	//Get cursor of database
	cur, err := r.collection.Find(r.ctx, filter)

	//Handle errors
	if err != nil {
		return nil, err
	}

	//Iterate over the cursor to get the loan
	for cur.Next(r.ctx) {
		var payment models.Payment
		err = cur.Decode(&payment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}
