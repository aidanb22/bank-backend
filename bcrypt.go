package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	originalPassword := "aidan2113"

	// 1) Hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v\n", err)
	}
	fmt.Printf("Original plain-text: %s\n", originalPassword)
	fmt.Printf("Generated hash: %s\n", hashed)

	// 2) Compare using the same plain text
	err = bcrypt.CompareHashAndPassword(hashed, []byte(originalPassword))
	if err != nil {
		fmt.Println("Comparison with correct password => mismatch!", err)
	} else {
		fmt.Println("Comparison with correct password => success!")
	}

	// 3) Compare using a wrong password
	wrongPass := "thisIsWrong"
	err = bcrypt.CompareHashAndPassword(hashed, []byte(wrongPass))
	if err != nil {
		fmt.Printf("Comparison with wrong password => mismatch (expected)\n")
	} else {
		fmt.Printf("Comparison with wrong password => success?! (unexpected)\n")
	}
}
