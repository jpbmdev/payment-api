package repositories

import (
	"context"

	"github.com/jpbmdev/payment-api/database"
	"github.com/jpbmdev/payment-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// -----------------------------------------------
// -- Target Schema Repository
// -----------------------------------------------
type TargetSchemaRepository interface {
	Find(filter bson.M) (models.TargetSchemas, error)
	FindOneWithSort(sortfilter bson.M, targetSchema *models.TargetSchema) error
}

type targetSchemaRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

//Function to crete new targetSchema repository
func NewTargetSchemaRepository() TargetSchemaRepository {
	return &targetSchemaRepository{
		collection: database.GetCollection("targetShemas"),
		ctx:        context.Background(),
	}
}

func (r *targetSchemaRepository) Find(filter bson.M) (models.TargetSchemas, error) {
	targetShemas := []models.TargetSchema{}

	//Get cursor of database
	cur, err := r.collection.Find(r.ctx, filter)

	//Handle errors
	if err != nil {
		return nil, err
	}

	//Iterate over the cursor to get the targetSchemas
	for cur.Next(r.ctx) {
		var targetShema models.TargetSchema
		err = cur.Decode(&targetShema)
		if err != nil {
			return nil, err
		}
		targetShemas = append(targetShemas, targetShema)
	}

	return targetShemas, nil
}

func (r *targetSchemaRepository) FindOneWithSort(sortfilter bson.M, targetSchema *models.TargetSchema) error {
	filter := bson.D{}

	//Define sort to the query
	opts := options.FindOne().SetSort(sortfilter)
	//Get the targetSchema
	err := r.collection.FindOne(r.ctx, filter, opts).Decode(targetSchema)

	//Handle Errors
	if err != nil {
		return err
	}
	return nil
}
