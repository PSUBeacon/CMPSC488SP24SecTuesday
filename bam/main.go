package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var users = []User{
	{ID: 1, Username: "user1", Password: "pass1"},
}

var secretKey = []byte("your_secret_key")

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")

	port := 3002
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var userCredentials User
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := findUser(userCredentials.Username, userCredentials.Password)
	if user != nil {
		token := generateToken(user.ID)
		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func findUser(username, password string) *User {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return &u
		}
	}
	return nil
}

func generateToken(userID int) string {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}

	return signedToken
}
