package repositories

import (
	"context"
	"fmt"

	"github.com/jpbmdev/payment-api/database"
	"github.com/jpbmdev/payment-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// -----------------------------------------------
// -- Loan Repository
// -----------------------------------------------
type LoanRepository interface {
	InsertOne(loan models.Loan) error
	UpdateOne(filter bson.M, updateData bson.M) error
	FindOne(filter bson.M, loan *models.Loan) error
	Find(filter bson.M) (models.Loans, error)
	FindPaginate(filter bson.M, pageSize int, page int) (models.Loans, error)
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

	//Insert loan into the database
	_, err = r.collection.InsertOne(r.ctx, loan)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (r *loanRepository) UpdateOne(filter bson.M, updateData bson.M) error {
	var err error

	//Insert user into the database
	_, err = r.collection.UpdateOne(r.ctx, filter, updateData)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (r *loanRepository) FindOne(filter bson.M, loan *models.Loan) error {

	//Get the loan
	err := r.collection.FindOne(r.ctx, filter).Decode(loan)

	//Handle Errors
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

	//Iterate over the cursor to get the loan
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

func (r *loanRepository) FindPaginate(filter bson.M, pageSize int, page int) (models.Loans, error) {
	loans := []models.Loan{}

	fmt.Println(page)
	fmt.Println(pageSize)

	//Add pagination to query
	opts := options.Find().SetLimit(int64(pageSize)).SetSkip(int64((page - 1) * pageSize))

	//Get cursor of database
	cur, err := r.collection.Find(r.ctx, filter, opts)

	//Handle errors
	if err != nil {
		return nil, err
	}

	//Iterate over the cursor to get the loans
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
