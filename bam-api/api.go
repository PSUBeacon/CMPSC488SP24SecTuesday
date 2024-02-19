package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    //"github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)
import "github.com/gin-contrib/cors"


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
        // Add more user-only routes as needed
    }

    

    err := r.Run(":8080")
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
    //  authentication logic goes here
    // if authentication is successful, create a JWT token.

    // load .env file which is in gitignore
    //err := godotenv.Load(".env")
    //if err != nil {
    //  log.Fatal("Error loading .env file")
    //}

    //tmpUsername := os.Getenv("TMP_USERNAME")      // get this from mongodb
    //tmpPasswordHash := os.Getenv("TMP_PASS_HASH") // get this from mongodb

    tmpUsername := "beaconuser"
    // Use the bcrypt hash of the password here, not the plaintext
    tmpPasswordHash := "$2a$10$oxa8KNIQRDEmPhadO7PQrecW.Dbl2KJi/7cN0HXxHyAPBiVMGFhaG"

    fmt.Printf("tmpUsername: %s\n", tmpUsername)
    fmt.Printf("tmpPasswordHash: %s\n", tmpPasswordHash)

    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    // BindJSON will return an error if the JSON is invalid
    if err := c.BindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if loginData.Username != tmpUsername {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // CompareHashAndPassword will return an error if the password does not match the hash
    if err := bcrypt.CompareHashAndPassword([]byte(tmpPasswordHash), []byte(loginData.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // jwt token creation:
    claims := jwt.MapClaims{
        "username": loginData.Username,                    // replace with matching value from mongoDB users table
        "role":     "admin",                               // temporary role
        "exp":      time.Now().Add(time.Hour * 24).Unix(), // token expiration time
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

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