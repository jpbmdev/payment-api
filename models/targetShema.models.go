package models

import (
	"github.com/tkanos/go-dtree"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- TargetSchema models & dtos
// -----------------------------------------------

type TargetParams struct {
	Rate float64 `json:"rate"`
	Max  int     `json:"max"`
}

type TargetSchema struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Version      string             `json:"version"`
	New          TargetParams       `json:"new"`
	Frequent     TargetParams       `json:"frequent"`
	Premium      TargetParams       `json:"premium"`
	DesicionTree []dtree.Tree       `json:"desicionTree"`
}

type TargetSchemas []TargetSchema

//THIS MODEL IS USED ONLY FOR SWAGGER (HAVE TROUBLE CREATING THE DOCS WITH dtree.Tree)
type TargetSchemaSwagger struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Version      string             `json:"version"`
	New          TargetParams       `json:"new"`
	Frequent     TargetParams       `json:"frequent"`
	Premium      TargetParams       `json:"premium"`
	DesicionTree []Tree             `json:"desicionTree"`
}
