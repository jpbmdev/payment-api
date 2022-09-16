package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- Loan models & dtos
// -----------------------------------------------

type AddPaymentToLoanDto struct {
	Date   string  `json:"date" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0" `
}

type CreateLoanDto struct {
	Amount    float64            `json:"amount" binding:"required,gte=1"`
	Term      int                `json:"term" binding:"required,gte=1"`
	UserId    primitive.ObjectID `json:"userId" binding:"required"`
	StartDate string             `json:"StartDate" binding:"required"`
}

type LoanHistory struct {
	MonthStart  time.Time          `json:"monthStart"`
	MonthEnd    time.Time          `json:"monthEnd"`
	Accumulated float64            `json:"accumulated"`
	MonthDebt   float64            `json:"monthDebt"`
	PaymentId   primitive.ObjectID `json:"paymentId,omitempty" bson:"paymentId,omitempty"`
}

type Loan struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Amount         float64            `json:"amount"`
	Term           int                `json:"term"`
	Rate           float64            `json:"rate"`
	UserId         primitive.ObjectID `json:"userId" bson:"userId"`
	TargetSchemaId primitive.ObjectID `json:"targetSchemaId" bson:"targetSchemaId"`
	TargetName     string             `json:"targetName" bson:"targetName"`
	StartDate      time.Time          `json:"StartDate" bson:"startDate"`
	EndDate        time.Time          `json:"endDate" bson:"endDate"`
	Quota          float64            `json:"quota"`
	Debt           float64            `json:"debt"`
	LoanHistory    []LoanHistory      `json:"loanHistory" bson:"loanHistory"`
}

type Loans []Loan
