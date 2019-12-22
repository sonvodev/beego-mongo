package modules

import (
	"log"
	"context"
	"wordbook/processors"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
)

type AppBaseService struct {
	CollectionName string
}

var (
	collection *mongo.Collection
)

func (this *AppBaseService) GetCollection() {
	collection = processors.GetCollection(this.CollectionName)
}

func (this *AppBaseService) GetList(filter interface{}, index int64, size int64) (*mongo.Cursor, int64, error) {

	log.Println("Quering on ", this.CollectionName)
	res, resErr := collection.Find(context.TODO(), filter)
	if resErr != nil {
		return nil, 0, resErr
	}

	count, cerr := collection.CountDocuments(context.TODO(),filter)
	if cerr != nil {
		return nil, 0, cerr
	}
	return res, count, nil
}
