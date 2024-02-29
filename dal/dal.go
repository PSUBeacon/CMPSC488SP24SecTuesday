package dal

//go get go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
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

func FetchUser(client *mongo.Client, key, value string) (User, error) {
	collection := client.Database(dbName).Collection("users")
	filter := bson.M{key: value}

	var user User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
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
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.Background())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("user_password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	// Example: Creating a new user
	newUser := User{
		Name:     "john",
		Password: string(hashedPassword),
		Email:    "john@example.com",
	}

	err = CreateUser(client, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully!")

	// Example: Fetching a user by username
	//fetchedUser, err := FetchUser(client, "name", "john")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(fetchedUser)
	//fmt.Println(fetchedUser.Name)
}
