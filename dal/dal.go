package dal

//package main

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

// Secure connection to DB through admin user
var (
	mongoURI = "mongodb://localhost:27017"
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
	TemperatureSettings int       `json:"TemperatureSettings"`
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
	Mode              string    `json:"Mode"`
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
	TemperatureSettings int       `json:"TemperatureSettings"`
	TimerStopTime       time.Time `json:"TimerStopTime"`
	EnergyConsumption   int       `json:"EnergyConsumption"`
	LastChanged         time.Time `json:"LastChanged"`
}

type SecuritySystem struct {
	UUID              string    `json:"UUID"`
	Location          string    `json:"Location"`
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
	TemperatureSettings int       `json:"TemperatureSettings"`
	TimerStopTime       time.Time `json:"TimerStopTime"`
	EnergyConsumption   int       `json:"EnergyConsumption"`
	LastChanged         time.Time `json:"LastChanged"`
}

type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Role     string `json:"Role"`
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

type UUIDsConfig struct {
	LightingUUIDs   []Pi
	HvacUUIDs       []Pi
	SecurityUUIDs   []Pi
	AppliancesUUIDs []Pi
	EnergyUUIDs     []Pi
}

type Pi struct {
	Pinum int    `json:"Pinum"`
	UUID  string `json:"UUID"`
}

type MessagingStruct struct {
	UUID     string `json:"UUID"`
	Name     string `json:"Name"`     //type item being changes ex(Lighting) or (HVAC)
	AppType  string `json:"AppType"`  //Which Type of appliance it is
	Function string `json:"Function"` //function being changed ex(brightness)
	Change   string `json:"Change"`   //actual change being made ex(100) for brightness
}

type LoggingStruct struct {
	DeviceID string    `json:"DeviceID"`
	Function string    `json:"Function"`
	Change   string    `json:"Change"`
	Time     time.Time `json:"Time"`
}

func FetchCollections(client *mongo.Client, dbName string) (*SmartHomeDB, error) {
	smartHomeDB := &SmartHomeDB{}

	// Define a helper function to fetch and decode documents
	fetchAndDecode := func(collectionName string, results interface{}) error {
		cursor, err := client.Database(dbName).Collection(collectionName).Find(context.Background(), bson.D{})
		if err != nil {
			return err
		}
		defer cursor.Close(context.Background())
		if err = cursor.All(context.Background(), results); err != nil {
			return err
		}
		return nil
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

func FetchUser(client *mongo.Client, userName string) (User, error) {

	collection := client.Database(dbName).Collection("Users")

	// Create a filter to specify the criteria of the query
	filter := bson.M{"username": userName}

	// Finding multiple documents returns a cursor
	var user User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// User was not found
			return User{}, fmt.Errorf("no user found with username: %s", userName)
		}
		// Some other error occurred
		return User{}, err
	}

	fmt.Printf("Username: %s, Role: %s\n", user.Username, user.Role)
	return user, nil
}

func UpdateMessaging(client *mongo.Client, UUID []byte, name string, apptype string, function string, change string) {
	var messageRequest MessagingStruct
	messageRequest.UUID = string(UUID)
	messageRequest.Name = name
	messageRequest.AppType = apptype
	messageRequest.Function = function
	messageRequest.Change = change
	message, err := json.Marshal(messageRequest)
	if err != nil {
		fmt.Printf("Error marshaling JSON message: %v", err)
		return
	}
	messaging.BroadCastMessage(message)

	var logg LoggingStruct
	// Update the MongoDB collection using JSON
	collection := client.Database("smartHomeDB").Collection(apptype)
	Logging := client.Database("smartHomeDB").Collection("Logging")
	filter := bson.M{"UUID": messageRequest.UUID}

	var updatedValue interface{}
	switch function {
	case "WashTime", "TemperatureSettings", "EnergyConsumption", "Brightness", "Power":
		updatedValue, _ = strconv.Atoi(change)
	default:
		updatedValue = change
	}

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			function: updatedValue,
		},
	}

	updateJSON, err := json.Marshal(update)
	if err != nil {
		fmt.Printf("Error marshaling update JSON: %v", err)
		return
	}

	var updateDoc bson.M
	if err := bson.UnmarshalExtJSON(updateJSON, true, &updateDoc); err != nil {
		fmt.Printf("Error unmarshaling update JSON: %v", err)
		return
	}

	_, err = collection.UpdateOne(context.Background(), filter, updateDoc)
	if err != nil {
		fmt.Printf("Error updating document: %v", err)
		return
	}

	logg.DeviceID = name
	logg.Function = function
	logg.Change = change
	logg.Time = time.Now()
	_, err = Logging.InsertOne(context.Background(), logg)

	return
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
	//fetchedUser, err := FetchUser(client, "Owner")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("User name: %s\n", fetchedUser.User)
	//fmt.Printf("Password: %v\n", fetchedUser.CustomData["password"])
	//fmt.Printf("UserID: %x\n", fetchedUser.UserID.Data)
	//fmt.Printf("Role: %s\n", fetchedUser.Role.Role)
	//fmt.Printf("Role DB: %s\n", fetchedUser.Role.DB)

	//Testing IoT functions
	// Fetch the IoT system data
	//smartHomeDB, err := FetchCollections(client, dbName) // Fetches and populates data
	//if err != nil {
	//	log.Fatalf("Error fetching IoT data: %v", err)
	//}

	//fmt.Printf("HVAC Temperature: %s\n", smartHomeDB.HVAC.Temperature)

	//fmt.Printf("Dishwasher Status: %s\n", smartHomeDB.Dishwasher.Status)

	//fmt.Printf("Oven UUID: %s\n", smartHomeDB.Oven.UUID)

	// Print the contents of smartHomeDB
	//fmt.Printf(PrintSmartHomeDBContents(smartHomeDB))

	fmt.Println(FetchUser(client, "beaconuser"))

}
