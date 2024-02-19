package main

//go get go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to represent your data model
type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

// MongoDB configuration
var (
	mongoURI       = "mongodb://localhost:27017" // MongoDB server URI
	dbName         = "mydb"                      // Database name
	collectionName = "users"                     // Collection name
)

// Connect to MongoDB and return a MongoDB client
func connectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// Create a new user in the MongoDB database
func createUser(client *mongo.Client, user User) error {
	collection := client.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func main() {
	// Connect to MongoDB
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Example: Creating a new user
	newUser := User{
		Username: "john_doe",
		Email:    "john@example.com",
	}

	err = createUser(client, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully!")
}
