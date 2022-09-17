package services

import (
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- Payment service
// -----------------------------------------------
type PaymentService interface {
	CreatePayment(payment models.Payment) error
	GetPaymentsByLoanId(loanId primitive.ObjectID) (models.Payments, error)
}

type paymentService struct {
	respository repositories.PaymentRepository
}

//Function to crete new payment service
func NewPaymentService() PaymentService {
	return &paymentService{
		respository: repositories.NewPaymentRepository(),
	}
}

func (s *paymentService) CreatePayment(payment models.Payment) error {
	//Create payment in database
	err := s.respository.InsertOne(payment)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (s *paymentService) GetPaymentsByLoanId(loanId primitive.ObjectID) (models.Payments, error) {
	//Create query to find all payments of a loan
	filter := bson.M{
		"loanId": loanId,
	}

	users, err := s.respository.Find(filter)

	if err != nil {
		return nil, err
	}
	return users, nil
}
