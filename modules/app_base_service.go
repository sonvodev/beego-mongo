package modules

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AppBaseService struct {
		Collection *mongo.Collection
	}

	FindOptions struct {
		Limit *int64
	}
)

func (this *AppBaseService) GetList(filter interface{}, index int64, size int64) (*mongo.Cursor, int64, error) {

	limit := int64(size)
	skip := int64(index * size)
	findOptions := options.FindOptions{Limit: &limit, Skip: &skip}

	cur, resErr := this.Collection.Find(context.TODO(), filter, &findOptions)

	if resErr != nil {
		return nil, 0, resErr
	}

	count, resErr := this.Collection.CountDocuments(context.TODO(), filter)
	if resErr != nil {
		return nil, 0, resErr
	}

	return cur, count, nil
}

func (this *AppBaseService) InsertOne(doc interface{}) (*mongo.InsertOneResult, error) {

	cur, err := this.Collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

func (this *AppBaseService) InsertMany(doc []interface{}) (*mongo.InsertManyResult, error) {

	cur, err := this.Collection.InsertMany(context.TODO(), doc)
	if err != nil {
		return nil, err
	}

	return cur, nil
}
