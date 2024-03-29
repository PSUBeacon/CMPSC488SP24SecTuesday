package dal

//go get go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to represent your data model
type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Role     string `bson:"role"`
}

// MongoDB configuration
var (
	mongoURI = "mongodb://localhost:27017" // MongoDB server URI
	dbName   = "smartHomeDB"               // Collection name
)

// Connect to MongoDB and return a MongoDB client
func ConnectToMongoDB() (*mongo.Client, error) {
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
func CreateUser(client *mongo.Client, user User) error {
	collection := client.Database(dbName).Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func FetchUser(client *mongo.Client, key, value string) (*User, error) {
	collection := client.Database(dbName).Collection("users")
	filter := bson.M{key: value}

	var user User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err // Return nil user and error if user not found or error occurs
	}
	return &user, nil // Return pointer to user and nil error if user found
}

func deleteUser(client *mongo.Client, key, value string) error {
	collection := client.Database(dbName).Collection("users")
	filter := bson.M{key: value}

	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

func main() {
	// Connect to MongoDB
	client, err := ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Example: Creating a new user
	newUser := User{
		Name:     "john_doess",
		Password: "passlel",
		Email:    "john@example.com",
	}

	err = CreateUser(client, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully!")

	// Example: Fetching a user by username
	fetchedUser, err := FetchUser(client, "name", "john_doess")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fetchedUser)
	fmt.Println(fetchedUser.Name)
}
