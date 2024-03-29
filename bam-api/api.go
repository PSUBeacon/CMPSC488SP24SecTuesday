package main

import (
	"CMPSC488SP24SecTuesday/dal"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	//"github.com/joho/godotenv"
	"CMPSC488SP24SecTuesday/dal"
	"github.com/gin-contrib/cors"
	//"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)


func main() {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	// unprotected endpoints no auth needed
	r.GET("/status", statusResp)
	r.POST("/login", loginHandler)
	r.GET("/logout", logout)
  r.POST("/signup", signupHandler)


	// Apply JWT middleware to protected routes
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(authMiddleware())

	// This ones should be protected
	protectedRoutes.POST("/lighting", updateIoT)
	protectedRoutes.POST("/hvac", updateIoT)
	protectedRoutes.POST("/security", updateIoT)
	protectedRoutes.POST("/appliances", getAppliancesData)
	protectedRoutes.POST("/energy", getAppliancesData)

	// use JWT middleware for all protected routes
	//r.Use(authMiddleware())


	// route group for admin endpoints
	adminGroup := r.Group("/admin")
	adminGroup.Use(adminMiddleware())
	{
		// Example route requiring admin role
		adminGroup.GET("/admin-dashboard", adminDashboardHandler)
		// Add more admin-only routes as needed
	}

	// Route group for user endpoints
	userGroup := r.Group("/user")
	userGroup.Use(userMiddleware())
	{
		// Example route for user profile
		userGroup.GET("/dashboard", userProfileHandler)
		// Add more user-only routes as needed
	}

	err := r.Run(":8081")
	if err != nil {
		log.Fatal("Server startup error:", err)
	}

} // closes main 

func updateIoT(c *gin.Context) {
	//var req dal.UpdateLightingRequest
	var req dal.MessagingStruct
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("Received request body: %s\n", string(requestBody))
	// Reset the request body to be able to parse it again
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dal.UpdateMessaging(client, []byte(req.UUID), req.Name, req.AppType, req.Function, req.Change)
	// Respond to the request indicating success.
	c.JSON(http.StatusOK, gin.H{"message": "Lighting updated successfully"})

}

func getAppliancesData(c *gin.Context) {
	//var req dal.UpdateLightingRequest
	var req dal.SmartHomeDB
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("Received request body: %s\n", string(requestBody))
	// Reset the request body to be able to parse it again
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch smart home data
	smartHomeDB, err := dal.FetchCollections(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
		return
	}

	//fmt.Printf("%+v\n", smartHomeDB)

	// Respond to the FE request, indicating success.
	c.JSON(http.StatusOK, smartHomeDB)

} // closes get apliances 

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
	session := sessions.Default(c)
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}


	// Connect to MongoDB
	client, err := dal.ConnectToMongoDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}
	defer client.Disconnect(context.Background())
	
	// Compare the password hash using bcrypt.CompareHashAndPassword
  fetchedUser, err := dal.FetchUser(client, "name", loginData.Username)
  if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username (not found in DB"})
		return
	}
  
	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(loginData.Password))
	if err != nil {
 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid  password"})
		return
	}

	// JWT token creation
	claims := jwt.MapClaims{
		"username": fetchedUser.Name,
		"role":     "admin",                               // Replace with the actual role from MongoDB
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	session.Set(userkey, loginData.Username)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	// Return the JWT token in the response

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
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

// handle admin
func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve user claims from the request context
		userClaims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User claims not found"})
			c.Abort()
			return
		}

		// Extract role from user claims
		role, ok := userClaims.(jwt.MapClaims)["role"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found in user claims"})
			c.Abort()
			return
		}

		// Perform role-based authorization check
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
			c.Abort()
			return
		}

		// Proceed to the next middleware or route handler
		c.Next()
	}
}

// handle user
func userMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve user claims from the request context
		userClaims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User claims not found"})
			c.Abort()
			return
		}

		// Extract role from user claims
		role, ok := userClaims.(jwt.MapClaims)["role"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found in user claims"})
			c.Abort()
			return
		}

		// Perform role-based authorization check
		if role != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
			c.Abort()
			return
		}

		// Proceed to the next middleware or route handler
		c.Next()
	}
}

// handle admin dashboard
func adminDashboardHandler(c *gin.Context) {
	// Retrieve user claims from the request context
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User claims not found"})
		return
	}

	// Extract role from user claims
	role, ok := userClaims.(jwt.MapClaims)["role"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Role not found in user claims"})
		return
	}

	// Perform role-based authorization check
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient privileges"})
		return
	}

	// Proceed with the handler logic for admin dashboard
	// For example, return information specific to the admin dashboard
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the admin dashboard!"})
}

// handle user dashboard
func userProfileHandler(c *gin.Context) {
	// Retrieve user claims from the request context
	userClaims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User claims not found"})
		return
	}

	// Extract username from user claims
	username, ok := userClaims.(jwt.MapClaims)["username"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found in user claims"})
		return
	}

	// Proceed with the handler logic for user profile
	// For example, return user-specific data or render a page
	userProfile := getUserProfile(username)
	if userProfile == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User profile not found"})
		return
	}

	// Return user profile data
	c.JSON(http.StatusOK, gin.H{"username": userProfile.FirstName, "email": userProfile.Email})
}

func getUserProfile(username string) *UserProfile {
	// Example implementation to fetch user profile from the database
	// You should replace this with your actual implementation
	// This function should query the database to retrieve the user profile based on the provided username
	// Return nil if user profile not found or an error occurs

	// Example:
	// userProfile, err := db.GetUserProfileByUsername(username)
	// if err != nil {
	//     return nil
	// }
	// return userProfile

	// For demonstration purposes, return a hardcoded user profile
	return &UserProfile{
		//Username: username,
		Email: "user@example.com",
	}
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
				"UUID":              smartHomeDB.Dishwasher[0].UUID,
				"Status":            smartHomeDB.Dishwasher[0].Status,
				"WashTime":          smartHomeDB.Dishwasher[0].WashTime,
				"TimerStopTime":     smartHomeDB.Dishwasher[0].TimerStopTime,
				"EnergyConsumption": smartHomeDB.Dishwasher[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.Dishwasher[0].LastChanged,
			},
			"Fridge": gin.H{
				"UUID":                smartHomeDB.Fridge[0].UUID,
				"Status":              smartHomeDB.Fridge[0].Status,
				"TemperatureSettings": smartHomeDB.Fridge[0].TemperatureSettings,
				"EnergyConsumption":   smartHomeDB.Fridge[0].EnergyConsumption,
				"LastChanged":         smartHomeDB.Fridge[0].LastChanged,
				"EnergySaveMode":      smartHomeDB.Fridge[0].EnergySaveMode,
			},
			"HVAC": gin.H{
				"UUID":              smartHomeDB.HVAC[0].UUID,
				"Location":          smartHomeDB.HVAC[0].Location,
				"Temperature":       smartHomeDB.HVAC[0].Temperature,
				"Humidity":          smartHomeDB.HVAC[0].Humidity,
				"FanSpeed":          smartHomeDB.HVAC[0].FanSpeed,
				"Status":            smartHomeDB.HVAC[0].Status,
				"EnergyConsumption": smartHomeDB.HVAC[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.HVAC[0].LastChanged,
			},
			"Lighting": gin.H{
				"UUID":              smartHomeDB.Lighting[0].UUID,
				"Location":          smartHomeDB.Lighting[0].Location,
				"Brightness":        smartHomeDB.Lighting[0].Brightness,
				"Status":            smartHomeDB.Lighting[0].Status,
				"EnergyConsumption": smartHomeDB.Lighting[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.Lighting[0].LastChanged,
			},
			"Microwave": gin.H{
				"UUID":              smartHomeDB.Microwave[0].UUID,
				"Status":            smartHomeDB.Microwave[0].Status,
				"Power":             smartHomeDB.Microwave[0].Power,
				"TimerStopTime":     smartHomeDB.Microwave[0].TimerStopTime,
				"EnergyConsumption": smartHomeDB.Microwave[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.Microwave[0].LastChanged,
			},
			"Oven": gin.H{
				"UUID":                smartHomeDB.Oven[0].UUID,
				"Status":              smartHomeDB.Oven[0].Status,
				"TemperatureSettings": smartHomeDB.Oven[0].TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Oven[0].TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Oven[0].EnergyConsumption,
				"LastChanged":         smartHomeDB.Oven[0].LastChanged,
			},
			"SecuritySystem": gin.H{
				"UUID":              smartHomeDB.SecuritySystem[0].UUID,
				"Location":          smartHomeDB.SecuritySystem[0].Location,
				"Status":            smartHomeDB.SecuritySystem[0].Status,
				"EnergyConsumption": smartHomeDB.SecuritySystem[0].EnergyConsumption,
				"LastTriggered":     smartHomeDB.SecuritySystem[0].LastTriggered,
			},
			"SolarPanel": gin.H{
				"UUID":                 smartHomeDB.SolarPanel[0].UUID,
				"PanelID":              smartHomeDB.SolarPanel[0].PanelID,
				"Status":               smartHomeDB.SolarPanel[0].Status,
				"EnergyGeneratedToday": smartHomeDB.SolarPanel[0].EnergyGeneratedToday,
				"PowerOutput":          smartHomeDB.SolarPanel[0].PowerOutput,
				"LastChanged":          smartHomeDB.SolarPanel[0].LastChanged,
			},
			"Toaster": gin.H{
				"UUID":                smartHomeDB.Toaster[0].UUID,
				"Status":              smartHomeDB.Toaster[0].Status,
				"TemperatureSettings": smartHomeDB.Toaster[0].TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Toaster[0].TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Toaster[0].EnergyConsumption,
				"LastChanged":         smartHomeDB.Toaster[0].LastChanged,
			},
		})

	case "read": //child role
		c.JSON(http.StatusOK, gin.H{
			"message":     "Welcome to the Owner dashboard",
			"accountType": "Child",
			"Dishwasher": gin.H{
				"UUID":              smartHomeDB.Dishwasher[0].UUID,
				"Status":            smartHomeDB.Dishwasher[0].Status,
				"WashTime":          smartHomeDB.Dishwasher[0].WashTime,
				"TimerStopTime":     smartHomeDB.Dishwasher[0].TimerStopTime,
				"EnergyConsumption": smartHomeDB.Dishwasher[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.Dishwasher[0].LastChanged,
			},
			"Fridge": gin.H{
				"UUID":                smartHomeDB.Fridge[0].UUID,
				"Status":              smartHomeDB.Fridge[0].Status,
				"TemperatureSettings": smartHomeDB.Fridge[0].TemperatureSettings,
				"EnergyConsumption":   smartHomeDB.Fridge[0].EnergyConsumption,
				"LastChanged":         smartHomeDB.Fridge[0].LastChanged,
				"EnergySaveMode":      smartHomeDB.Fridge[0].EnergySaveMode,
			},
			"HVAC": gin.H{
				"UUID":              smartHomeDB.HVAC[0].UUID,
				"Location":          smartHomeDB.HVAC[0].Location,
				"Temperature":       smartHomeDB.HVAC[0].Temperature,
				"Humidity":          smartHomeDB.HVAC[0].Humidity,
				"FanSpeed":          smartHomeDB.HVAC[0].FanSpeed,
				"Status":            smartHomeDB.HVAC[0].Status,
				"EnergyConsumption": smartHomeDB.HVAC[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.HVAC[0].LastChanged,
			},
			"Lighting": gin.H{
				"UUID":              smartHomeDB.Lighting[0].UUID,
				"Location":          smartHomeDB.Lighting[0].Location,
				"Brightness":        smartHomeDB.Lighting[0].Brightness,
				"Status":            smartHomeDB.Lighting[0].Status,
				"EnergyConsumption": smartHomeDB.Lighting[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.Lighting[0].LastChanged,
			},
			"Microwave": gin.H{
				"UUID":              smartHomeDB.Microwave[0].UUID,
				"Status":            smartHomeDB.Microwave[0].Status,
				"Power":             smartHomeDB.Microwave[0].Power,
				"TimerStopTime":     smartHomeDB.Microwave[0].TimerStopTime,
				"EnergyConsumption": smartHomeDB.Microwave[0].EnergyConsumption,
				"LastChanged":       smartHomeDB.Microwave[0].LastChanged,
			},
			"Oven": gin.H{
				"UUID":                smartHomeDB.Oven[0].UUID,
				"Status":              smartHomeDB.Oven[0].Status,
				"TemperatureSettings": smartHomeDB.Oven[0].TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Oven[0].TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Oven[0].EnergyConsumption,
				"LastChanged":         smartHomeDB.Oven[0].LastChanged,
			},
			"SecuritySystem": gin.H{
				"UUID":              smartHomeDB.SecuritySystem[0].UUID,
				"Location":          smartHomeDB.SecuritySystem[0].Location,
				"Status":            smartHomeDB.SecuritySystem[0].Status,
				"EnergyConsumption": smartHomeDB.SecuritySystem[0].EnergyConsumption,
				"LastTriggered":     smartHomeDB.SecuritySystem[0].LastTriggered,
			},
			"SolarPanel": gin.H{
				"UUID":                 smartHomeDB.SolarPanel[0].UUID,
				"PanelID":              smartHomeDB.SolarPanel[0].PanelID,
				"Status":               smartHomeDB.SolarPanel[0].Status,
				"EnergyGeneratedToday": smartHomeDB.SolarPanel[0].EnergyGeneratedToday,
				"PowerOutput":          smartHomeDB.SolarPanel[0].PowerOutput,
				"LastChanged":          smartHomeDB.SolarPanel[0].LastChanged,
			},
			"Toaster": gin.H{
				"UUID":                smartHomeDB.Toaster[0].UUID,
				"Status":              smartHomeDB.Toaster[0].Status,
				"TemperatureSettings": smartHomeDB.Toaster[0].TemperatureSettings,
				"TimerStopTime":       smartHomeDB.Toaster[0].TimerStopTime,
				"EnergyConsumption":   smartHomeDB.Toaster[0].EnergyConsumption,
				"LastChanged":         smartHomeDB.Toaster[0].LastChanged,
			},
		})
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role or insufficient privileges"})
	}

func signupHandler(c *gin.Context) {
	var userData UserProfile

	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
  
	// Validate user data (e.g., check for required fields, validate email format, etc.)

	// Connect to MongoDB
	client, err := dal.ConnectToMongoDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}
	defer client.Disconnect(context.Background())

	// Check if the username already exists in the database
	existingUser, err := dal.FetchUser(client, "username", userData.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password before storing it in the database
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	//if err != nil {
	//    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
	//    return
	//}

	// Create a new user object with hashed password
	newUser := dal.User{
		Email:    userData.Email,
		Name:     userData.FirstName + " " + userData.LastName,
		Password: userData.Password, // Replace with hashedPassword
		Role:     "None",            // Set user role
	}

	// Insert the new user into the database
	err = dal.CreateUser(client, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// struct to represent user profile data
type UserProfile struct {
	Email     string `json:"email"` // this is the email and also the username
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
