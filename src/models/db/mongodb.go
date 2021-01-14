package db

import (
	"context"
	"ginRestFulApi/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoConnection *mongo.Database = nil

//GetConnection is for get mongo connection
func GetConnection() *mongo.Database {
	if mongoConnection == nil {
		dbUrl := utils.EnvVar("DB_URL", "")

		// Database Config
		clientOptions := options.Client().ApplyURI(dbUrl)
		client, err := mongo.NewClient(clientOptions)

		//Set up a context required by mongo.Connect
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		err = client.Connect(ctx)

		//Cancel context to avoid memory leak
		defer cancel()

		// Ping our db connection
		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			log.Fatal("Couldn't connect to the database", err)
		} else {
			log.Println("Connected!")
		}
		// Connect to the database
		mongoConnection := client.Database(utils.EnvVar("DB_NAME", ""))
		return mongoConnection
	}

	return mongoConnection

}
