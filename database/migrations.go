package database

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/jpbmdev/payment-api/config"
	"github.com/jpbmdev/payment-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Function to create connection with mongodb collection
func Migrations() {
	//Get connection data
	connectionString := config.ConfigSchema.DBConnectionString

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

	//Check if the database exists
	query := bson.D{}
	names, err := client.ListDatabaseNames(ctx, query)
	if err != nil {
		panic(err.Error())
	}
	find := false
	for i := range names {
		if names[i] == config.ConfigSchema.DBName {
			find = true
		}
	}

	//If the database does not exists create id with a targetSchena and test users
	if !find {
		//Create db
		client.Database(config.ConfigSchema.DBName).CreateCollection(ctx, "targetShemas")
		//Load targetSchema
		targetSchemaBytes, _ := ioutil.ReadFile("database/migrations/testTree.json")
		targetSchema := models.TargetSchemaSwagger{}
		json.Unmarshal(targetSchemaBytes, &targetSchema)
		client.Database(config.ConfigSchema.DBName).Collection("targetShemas").InsertOne(ctx, targetSchema)
		//LoadUsers
		usersBytes, _ := ioutil.ReadFile("database/migrations/users.json")
		var users models.Users
		json.Unmarshal(usersBytes, &users)
		newUsers := make([]interface{}, len(users))
		for i := range users {
			newUsers[i] = users[i]
		}
		client.Database(config.ConfigSchema.DBName).Collection("users").InsertMany(ctx, newUsers)
		//Load Loans
		loanBytes, _ := ioutil.ReadFile("database/migrations/loans.json")
		var loans models.Loans
		json.Unmarshal(loanBytes, &loans)
		newLoans := make([]interface{}, len(loans))
		for i := range loans {
			newLoans[i] = loans[i]
		}
		client.Database(config.ConfigSchema.DBName).Collection("loans").InsertMany(ctx, newLoans)
		//Load payments
		paymentsBytes, _ := ioutil.ReadFile("database/migrations/payments.json")
		var payments models.Payments
		json.Unmarshal(paymentsBytes, &payments)
		newPayments := make([]interface{}, len(payments))
		for i := range payments {
			newPayments[i] = payments[i]
		}
		client.Database(config.ConfigSchema.DBName).Collection("payments").InsertMany(ctx, newPayments)
	}

}
