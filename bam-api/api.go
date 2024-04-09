package main

import (
	"CMPSC488SP24SecTuesday/dal"
	"CMPSC488SP24SecTuesday/networktraffic"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
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

var r = gin.Default() // Configure CORS

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
	protectedRoutes.GET("/lighting", GetLights)

	protectedRoutes.POST("/security", updateIoT)
	protectedRoutes.GET("/security", GetSecurity)

	protectedRoutes.POST("/appliances", getAppliancesData)

	protectedRoutes.POST("/energy", updateIoT)

	//ADJUSTMENT:
	// Combined route group for both admin and user dashboards
	dashboardGroup := r.Group("/dashboard")

	dashboardGroup.Use() // Apply authMiddleware to protect the route
	{
		dashboardGroup.GET("/", dashboardHandler)
		dashboardGroup.GET("/GetDashboard", getDashboardItems)
		dashboardGroup.GET("/me", me)
		dashboardGroup.GET("/status", statusResp)
	}

	settingsGroup := r.Group("/settings")
	settingsGroup.Use(authMiddleware())
	{
		settingsGroup.GET("/")
		settingsGroup.GET("/GetUser", authMiddleware(), getUsers)
		settingsGroup.GET("/me", me)
		settingsGroup.GET("/status", statusResp)
	}

	hvacGroup := r.Group("/hvac")
	hvacGroup.Use(authMiddleware())
	{
		hvacGroup.GET("/", GetHVAC)
		hvacGroup.POST("/updateHVAC", updateThermostat)
		hvacGroup.GET("/GetHVAC", GetHVAC)
		hvacGroup.GET("/status", statusResp)
	}

	networkingGroup := r.Group("/networking")
	networkingGroup.Use(authMiddleware())
	{
		networkingGroup.GET("/", me)
		networkingGroup.GET("/GetNetLogs", GetNetLogs)
		networkingGroup.GET("/GetNetPcapLogs", GetNetPcapLogs)
		networkingGroup.GET("/status", statusResp)
	}

	securityGroup := r.Group("/security")
	securityGroup.Use(authMiddleware())
	{
		securityGroup.GET("/", me)
		securityGroup.GET("/GetSecurity", GetSecurity)
		securityGroup.POST("/system", VerifySecurityCode)
		securityGroup.GET("/status", statusResp)
	}

	appliancesGroup := r.Group("/appliances", dashboardHandler)
	appliancesGroup.Use(authMiddleware())
	{
		appliancesGroup.GET("/", me)
		appliancesGroup.GET("/me", me)
		appliancesGroup.GET("/status", statusResp)
	}

	energyGroup := r.Group("/energy")
	appliancesGroup.Use(authMiddleware())
	{
		energyGroup.GET("/", me)
		energyGroup.POST("/GetEnergy", getAppliancesData)
		energyGroup.GET("/status", statusResp)
	}
	go func() {
		err := r.Run(":8081")
		if err != nil {

		}
	}()
	//if err != nil {
	//	log.Fatal("Server startup error:", err)
	//}
	// Create a channel to receive the missingPi array from BlockReceiver
	//

	select {}
}

func getUsers(c *gin.Context) {

	//session := sessions.Default(c)
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return
	}

	if username == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in session"})
		return
	}

	// Convert username to string
	usernameStr, ok := username.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username type"})
		return
	}

	// Fetch user by username from MongoDB
	fetchedUser, err := dal.FetchUser(client, usernameStr)
	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}
	fmt.Println("fetchedUser", fetchedUser)

	// Create a struct to hold the user information
	type userInfo struct {
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Role      string `json:"role"`
	}
	var userData userInfo

	userData.Username = fetchedUser.Username
	userData.FirstName = fetchedUser.FirstName
	userData.LastName = fetchedUser.LastName
	userData.Role = fetchedUser.Role
	fmt.Printf("userData: %+v\n", userData)

	c.JSON(http.StatusOK, userData)
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

func updateThermostat(c *gin.Context) {
	var req dal.ThermMessagingStruct
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println([]byte(req.UUID), req.Name, req.AppType, req.Function, req.Change)
	//fmt.Printf("\n%T\n", req.Change)
	dal.UpdateThermMessaging(client, []byte(req.UUID), req.Name, req.AppType, req.Function, req.Change)

}

func VerifySecurityCode(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var req struct {
		Code   string `json:"code"`
		Status []byte `json:"status"`
	}
	fmt.Println("Alarm Status: ", req.Status)
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	securityKey := os.Getenv("SECURITY_KEY")
	if securityKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Security key not found"})
		return
	}

	if req.Code == securityKey {
		c.JSON(http.StatusOK, gin.H{"message": "Security code verified"})
		fmt.Println("Key is valid ", req.Code)
		dal.UpdateMessaging(client, []byte("502857"), "Security", "SecuritySystem", "Status", string(req.Status))
		dal.UpdateMessaging(client, []byte("502858"), "Security", "SecuritySystem", "Status", string(req.Status))
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid security code"})
	}
}

func GetHVAC(c *gin.Context) {

	//fmt.Printf("Room name: %s\n", room)
	hvacs, err := dal.FetchHVAC(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("hvacs: ", hvacs)
	c.JSON(http.StatusOK, hvacs)

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

func getDashboardItems(c *gin.Context) {
	// Fetch the user's account type from the session
	accountType := c.GetString("accountType")
	if accountType == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Fetch the smart home data from the database
	smartHomeData, err := dal.FetchCollections(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch smart home data"})
		return
	}

	// Prepare the response data
	dashboardData := make(map[string]interface{})

	// Extract the relevant data for each category
	dashboardData["hvac"] = smartHomeData.HVAC
	dashboardData["lighting"] = smartHomeData.Lighting
	dashboardData["securitySystem"] = smartHomeData.SecuritySystem
	dashboardData["solarPanel"] = smartHomeData.SolarPanel

	// Extract appliance data
	applianceData := dal.Appliances{
		Dishwasher: smartHomeData.Dishwasher,
		Fridge:     smartHomeData.Fridge,
		Toaster:    smartHomeData.Toaster,
		Lighting:   smartHomeData.Lighting,
		Microwave:  smartHomeData.Microwave,
		Oven:       smartHomeData.Oven,
	}
	dashboardData["appliances"] = applianceData

	c.JSON(http.StatusOK, gin.H{"data": dashboardData})
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
	} else {
		session.Set("username", loginData.Username)
		//session.Save()
	}
	//c.Set("smartHomeDB", smartHomeDB)

	// Print smart home data
	//smartHomeData := dal.PrintSmartHomeDBContents(smartHomeDB)

	// JWT token creation
	claims := jwt.MapClaims{
		"username": fetchedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the token"})
		return
	}
	// Set the username in the session
	//session.Set("username", loginData.Username)
	session.Save()

	//fmt.Println(session.Get("username"))
	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString,
		"username": fetchedUser.Username,
		"password": fetchedUser.Password})
	fmt.Println(session.Get("username"))

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

func signupHandler(c *gin.Context) {
	var signupData struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Password  string `json:"password"`
		Username  string `json:"username"`
	}

	if err := c.BindJSON(&signupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	}

	//existingUser, err := dal.FetchUser(client, signupData.Username)
	//if err == nil {
	//	c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
	//	return
	//}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := &dal.User{
		Username:  signupData.Username,
		Password:  string(hashedPassword),
		FirstName: signupData.Firstname,
		LastName:  signupData.Lastname,
		Role:      "user",
	}

	err = dal.CreateUser(client, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created succesfully"})

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
		// Check for JWT token in the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			// Extract the token from the Authorization header
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString := parts[1]

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method")
					}
					return jwtKey, nil
				})

				if err == nil && token.Valid {
					if claims, ok := token.Claims.(jwt.MapClaims); ok {
						// Set the user claims in the request context
						fmt.Printf("Authorized")
						c.Set("user", claims)
						c.Next()
						return
					}
				}
			}
		}

		session := sessions.Default(c)
		username := session.Get("username")

		fmt.Printf("Session: %+v\n", session)
		fmt.Printf("Username: %+v\n", username)

		if username == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
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
	logs, err := dal.FetchLogging(client, "smartHomeDB")
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func GetNetPcapLogs(c *gin.Context) {
	logs, err :=  networktraffic.GetNetEvents() //.FetchLogging(client, "smartHomeDB")
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func GetSecurity(c *gin.Context) {
	security, err := dal.FetchSecurity(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, security)
}
