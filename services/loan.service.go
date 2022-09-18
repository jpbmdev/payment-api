package services

import (
	"math"
	"time"

	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- Loan service
// -----------------------------------------------
type LoanService interface {
	CreateLoan(loan models.Loan) error
	CalculateQuota(term float64, rate float64, amount float64) float64
	FindLastYearLoans(userId primitive.ObjectID, loanStartDate time.Time) (models.Loans, error)
	FindLoansByDate(fromDate time.Time, toDate time.Time, pageSize int, page int) (models.Loans, error)
	FindLoanById(id primitive.ObjectID, loan *models.Loan) error
	UpdateLoanPayment(id primitive.ObjectID, debt float64, loanHistory []models.LoanHistory) error
	FindLoansWithDebt(target string, date time.Time) (models.Loans, error)
	FindLoansByUserId(userId primitive.ObjectID) (models.Loans, error)
}

type loanService struct {
	respository repositories.LoanRepository
}

//Function to crete new loan service
func NewLoanService() LoanService {
	return &loanService{
		respository: repositories.NewLoanRepository(),
	}
}

func (s *loanService) CreateLoan(loan models.Loan) error {
	//Create loan in database
	err := s.respository.InsertOne(loan)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

func (s *loanService) UpdateLoanPayment(id primitive.ObjectID, debt float64, loanHistory []models.LoanHistory) error {
	//Create query
	filter := bson.M{"_id": id}

	//Define data to update
	updateData := bson.M{"$set": bson.M{"debt": debt, "loanHistory": loanHistory}}

	//Update loan in database
	err := s.respository.UpdateOne(filter, updateData)

	//Handle errors
	if err != nil {
		return err
	}
	return nil
}

//Function to calculate the month quota
func (s *loanService) CalculateQuota(term float64, rate float64, amount float64) float64 {
	r := rate / 12
	quota := (r + (r / (math.Pow(1+r, term) - 1))) * amount
	//Round to two decimals
	return math.Round(quota*100) / 100
}

func (s *loanService) FindLoanById(id primitive.ObjectID, loan *models.Loan) error {
	//Create Sort query
	filter := bson.M{"_id": id}

	err := s.respository.FindOne(filter, loan)

	//Handle Errors
	if err != nil {
		return err
	}
	return nil
}

func (s *loanService) FindLoansByUserId(userId primitive.ObjectID) (models.Loans, error) {
	//Create query to find the loans a user started the year before this new loan
	filter := bson.M{
		"userId": userId,
	}

	loans, err := s.respository.Find(filter)

	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (s *loanService) FindLastYearLoans(userId primitive.ObjectID, loanStartDate time.Time) (models.Loans, error) {
	//Create query to find the loans a user started the year before this new loan
	filter := bson.M{
		"userId": userId,
		"$expr": bson.M{
			"$eq": []interface{}{
				bson.M{"$year": "$startDate"},
				loanStartDate.Year() - 1,
			},
		}}

	loans, err := s.respository.Find(filter)

	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (s *loanService) FindLoansWithDebt(target string, date time.Time) (models.Loans, error) {

	andQuery := []bson.M{}
	andQuery = append(
		andQuery,
		bson.M{"debt": bson.M{
			"$ne": 0,
		}},
	)

	if !date.IsZero() {
		andQuery = append(
			andQuery,
			bson.M{"startDate": bson.M{
				"$lte": date,
			}},
		)
	}

	if target != "" {
		andQuery = append(
			andQuery,
			bson.M{"targetName": target},
		)
	}

	filter := bson.M{"$and": andQuery}

	loans, err := s.respository.Find(filter)

	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (s *loanService) FindLoansByDate(fromDate time.Time, toDate time.Time, pageSize int, page int) (models.Loans, error) {
	//Create query to find the loans started in a range
	filter := bson.M{}

	if !fromDate.IsZero() {
		filter["startDate"] = bson.M{
			"$gte": fromDate,
		}
	}

	if !toDate.IsZero() {
		filter["startDate"] = bson.M{
			"$lte": toDate,
		}
	}

	loans, err := s.respository.FindPaginate(filter, pageSize, page)

	if err != nil {
		return nil, err
	}
	return loans, nil
}
