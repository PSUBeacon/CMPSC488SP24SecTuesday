package main

import (
	"CMPSC488SP24SecTuesday/dal"
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"

	//"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
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
	r.POST("/lighting", updateLighting)

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

type UpdateLightingRequest struct {
	UUID       string `json:"UUID"`
	Status     bool   `json:"Status"`
	Brightness int    `json:"Brightness"`
}

func updateLighting(c *gin.Context) {
	var req UpdateLightingRequest
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("Received request body: %s\n", string(requestBody))
	// Reset the request body to be able to parse it again
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Iotlighting function
	dal.Iotlighting([]byte(req.UUID), req.Status, req.Brightness)

	// Respond to the request indicating success.
	c.JSON(http.StatusOK, gin.H{"message": "Lighting updated successfully"})
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

	// Fetch user by username from MongoDB
	fetchedUser, err := dal.FetchUser(client, loginData.Username)
	fmt.Printf("Username:", fetchedUser.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // Use generic error message
		return
	}

	passwordFromDB := fetchedUser.Password
	fmt.Printf("Password: %s\n", passwordFromDB)

	// Compare the password hash using bcrypt.CompareHashAndPassword
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(loginData.Password))
	if err != nil {
		fmt.Printf("Error: password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // Use generic error message
		return
	}

	// JWT token creation
	claims := jwt.MapClaims{
		"username": fetchedUser.Username,
		"role":     fetchedUser.Role,                      // Use the actual role from the fetched user
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the token"})
		return
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
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

// Combined dashboard handler for both admin or user
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
			"message":     "Welcome to the Owner dashboard",
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
}
