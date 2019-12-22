package users

import (
	"context"
	"log"
	"wordbook/modules"

	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
}

var (
	super *modules.AppBaseService = &modules.AppBaseService{CollectionName: "users"}
)

func init() {
	super.GetCollection()
}

func (this *UserService) GetListWords(filter UserParameter) ([]*UserDocument, int64, error) {

	condition := bson.M{
	}

	cur, count, resErr := super.GetList(condition, filter.BaseParameterModel.Index, filter.BaseParameterModel.Size)

	
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
