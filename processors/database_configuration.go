package processors

import (
	"context"
	"log"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
	err    error
)

func init() {
	uri := beego.AppConfig.String("DBConnectionString")
	clientOptions := options.Client().ApplyURI(uri)

	client, err = mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
		log.Println("Connect to server failed")
	}

	log.Println("Connected to server !!!")
}

func GetCollection(collectionName string) *mongo.Collection {

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
		return nil
	} else {
		log.Println("Get collection succeeded !!!")
		return client.Database(beego.AppConfig.String("DBName")).Collection(collectionName)
	}
}
