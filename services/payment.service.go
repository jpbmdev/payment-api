package services

import (
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/repositories"
)

// -----------------------------------------------
// -- Payment service
// -----------------------------------------------
type PaymentService interface {
	CreatePayment(payment models.Payment) error
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
