package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ResponseListingModel struct {
	cur *mongo.Cursor
	err error
}
