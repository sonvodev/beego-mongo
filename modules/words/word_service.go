package words

import (
	"context"
	"log"
	"wordbook/modules"

	"gopkg.in/mgo.v2/bson"
)

type WordService struct {
}

var (
	super *modules.AppBaseService = &modules.AppBaseService{CollectionName: "words"}
)

func init() {
	super.GetCollection()
}

func (this *WordService) GetListWords(filter WordParameter) ([]*WordDocument, int64, error) {

	condition := bson.M{
	}

	cur, count, resErr := super.GetList(condition, filter.BaseParameterModel.Index, filter.BaseParameterModel.Size)
	defer cur.Close(context.TODO())

	
	if resErr != nil {
		return nil, 0, resErr
	}

	var docs []*WordDocument
	for cur.Next(context.TODO()) {
		var u WordDocument
		resErr = cur.Decode(&u)
		if resErr != nil {
			log.Fatal("An exception occured on decoding the user document", resErr)
		}

		docs = append(docs, &u)
	}

	return docs, count, nil
}
