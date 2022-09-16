package repositories

import (
	"context"

	"github.com/jpbmdev/payment-api/database"
	"github.com/jpbmdev/payment-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// -----------------------------------------------
// -- Loan Repository
// -----------------------------------------------
type LoanRepository interface {
	InsertOne(loan models.Loan) error
	Find(filter bson.M) (models.Loans, error)
}

type loanRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

//Function to crete new loan repository
func NewLoanRepository() LoanRepository {
	return &loanRepository{
		collection: database.GetCollection("loans"),
		ctx:        context.Background(),
	}
}

func (r *loanRepository) InsertOne(loan models.Loan) error {
	var err error

	//Insert user into the database
	_, err = r.collection.InsertOne(r.ctx, loan)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (r *loanRepository) Find(filter bson.M) (models.Loans, error) {
	loans := []models.Loan{}

	//Get cursor of database
	cur, err := r.collection.Find(r.ctx, filter)

	//Handle errors
	if err != nil {
		return nil, err
	}

	//Iterate over the cursor to get the targetSchemas
	for cur.Next(r.ctx) {
		var loan models.Loan
		err = cur.Decode(&loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, loan)
	}

	return loans, nil
}
