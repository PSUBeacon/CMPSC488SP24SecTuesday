package dal

//go get go.mongodb.org/mongo-driver/mongo

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" 
	"log"
 
	"strconv"
	"time"
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
	Users          []User
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

func main() {
	// Connect to MongoDB
	client, err := ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// 	// Example: Creating a new user
	// 	newUser := User{
	// 		Name:     "john_doess",
	// 		Password: "passlel",
	// 		Email:    "john@example.com",
	// 	}

	// 	err = CreateUser(client, newUser)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("User created successfully!")

	// Example: Fetching a user by username
	// 	fetchedUser, err := FetchUser(client, "name", "john_doess")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(fetchedUser)
	// 	fmt.Println(fetchedUser.Name)

}
