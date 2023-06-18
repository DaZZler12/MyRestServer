package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID        primitive.ObjectID `json:"id"`
	Brand     string             `json:"brand" bson:"brand"`
	Model     string             `json:"model" bson:"model"`
	Item_Name string             `json:"item_name" bson:"item_name"`
	Year      int64              `json:"year" bson:"year"`
	Price     float64            `json:"price" bson:"price"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type User struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Country  string             `json:"country"`
	Password string             `json:"password"`
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
