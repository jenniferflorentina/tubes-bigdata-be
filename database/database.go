package database

import (
	"context"
	"fmt"
	"log"

	"TubesBigData/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collection *mongo.Collection
var Client *mongo.Client

func Connect() error {
	var err error
	var connectionString string

	// Load credentials
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	dbName := config.Config("DB_NAME")
	collName := config.Config("COLL_NAME")

	connectionString = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)

	log.Println("Connection String: " + connectionString)

	// Create a new Client
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalln("Connect:", err)
		return err
	}

	// Test DB connection by pinging
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalln("Ping:", err)
		return err
	}

	Collection = Client.Database(dbName).Collection(collName)

	log.Println("Opened database connection and loaded collection")

	return nil
}
