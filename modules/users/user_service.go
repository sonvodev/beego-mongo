package users

import (
	"log"
	"wordbook/modules"
	"wordbook/processors"

	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct {
}

var (
	super modules.AppBaseService = modules.AppBaseService{Collection: processors.GetCollection("users")}
)

func (this *UserService) GetListWords(param *UserParameter) ([]*UserDocument, int64, error) {

	condition := bson.M{}

	cur, count, resErr := super.GetList(condition, param.BaseParameterModel.Index-1, param.BaseParameterModel.Size)
	defer cur.Close(context.Background())

	if resErr != nil {
		return nil, 0, resErr
	}

	var docs []*UserDocument
	for cur.Next(context.TODO()) {
		var u UserDocument
		resErr = cur.Decode(&u)
		if resErr != nil {
			log.Fatal("An exception occured on decoding the user document", resErr)
		}

		docs = append(docs, &u)
	}

	return docs, count, nil
}
