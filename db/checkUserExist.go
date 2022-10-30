package db

import (
	"context"
	"time"

	"github.com/rodrigoSolano/twitter_go/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExist is a function to check if a user already exist in the database
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("twitter")
	col := db.Collection("users")
	condition := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
