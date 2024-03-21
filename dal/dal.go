package dal

//package main

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Secure connection to DB through admin user
var (
	mongoURI = "mongodb://adminDB:%2525R37%3DLA%5EZX@localhost/admin"
	dbName   = "smartHomeDB"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping MongoDB server for connection verification
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// IOT Structure to fit system info
type Dishwasher struct {
	UUID              string    `json:"UUID"`
	Status            bool      `json:"Status"`
	WashTime          int       `json:"WashTime"`
	TimerStopTime     time.Time `json:"TimerStopTime"`
	EnergyConsumption int       `json:"EnergyConsumption"`
	LastChanged       time.Time `json:"LastChanged"`
}

type Fridge struct {
	UUID                string    `json:"UUID"`
	Status              bool      `json:"Status"`
	TemperatureSettings string    `json:"TemperatureSettings"`
	EnergyConsumption   int       `json:"EnergyConsumption"`
	LastChanged         time.Time `json:"LastChanged"`
	EnergySaveMode      bool      `json:"EnergySaveMode"`
}

type HVAC struct {
	UUID              string    `json:"UUID"`
	Location          string    `json:"Location"`
	Temperature       int       `json:"Temperature"`
	Humidity          int       `json:"Humidity"`
	FanSpeed          int       `json:"FanSpeed"`
	Status            bool      `json:"Status"`
	EnergyConsumption int       `json:"EnergyConsumption"`
	LastChanged       time.Time `json:"LastChanged"`
}

type Lighting struct {
	UUID              string    `json:"UUID"`
	Location          string    `json:"Location"`
	Brightness        int       `json:"Brightness"`
	Status            bool      `json:"Status"`
	EnergyConsumption int       `json:"EnergyConsumption"`
	LastChanged       time.Time `json:"LastChanged"`
}

type Microwave struct {
	UUID              string    `json:"UUID"`
	Status            bool      `json:"Status"`
	Power             int       `json:"Power"`
	TimerStopTime     time.Time `json:"TimerStopTime"`
	EnergyConsumption int       `json:"EnergyConsumption"`
	LastChanged       time.Time `json:"LastChanged"`
}

type Oven struct {
	UUID                string    `json:"UUID"`
	Status              bool      `json:"Status"`
	TemperatureSettings string    `json:"TemperatureSettings"`
	TimerStopTime       time.Time `json:"TimerStopTime"`
	EnergyConsumption   int       `json:"EnergyConsumption"`
	LastChanged         time.Time `json:"LastChanged"`
}

type SecuritySystem struct {
	UUID              string    `json:"UUID"`
	Location          string    `json:"Location"`
	SensorType        string    `json:"SensorType"`
	Status            bool      `json:"Status"`
	EnergyConsumption int       `json:"EnergyConsumption"`
	LastTriggered     time.Time `json:"LastTriggered"`
}

type SolarPanel struct {
	UUID                 string    `json:"UUID"`
	PanelID              string    `json:"PanelID"`
	Status               bool      `json:"Status"`
	EnergyGeneratedToday int       `json:"EnergyGeneratedToday"`
	PowerOutput          int       `json:"PowerOutput"`
	LastChanged          time.Time `json:"LastChanged"`
}

type Toaster struct {
	UUID                string    `json:"UUID"`
	Status              bool      `json:"Status"`
	TemperatureSettings string    `json:"TemperatureSettings"`
	TimerStopTime       time.Time `json:"TimerStopTime"`
	EnergyConsumption   int       `json:"EnergyConsumption"`
	LastChanged         time.Time `json:"LastChanged"`
}

type SmartHomeDB struct {
	Dishwasher     []Dishwasher
	Fridge         []Fridge
	HVAC           []HVAC
	Lighting       []Lighting
	Microwave      []Microwave
	Oven           []Oven
	SecuritySystem []SecuritySystem
	SolarPanel     []SolarPanel
	Toaster        []Toaster
}

// messaging struct to send update requests to IoT devices
type messagingStruct struct {
	UUID     string `json:"UUID"`
	Function string `json:"Function"`
	Change   string `json:"Change"`
}

func FetchCollections(client *mongo.Client, dbName string) (*SmartHomeDB, error) {
	smartHomeDB := &SmartHomeDB{}

	// Define a helper function to fetch and decode documents
	fetchAndDecode := func(collectionName string, result interface{}) error {
		return client.Database(dbName).Collection(collectionName).FindOne(context.Background(), bson.D{}).Decode(result)
	}

	// Fetch each collection
	if err := fetchAndDecode("Dishwasher", &smartHomeDB.Dishwasher); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("Fridge", &smartHomeDB.Fridge); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("HVAC", &smartHomeDB.HVAC); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("Lighting", &smartHomeDB.Lighting); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("Microwave", &smartHomeDB.Microwave); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("Oven", &smartHomeDB.Oven); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("SecuritySystem", &smartHomeDB.SecuritySystem); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("SolarPanel", &smartHomeDB.SolarPanel); err != nil {
		return nil, err
	}
	if err := fetchAndDecode("Toaster", &smartHomeDB.Toaster); err != nil {
		return nil, err
	}

	return smartHomeDB, nil
}

///////////////////////////////////////////////////////////

// User structure to fit system Info
type User struct {
	User       string                 `json:"user"`
	UserID     primitive.Binary       `json:"userId"`
	CustomData map[string]interface{} `json:"customData"`
	Role       struct {
		Role string `json:"role"`
		DB   string `json:"db"`
	} `json:"role"`
}

func FetchUser(client *mongo.Client, userName string) (User, error) {
	var tempResult struct {
		Users []struct {
			User       string                 `bson:"user"`
			UserID     primitive.Binary       `bson:"userId"`
			CustomData map[string]interface{} `bson:"customData"`
			Roles      []struct {
				Role string `bson:"role"`
				DB   string `bson:"db"`
			} `bson:"roles"`
		} `bson:"users"`
	}

	cmd := bson.D{
		{Key: "usersInfo", Value: bson.M{
			"user": userName,
			"db":   dbName,
		}},
	}

	err := client.Database(dbName).RunCommand(context.Background(), cmd).Decode(&tempResult)
	if err != nil {
		return User{}, err
	}

	if len(tempResult.Users) > 0 {
		user := tempResult.Users[0]
		var simplifiedUser User
		simplifiedUser.User = user.User
		simplifiedUser.UserID = user.UserID
		simplifiedUser.CustomData = user.CustomData
		if len(user.Roles) > 0 {
			simplifiedUser.Role.Role = user.Roles[0].Role
			simplifiedUser.Role.DB = user.Roles[0].DB
		}
		return simplifiedUser, nil
	}

	return User{}, fmt.Errorf("no user found with username: %s", userName)
}

func Iotlighting(UUID []byte, status string, dim string) {

	client, err := ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.Background())
	smartHomeDB, err := FetchCollections(client, dbName) // Fetches and populates data

	if err != nil {
		log.Fatalf("Error fetching IoT data: %v", err)
	}
	for _, light := range smartHomeDB.Lighting {
		if bytes.Equal([]byte(light.UUID), UUID) {
			light.Status = status
			light.Brightness = dim
			var infoChange messagingStruct
			infoChange.UUID = string(UUID)
			infoChange.Function = "status"
			infoChange.Change = status
			message, _ := json.MarshalIndent(infoChange, "", "  ")
			fmt.Printf("This is the message: ", message)
			messaging.BroadCastMessage(message)

		}
	}
}

//func PrintSmartHomeDBContents(smartHomeDB *SmartHomeDB) string {
//	return fmt.Sprintf(
//		"Dishwasher: %+v\nFridge: %+v\nHVAC: %+v\nLighting: %+v\nMicrowave: %+v\nOven: %+v\nSecuritySystem: %+v\nSolarPanel: %+v\nToaster: %+v\n",
//		*smartHomeDB.Dishwasher,
//		*smartHomeDB.Fridge,
//		*smartHomeDB.HVAC,
//		*smartHomeDB.Lighting,
//		*smartHomeDB.Microwave,
//		*smartHomeDB.Oven,
//		*smartHomeDB.SecuritySystem,
//		*smartHomeDB.SolarPanel,
//		*smartHomeDB.Toaster,
//	)
//}

// connect to db if exists, else return error log
func main() {
	client, err := ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	////Testing fetchedUser function
	fetchedUser, err := FetchUser(client, "Owner")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User name: %s\n", fetchedUser.User)
	fmt.Printf("Password: %v\n", fetchedUser.CustomData["password"])
	fmt.Printf("UserID: %x\n", fetchedUser.UserID.Data)
	fmt.Printf("Role: %s\n", fetchedUser.Role.Role)
	fmt.Printf("Role DB: %s\n", fetchedUser.Role.DB)

	//Testing IoT functions
	// Fetch the IoT system data
	//smartHomeDB, err := FetchCollections(client, dbName) // Fetches and populates data
	if err != nil {
		log.Fatalf("Error fetching IoT data: %v", err)
	}

	//fmt.Printf("HVAC Temperature: %s\n", smartHomeDB.HVAC.Temperature)

	//fmt.Printf("Dishwasher Status: %s\n", smartHomeDB.Dishwasher.Status)

	//fmt.Printf("Oven UUID: %s\n", smartHomeDB.Oven.UUID)

	// Print the contents of smartHomeDB
	//fmt.Printf(PrintSmartHomeDBContents(smartHomeDB))

}
