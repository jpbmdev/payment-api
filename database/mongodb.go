package database

import (
	"context"
	"time"

	"github.com/jpbmdev/payment-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Function to create connection with mongodb collection
func GetCollection(collection string) *mongo.Collection {
	//Get connection data
	connectionString := config.ConfigSchema.DBConnectionString
	dbName := config.ConfigSchema.DBName

	//create mongodb client
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))

	if err != nil {
		panic(err.Error())
	}

	//create context and connect to the client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	//Return mongodb collection
	return client.Database(dbName).Collection(collection)

}
