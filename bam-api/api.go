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
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var client *mongo.Client

const userkey = "user"

var secret = []byte("secret")

var disconnectedPiNums []int

func UpdateMissingPi(PiNum []string) {
	for i := 0; i < len(PiNum); i++ {
		PiNumInt, _ := strconv.Atoi(PiNum[i])
		disconnectedPiNums = append(disconnectedPiNums, PiNumInt)
	}
}

func main() {
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

	// Apply JWT middleware to protected routes
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(authMiddleware())

	// This ones should be protected
	protectedRoutes.POST("/lighting", updateIoT)
	protectedRoutes.GET("/lighting", GetLights)

	protectedRoutes.POST("/hvac")

	protectedRoutes.POST("/security")
	protectedRoutes.GET("/security", GetSecurity)

	protectedRoutes.POST("/appliances", getAppliancesData)
	protectedRoutes.POST("/energy", getAppliancesData)

	//ADJUSTMENT:

	// Combined route group for both admin and user dashboards
	dashboardGroup := r.Group("/dashboard", dashboardHandler)
	dashboardGroup.Use(authMiddleware()) // Apply authMiddleware to protect the route
	{
		// Combined dashboard route for admin and user
		dashboardGroup.POST("/", me)
		dashboardGroup.GET("/me", me) // Use an empty string for the base path of the group
		dashboardGroup.GET("/status", statusResp)
	}

	networkingGroup := r.Group("/networking")
	networkingGroup.Use()
	{
		networkingGroup.GET("/", me)
		networkingGroup.GET("/GetNetLogs", GetNetLogs)
		networkingGroup.GET("/status", statusResp)
	}

	securityGroup := r.Group("/security")
	securityGroup.Use()
	{
		securityGroup.GET("/", me)
		securityGroup.GET("/GetSecurity", GetSecurity)
		securityGroup.GET("/status", statusResp)
	}

	appliancesGroup := r.Group("/appliances", dashboardHandler)
	appliancesGroup.Use()
	{
		appliancesGroup.GET("/", me)
		appliancesGroup.GET("/me", me)
		appliancesGroup.GET("/status", statusResp)
	}

	go r.Run(":8081")
	//if err != nil {
	//	log.Fatal("Server startup error:", err)
	//}
	// Create a channel to receive the missingPi array from BlockReceiver
	//

	select {}
}

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

	//var UUIDsData dal.UUIDsConfig
	//jsonconfigData, _ := os.ReadFile("config.json")
	//_ = json.Unmarshal(jsonconfigData, &UUIDsData)

	//allPis := [][]dal.Pi{
	//	UUIDsData.LightingUUIDs,
	//	UUIDsData.HvacUUIDs,
	//	UUIDsData.SecurityUUIDs,
	//	UUIDsData.AppliancesUUIDs,
	//	UUIDsData.EnergyUUIDs,
	//}
	//UpdateMissingPi(messaging.)
	//foundPi, found := findPiByUUID(allPis, req.UUID)
	//if found {
	//	for i := 0; i < len(disconnectedPiNums); i++ {
	//		if foundPi.Pinum == disconnectedPiNums[i] {
	//			fmt.Println("Pi is disconnected")
	//		}
	//	}
	//}
	//if !found {
	//	dal.UpdateMessaging(client, []byte(req.UUID), req.Name, req.AppType, req.Function, req.Change)
	//	c.JSON(http.StatusOK, gin.H{"message": "IOT updated successfully"})
	//}
	dal.UpdateMessaging(client, []byte(req.UUID), req.Name, req.AppType, req.Function, req.Change)

}

func GetLights(c *gin.Context) {
	room := c.Query("roomName")
	if room == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room name is required"})
		return
	}
	//fmt.Printf("Room name: %s\n", room)
	lights, err := dal.FetchLights(client, "smartHomeDB", room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("Light:", lights)
	c.JSON(http.StatusOK, lights)
}

func getAppliancesData(c *gin.Context) {
	//var req dal.UpdateLightingRequest
	var req dal.SmartHomeDB
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("Received request body: %s\n", string(requestBody))
	// Reset the request body to be able to parse it again
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	if err := c.ShouldBindJSON(&req); err != nil {
		//fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch smart home data
	smartHomeDB, err := dal.FetchCollections(client, "smartHomeDB")
	if err != nil {
		//fmt.Printf("SMDB data error: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
		return
	}

	//fmt.Printf("%+v\n", smartHomeDB)

	// Respond to the FE request, indicating success.
	c.JSON(http.StatusOK, smartHomeDB)

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
	session := sessions.Default(c)
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Printf("user: %s", loginData.Username)
	fmt.Printf("pass: %s", loginData.Password)
	// Fetch user by username from MongoDB
	fetchedUser, err := dal.FetchUser(client, loginData.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // Use generic error message
		return
	}

	passwordFromDB := fetchedUser.Password

	// Compare the password hash using bcrypt.CompareHashAndPassword
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(loginData.Password))
	if err != nil {
		fmt.Printf("Error: password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // Use generic error message

		return
	}
	//c.Set("smartHomeDB", smartHomeDB)

	// Print smart home data
	//smartHomeData := dal.PrintSmartHomeDBContents(smartHomeDB)

	// JWT token creation
	claims := jwt.MapClaims{
		"username": fetchedUser.Username,
		"role":     fetchedUser.Role, // Use the actual role from the fetched user

		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the token"})
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

	case "admin": //admin

		if smartHomeDB == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Owner dashboard",

			"accountType": "Admin",
		})

	case "user": //user
		c.JSON(http.StatusOK, gin.H{
			"message":     "Welcome to the Owner dashboard",
			"accountType": "User",
		})
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role or insufficient privileges"})
	}

	// Determine the response based on the user's role
	switch role {
	case "readWrite": //owner role
		if smartHomeDB == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
			return
		}
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role or insufficient privileges"})
	}
}

func GetNetLogs(c *gin.Context) {

	//logs, err := dal.FetchLogging(client, "smartHomeDB")
	//if err != nil {
	//	fmt.Printf(err.Error())
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//fmt.Printf("Logs:", logs)
	//c.JSON(http.StatusOK, logs)

	logs, err := dal.FetchLogging(client, "smartHomeDB")
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Remove the problematic line fmt.Printf("Logs:", logs)

	c.JSON(http.StatusOK, logs)
}

//func getHVAC(c *gin.Context) {
//}

func GetSecurity(c *gin.Context) {
	security, err := dal.FetchSecurity(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, security)
}

//
//func GetAppliances(c *gin.Context) {
//}
//
//func GetEnergy(c *gin.Context) {
//}
