package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	LoanId primitive.ObjectID `json:"loanId,omitempty" bson:"loanId,omitempty"`
	Date   time.Time          `json:"date"`
	Amount float64            `json:"amount"`
}
