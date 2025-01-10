package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Account represents a user's bank account.
type Account struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	AccountNumber string             `bson:"account_number" json:"account_number"`
	Balance       float64            `bson:"balance" json:"balance"`
	CreatedAt     int64              `bson:"created_at" json:"created_at"`
	UpdatedAt     int64              `bson:"updated_at" json:"updated_at"`
}
