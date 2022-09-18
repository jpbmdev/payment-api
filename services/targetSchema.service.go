package services

import (
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

// -----------------------------------------------
// -- Target Schema service
// -----------------------------------------------
type TargetSchemaService interface {
	GetTargetSchemas() (models.TargetSchemas, error)
	FindLatestTargetSchema(targetSchema *models.TargetSchema) error
}

type targetSchemaService struct {
	respository repositories.TargetSchemaRepository
}

//Function to crete new targetSchema service
func NewTargetSchemaService() TargetSchemaService {
	return &targetSchemaService{
		respository: repositories.NewTargetSchemaRepository(),
	}
}

func (s *targetSchemaService) GetTargetSchemas() (models.TargetSchemas, error) {
	//Create query
	filter := bson.M{}

	targetSchemas, err := s.respository.Find(filter)

	//Handle Errors
	if err != nil {
		return nil, err
	}
	return targetSchemas, nil
}

func (s *targetSchemaService) FindLatestTargetSchema(targetSchema *models.TargetSchema) error {
	//Create Sort query
	sortFilter := bson.M{"version": 1}

	err := s.respository.FindOneWithSort(sortFilter, targetSchema)

	//Handle Errors
	if err != nil {
		return err
	}
	return nil
}
