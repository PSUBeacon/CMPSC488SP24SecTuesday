package main

import (
	"CMPSC488SP24SecTuesday/dal"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	//"github.com/joho/godotenv"

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

	//r.GET("/lighting", updateLighting)
	//Messaging stuff

	// use JWT middleware for all protected routes
	r.Use(authMiddleware())

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
		userGroup.GET("/lighting", updateLighting)
		// Add more user-only routes as needed
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
	if err := c.ShouldBindJSON(&req); err != nil {
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

	// Connect to MongoDB
	client, err := dal.ConnectToMongoDB()
	//fmt.Printf("This is the client: ", client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to MongoDB"})
		return
	}
	defer client.Disconnect(context.Background())

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
	c.JSON(http.StatusOK, gin.H{"username": userProfile.Username, "email": userProfile.Email})
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
		Username: username,
		Email:    "user@example.com",
	}
}

// struct to represent user profile data
type UserProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
