package models

import (
	"github.com/tkanos/go-dtree"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- TargetSchema models & dtos
// -----------------------------------------------

type TargetParams struct {
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
	Max  int     `json:"max"`
}

type TargetSchema struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Version      string             `json:"version"`
	Targets      []TargetParams     `json:"targets"`
	DesicionTree []dtree.Tree       `json:"desicionTree"`
}

type TargetSchemas []TargetSchema

//THIS MODEL IS USED ONLY FOR SWAGGER (HAVE TROUBLE CREATING THE DOCS WITH dtree.Tree)
type TargetSchemaSwagger struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Version      string             `json:"version"`
	Targets      []TargetParams     `json:"targets"`
	DesicionTree []Tree             `json:"desicionTree"`
}
