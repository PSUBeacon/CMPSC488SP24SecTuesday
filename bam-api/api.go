package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	//"github.com/joho/godotenv"
	"CMPSC488SP24SecTuesday/dal"
	"github.com/gin-contrib/cors"
	"go.mongodb.org/mongo-driver/mongo"
	//"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var client *mongo.Client

func main() {
	//connect to mongo server
	var er error
	client, er = dal.ConnectToMongoDB()
	if er != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", er)
	}
	defer func() {
		if er = client.Disconnect(context.Background()); er != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", er)
		}
	}()

	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //for open access
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// unprotected endpoints no auth needed
	r.GET("/status", statusResp)
	r.POST("/login", loginHandler)

	// use JWT middleware for all protected routes
	r.Use(authMiddleware())

	//ADJUSTMENT:
	// Combined route group for both admin and user dashboards
	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use(authMiddleware()) // Apply authMiddleware to protect the route
	{
		// Combined dashboard route for admin and user
		dashboardGroup.GET("", dashboardHandler) // Use an empty string for the base path of the group
	}

	err := r.Run(":8081")
	if err != nil {
		log.Fatal("Server startup error:", err)
	}

}

func statusResp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func getJwtKey() string {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		// fallback value for the key
		key = "saflakjrwlierueorel121321!"
	}
	return key
}

// get env secret key for jwt
var jwtKey = []byte(getJwtKey())

func loginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	//ADJUSTMENT:Changed fetchUser parameters to reflect updated DAL
	// Fetch user by username from MongoDB
	fetchedUser, err := dal.FetchUser(client, loginData.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username (not found in DB"})
		return
	}

	//ADJUSTMENT: Changed function for password
	// Compare the password hash using bcrypt.CompareHashAndPassword
	//err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(loginData.Password))
	//if err != nil {
	fmt.Printf("%s\n", fetchedUser.UserID.Data)
	if fetchedUser.CustomData["password"] != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid  password"})
		return
	}

	// Example of fetching the smartHomeDB data, adjust based on actual implementation

	// Fetch smart home data
	smartHomeDB, err := dal.FetchCollections(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
		return
	}
	//c.Set("smartHomeDB", smartHomeDB)

	// Print smart home data
	smartHomeData := dal.PrintSmartHomeDBContents(smartHomeDB)

	//Delete later
	fmt.Printf(smartHomeData)
	//////////////////////////

	//ADJUSTMENT:Changed JWT generated payload
	// JWT token creation
	claims := jwt.MapClaims{

		"username": fetchedUser.User,
		"userID":   fetchedUser.UserID.Data,
		"role":     fetchedUser.Role.Role,                 // Replace with the actual role from MongoDB
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	////////////////////

	c.JSON(http.StatusOK, gin.H{"token": tokenString,
		"HVAC": smartHomeDB.HVAC.Temperature})
}

// check jwt auth and set user role
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Printf("Authorization header: %s\n", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set the user claims in the request context
			fmt.Printf("Authorized")
			c.Set("user", claims)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}

// Combined dashboard handler for both admin and child.
func dashboardHandler(c *gin.Context) {
	// Retrieve user claims from the request context
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User claims not found"})
		return
	}

	claims, ok := userClaims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing user claims"})
		return
	}

	// Extract role from user claims
	role, roleExists := claims["role"].(string)
	if !roleExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found in user claims"})
		return
	}

	// Fetch smart home data
	smartHomeDB, err := dal.FetchCollections(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
		return
	}

	// Determine the response based on the user's role
	switch role {
	case "readWrite": //owner role
		if smartHomeDB == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":     "Welcome to the Owner dashboard",
			"accountType": "Owner",
			"Dishwasher": gin.H{
				"UUID":              smartHomeDB.Dishwasher.UUID,
				"Status":            smartHomeDB.Dishwasher.Status,
				"WashTime":          smartHomeDB.Dishwasher.WashTime,
				"TimerStopTime":     smartHomeDB.Dishwasher.TimerStopTime,
				"EnergyConsumption": smartHomeDB.Dishwasher.EnergyConsumption,
				"LastChanged":       smartHomeDB.Dishwasher.LastChanged,
			},
			"Fridge": gin.H{
				"UUID":                smartHomeDB.Fridge.UUID,
				"Status":              smartHomeDB.Fridge.Status,
				"TemperatureSettings": smartHomeDB.Fridge.TemperatureSettings,
				"EnergyConsumption":   smartHomeDB.Fridge.EnergyConsumption,
				"LastChanged":         smartHomeDB.Fridge.LastChanged,
				"EnergySaveMode":      smartHomeDB.Fridge.EnergySaveMode,
			},
			"HVAC": gin.H{
				"UUID":              smartHomeDB.HVAC.UUID,
				"Location":          smartHomeDB.HVAC.Location,
				"Temperature":       smartHomeDB.HVAC.Temperature,
				"Humidity":          smartHomeDB.HVAC.Humidity,
				"FanSpeed":          smartHomeDB.HVAC.FanSpeed,
				"Status":            smartHomeDB.HVAC.Status,
				"EnergyConsumption": smartHomeDB.HVAC.EnergyConsumption,
				"LastChanged":       smartHomeDB.HVAC.LastChanged,
			},
			"Lighting": gin.H{
				"UUID":              smartHomeDB.Lighting.UUID,
				"Location":          smartHomeDB.Lighting.Location,
				"Brightness":        smartHomeDB.Lighting.Brightness,
				"Status":            smartHomeDB.Lighting.Status,
				"EnergyConsumption": smartHomeDB.Lighting.EnergyConsumption,
				"LastChanged":       smartHomeDB.Lighting.LastChanged,
			},
			"Microwave": gin.H{
				"UUID":              smartHomeDB.Microwave.UUID,
				"Status":            smartHomeDB.Microwave.Status,
				"Power":             smartHomeDB.Microwave.Power,
				"TimerStopTime":     smartHomeDB.Microwave.TimerStopTime,
				"EnergyConsumption": smartHomeDB.Microwave.EnergyConsumption,
				"LastChanged":       smartHomeDB.Microwave.LastChanged,
			},
			"Oven": gin.H{
				"UUID":                smartHomeDB.Oven.UUID,
				"Status":              smartHomeDB.Oven.Status,
				"TemperatureSettings": smartHomeDB.Oven.TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Oven.TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Oven.EnergyConsumption,
				"LastChanged":         smartHomeDB.Oven.LastChanged,
			},
			"SecuritySystem": gin.H{
				"UUID":              smartHomeDB.SecuritySystem.UUID,
				"Location":          smartHomeDB.SecuritySystem.Location,
				"SensorType":        smartHomeDB.SecuritySystem.SensorType,
				"Status":            smartHomeDB.SecuritySystem.Status,
				"EnergyConsumption": smartHomeDB.SecuritySystem.EnergyConsumption,
				"LastTriggered":     smartHomeDB.SecuritySystem.LastTriggered,
			},
			"SolarPanel": gin.H{
				"UUID":                 smartHomeDB.SolarPanel.UUID,
				"PanelID":              smartHomeDB.SolarPanel.PanelID,
				"Status":               smartHomeDB.SolarPanel.Status,
				"EnergyGeneratedToday": smartHomeDB.SolarPanel.EnergyGeneratedToday,
				"PowerOutput":          smartHomeDB.SolarPanel.PowerOutput,
				"LastChanged":          smartHomeDB.SolarPanel.LastChanged,
			},
			"Toaster": gin.H{
				"UUID":                smartHomeDB.Toaster.UUID,
				"Status":              smartHomeDB.Toaster.Status,
				"TemperatureSettings": smartHomeDB.Toaster.TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Toaster.TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Toaster.EnergyConsumption,
				"LastChanged":         smartHomeDB.Toaster.LastChanged,
			},
		})

	case "read": //child role
		c.JSON(http.StatusOK, gin.H{
			"message":     "Welcome to the Owner dashboard",
			"accountType": "Child",
			"Dishwasher": gin.H{
				"UUID":              smartHomeDB.Dishwasher.UUID,
				"Status":            smartHomeDB.Dishwasher.Status,
				"WashTime":          smartHomeDB.Dishwasher.WashTime,
				"TimerStopTime":     smartHomeDB.Dishwasher.TimerStopTime,
				"EnergyConsumption": smartHomeDB.Dishwasher.EnergyConsumption,
				"LastChanged":       smartHomeDB.Dishwasher.LastChanged,
			},
			"Fridge": gin.H{
				"UUID":                smartHomeDB.Fridge.UUID,
				"Status":              smartHomeDB.Fridge.Status,
				"TemperatureSettings": smartHomeDB.Fridge.TemperatureSettings,
				"EnergyConsumption":   smartHomeDB.Fridge.EnergyConsumption,
				"LastChanged":         smartHomeDB.Fridge.LastChanged,
				"EnergySaveMode":      smartHomeDB.Fridge.EnergySaveMode,
			},
			"HVAC": gin.H{
				"UUID":              smartHomeDB.HVAC.UUID,
				"Location":          smartHomeDB.HVAC.Location,
				"Temperature":       smartHomeDB.HVAC.Temperature,
				"Humidity":          smartHomeDB.HVAC.Humidity,
				"FanSpeed":          smartHomeDB.HVAC.FanSpeed,
				"Status":            smartHomeDB.HVAC.Status,
				"EnergyConsumption": smartHomeDB.HVAC.EnergyConsumption,
				"LastChanged":       smartHomeDB.HVAC.LastChanged,
			},
			"Lighting": gin.H{
				"UUID":              smartHomeDB.Lighting.UUID,
				"Location":          smartHomeDB.Lighting.Location,
				"Brightness":        smartHomeDB.Lighting.Brightness,
				"Status":            smartHomeDB.Lighting.Status,
				"EnergyConsumption": smartHomeDB.Lighting.EnergyConsumption,
				"LastChanged":       smartHomeDB.Lighting.LastChanged,
			},
			"Microwave": gin.H{
				"UUID":              smartHomeDB.Microwave.UUID,
				"Status":            smartHomeDB.Microwave.Status,
				"Power":             smartHomeDB.Microwave.Power,
				"TimerStopTime":     smartHomeDB.Microwave.TimerStopTime,
				"EnergyConsumption": smartHomeDB.Microwave.EnergyConsumption,
				"LastChanged":       smartHomeDB.Microwave.LastChanged,
			},
			"Oven": gin.H{
				"UUID":                smartHomeDB.Oven.UUID,
				"Status":              smartHomeDB.Oven.Status,
				"TemperatureSettings": smartHomeDB.Oven.TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Oven.TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Oven.EnergyConsumption,
				"LastChanged":         smartHomeDB.Oven.LastChanged,
			},
			"SecuritySystem": gin.H{
				"UUID":              smartHomeDB.SecuritySystem.UUID,
				"Location":          smartHomeDB.SecuritySystem.Location,
				"SensorType":        smartHomeDB.SecuritySystem.SensorType,
				"Status":            smartHomeDB.SecuritySystem.Status,
				"EnergyConsumption": smartHomeDB.SecuritySystem.EnergyConsumption,
				"LastTriggered":     smartHomeDB.SecuritySystem.LastTriggered,
			},
			"SolarPanel": gin.H{
				"UUID":                 smartHomeDB.SolarPanel.UUID,
				"PanelID":              smartHomeDB.SolarPanel.PanelID,
				"Status":               smartHomeDB.SolarPanel.Status,
				"EnergyGeneratedToday": smartHomeDB.SolarPanel.EnergyGeneratedToday,
				"PowerOutput":          smartHomeDB.SolarPanel.PowerOutput,
				"LastChanged":          smartHomeDB.SolarPanel.LastChanged,
			},
			"Toaster": gin.H{
				"UUID":                smartHomeDB.Toaster.UUID,
				"Status":              smartHomeDB.Toaster.Status,
				"TemperatureSettings": smartHomeDB.Toaster.TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Toaster.TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Toaster.EnergyConsumption,
				"LastChanged":         smartHomeDB.Toaster.LastChanged,
			},
		})
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role or insufficient privileges"})
	}
}
