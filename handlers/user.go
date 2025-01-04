package handlers

import (
	"bank-app/database"
	"bank-app/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// -------------------
// STRUCT & CONSTANTS
// -------------------

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password,omitempty" json:"password"`
}

// -------------------
// HELPER FUNCTIONS
// -------------------

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPasswordHash(password, hash string) bool {
	log.Printf("Comparing provided password: '%s' with stored hash: '%s'", password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf("Password comparison failed: %v", err)
		return false
	}
	return true
}

// -------------------
// REGISTER USER
// -------------------

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Log the plain-text password length for debugging
	log.Printf("Registering user with email='%s', raw password='%s' [len=%d]",
		user.Email, user.Password, len(user.Password))

	// Reject empty or missing passwords
	if len(user.Password) == 0 {
		http.Error(w, "Password cannot be empty", http.StatusBadRequest)
		return
	}

	// Hash the plain-text password exactly once
	hashedBytes, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	hashedPassword := string(hashedBytes)

	log.Printf("Hashed password for user %s: %s", user.Email, hashedPassword)
	user.Password = hashedPassword

	// Insert the user into MongoDB
	collection := database.GetCollection("bank", "users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	log.Printf("Inserted user: %+v", user)

	// Optional: Immediately query the DB to confirm the stored password
	var checkUser User
	err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&checkUser)
	if err != nil {
		log.Printf("Immediately after insert, can't find user? %v", err)
	} else {
		log.Printf("Found user in DB: email=%s, password=%s",
			checkUser.Email, checkUser.Password)
	}

	// Hide password in the response
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// -------------------
// LOGIN USER
// -------------------

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt for email: '%s', password: '%s'", creds.Email, creds.Password)

	// Get the user from MongoDB
	collection := database.GetCollection("bank", "users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err = collection.FindOne(ctx, bson.M{"email": creds.Email}).Decode(&user)
	if err != nil {
		log.Printf("Email not found: %s", creds.Email)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	log.Printf("User found: %s, hashedPassword in DB: %s", user.Email, user.Password)

	// Double-check that we do NOT hash user.Password here!
	// Just compare the plain-text from the request to the stored hash
	log.Printf("Manual test: comparing '%s' to DB password '%s'", creds.Password, user.Password)
	if !checkPasswordHash(creds.Password, user.Password) {
		log.Println("Password mismatch")
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT on success
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	log.Println("Login successful")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// -------------------
// PROTECTED ROUTE
// -------------------

func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "You have accessed a protected route"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
