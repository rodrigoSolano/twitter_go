package db

import (
	"context"
	"time"

	"github.com/rodrigoSolano/twitter_go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser is a function to insert a new user in the database
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("twitter")
	col := db.Collection("users")
	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
