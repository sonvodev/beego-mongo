package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname string             `json:"firstname" bson:"firstname"`
	Lastname  string             `json:"lastname" bson:"lastname"`
	Gender    string             `json:"gender" bson:"gender"`
	Email    string             `json:"email" bson:"email"`
}
