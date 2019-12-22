package words

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WordDocument struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
}
