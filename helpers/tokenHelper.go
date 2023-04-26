package helper

import (
	"github.com/feliciacia/golang_gin_mongodb/golang_gin_mongodb/database"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	email      string
	first_name string
	last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")
