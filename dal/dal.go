package main

//go get go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to represent your data model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}

type SecuritySystem struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UUID              string             `bson:"UUID"`
	Location          string             `bson:"Location"`
	SensorType        string             `bson:"SensorType"`
	Status            bool               `bson:"Status"`
	EnergyConsumption int                `bson:"EnergyConsumption"`
	LastTriggered     time.Time          `bson:"LastTriggered"`
}

type Lighting struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UUID              string             `bson:"UUID"`
	Location          string             `bson:"Location"`
	Brightness        string             `bson:"Brightness"`
	Status            bool               `bson:"Status"`
	EnergyConsumption int                `bson:"EnergyConsumption"`
	LastChanged       time.Time          `bson:"LastChanged"`
}

type Microwave struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UUID              string             `bson:"UUID"`
	Status            bool               `bson:"Status"`
	Power             string             `bson:"Power"`
	TimerStopTime     time.Time          `bson:"TimerStopTime"`
	EnergyConsumption int                `bson:"EnergyConsumption"`
	LastChanged       time.Time          `bson:"LastChanged"`
}

type Dishwasher struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UUID              string             `bson:"UUID"`
	Status            bool               `bson:"Status"`
	WashTime          string             `bson:"WashTime"`
	TimerStopTime     time.Time          `bson:"TimerStopTime"`
	EnergyConsumption int                `bson:"EnergyConsumption"`
	LastChanged       time.Time          `bson:"LastChanged"`
}

type Fridge struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	UUID                string             `bson:"UUID"`
	Status              bool               `bson:"Status"`
	TemperatureSettings string             `bson:"TemperatureSettings"`
	EnergyConsumption   int                `bson:"EnergyConsumption"`
	LastChanged         time.Time          `bson:"LastChanged"`
	EnergySaveMode      bool               `bson:"EnergySaveMode"`
}

type HVAC struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UUID              string             `bson:"UUID"`
	Location          string             `bson:"Location"`
	Temperature       string             `bson:"Temperature"`
	Humidity          string             `bson:"Humidity"`
	FanSpeed          string             `bson:"FanSpeed"`
	Status            bool               `bson:"Status"`
	EnergyConsumption int                `bson:"EnergyConsumption"`
	LastChanged       time.Time          `bson:"LastChanged"`
}
type Oven struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	UUID                string             `bson:"UUID"`
	Status              bool               `bson:"Status"`
	TemperatureSettings string             `bson:"TemperatureSettings"`
	TimerStopTime       time.Time          `bson:"TimerStopTime"`
	EnergyConsumption   int                `bson:"EnergyConsumption"`
	LastChanged         time.Time          `bson:"LastChanged"`
}

type Toaster struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	UUID                string             `bson:"UUID"`
	Status              bool               `bson:"Status"`
	TemperatureSettings string             `bson:"TemperatureSettings"`
	TimerStopTime       time.Time          `bson:"TimerStopTime"`
	EnergyConsumption   int                `bson:"EnergyConsumption"`
	LastChanged         time.Time          `bson:"LastChanged"`
}

type SolarPanel struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	UUID                 string             `bson:"UUID"`
	PanelID              string             `bson:"PanelID"`
	Status               bool               `bson:"Status"`
	EnergyGeneratedToday int                `bson:"EnergyGeneratedToday"`
	PowerOutput          int                `bson:"PowerOutput"`
	LastChanged          time.Time          `bson:"LastChanged"`
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

// DEVELOPER ONLY FUNCTION
func deleteAllUsers(client *mongo.Client) error {
	collection := client.Database(dbName).Collection("users")

	_, err := collection.DeleteMany(context.Background(), bson.M{})
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
		Name:     "john_does",
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

	deleteAllUsers(client)
}
