package main

import (
	"CMPSC488SP24SecTuesday/dal"
	"CMPSC488SP24SecTuesday/networktraffic"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/JGLTechnologies/gin-rate-limit"
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

var username = ""

var r = gin.Default() // Configure CORS

func UpdateMissingPi(PiNum []string) {
	for i := 0; i < len(PiNum); i++ {
		PiNumInt, _ := strconv.Atoi(PiNum[i])
		disconnectedPiNums = append(disconnectedPiNums, PiNumInt)
	}
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
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
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 3,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// unprotected endpoints no auth needed
	r.GET("/status", statusResp)
	r.POST("/login", loginHandler)
	r.GET("/logout", logout)
	r.POST("/signup", signupHandler)
	r.DELETE("/users/:username", deleteUserHandler)
	r.POST("/users/:username/role", updateUserRoleHandler)
	// Apply JWT middleware to protected routes
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(authMiddleware())

	// This ones should be protected
	protectedRoutes.POST("/lighting", updateIoT)
	protectedRoutes.GET("/lighting", GetLights)

	protectedRoutes.POST("/security", updateIoT)
	protectedRoutes.GET("/security", GetSecurity)

	protectedRoutes.POST("/appliances", getAppliancesData)

	//protectedRoutes.POST("/energy", updateIoT)

	//ADJUSTMENT:
	// Combined route group for both admin and user dashboards
	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.Use()
	{
		dashboardGroup.GET("/")
		dashboardGroup.GET("/GetDashboard", getDashboardItems)
		dashboardGroup.GET("/me", me)
		dashboardGroup.GET("/status", statusResp)
	}

	settingsGroup := r.Group("/settings")
	settingsGroup.Use()
	{
		settingsGroup.GET("/")
		settingsGroup.GET("/GetUsers", getUsers)
		settingsGroup.GET("/me", me)
		settingsGroup.GET("/status", statusResp)
	}

	hvacGroup := r.Group("/hvac")
	hvacGroup.Use()
	{
		hvacGroup.GET("/", GetHVAC)
		hvacGroup.POST("/updateHVAC", mw, updateThermostat)
		hvacGroup.GET("/GetHVAC", GetHVAC)
		hvacGroup.GET("/status", statusResp)
	}

	networkingGroup := r.Group("/networking")
	networkingGroup.Use()
	{
		networkingGroup.GET("/", me)
		networkingGroup.GET("/GetNetLogs", GetNetLogs)
		networkingGroup.GET("/GetNetPcapLogs", GetNetPcapLogs)
		networkingGroup.GET("/status", statusResp)
	}

	securityGroup := r.Group("/security")
	securityGroup.Use()
	{
		securityGroup.GET("/", me)
		securityGroup.GET("/GetSecurity", GetSecurity)
		securityGroup.POST("/system", VerifySecurityCode)
		securityGroup.GET("/status", statusResp)
	}

	appliancesGroup := r.Group("/appliances")
	appliancesGroup.Use()
	{
		appliancesGroup.GET("/", me)
		appliancesGroup.GET("/me", me)
		appliancesGroup.GET("/status", statusResp)
	}

	energyGroup := r.Group("/energy")
	appliancesGroup.Use()
	{
		energyGroup.GET("/", me)
		energyGroup.POST("/updateEnergy", updateIoT)
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
	// Fetch all users from MongoDB
	fetchedUsers, err := dal.FetchAllUsers(client)
	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Create a struct to hold the user information
	type userInfo struct {
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Role      string `json:"role"`
	}

	var usersData []userInfo

	for _, user := range fetchedUsers {
		userData := userInfo{
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
		}
		usersData = append(usersData, userData)
	}

	c.JSON(http.StatusOK, usersData)
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

	if req.Change == "true" {
		req.Change = "false"
	} else {
		req.Change = "true"
	}

	fmt.Println(req)
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

	if req.Function == "Mode" {
		var req2 = req
		req2.Function = "Status"
		req2.Change = "true"

		dal.UpdateThermMessaging(client, []byte(req2.UUID), req2.Name, req2.AppType, req2.Function, req2.Change)
		time.Sleep(1 * time.Second)
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
	accountType, _ := c.Get("accountType")
	fmt.Println(accountType)
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
		Username  string `json:"username"`
		Password  string `json:"password"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Role      string `json:"role"`
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
		session.Set("accountType", fetchedUser.Role)
		username = loginData.Username
		//session.Save()
	}
	//c.Set("smartHomeDB", smartHomeDB)

	// Print smart home data
	//smartHomeData := dal.PrintSmartHomeDBContents(smartHomeDB)

	// JWT token creation
	claims := jwt.MapClaims{
		"username":  fetchedUser.Username,
		"firstname": fetchedUser.FirstName,
		"lastname":  fetchedUser.LastName,
		"role":      fetchedUser.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
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
		"username":  fetchedUser.Username,
		"password":  fetchedUser.Password,
		"firstname": fetchedUser.FirstName,
		"lastname":  fetchedUser.LastName,
		"role":      fetchedUser.Role,
	})

	fmt.Println(session.Get("username"))

}

func deleteUserHandler(c *gin.Context) {
	username := c.Param("username")

	err := dal.DeleteUser(client, username)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func updateUserRoleHandler(c *gin.Context) {
	username := c.Param("username")
	var roleUpdate struct {
		Role string `json:"role"`
	}
	if err := c.BindJSON(&roleUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := dal.UpdateUserRole(client, username, roleUpdate.Role)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
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
	username = ""
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

	//getUsers(c)
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
	logs, err := networktraffic.GetNetEvents() // Fetch the logs
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Since logs is already a JSON []byte, write it directly to the response body
	c.Data(http.StatusOK, "application/json", logs)
}

func GetSecurity(c *gin.Context) {
	security, err := dal.FetchSecurity(client, "smartHomeDB")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, security)
}
