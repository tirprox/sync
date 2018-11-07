package dreamwhite

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tirprox/sync/mongoclient"
	"log"
)

var Collection *mongo.Collection

func Import() {
	initMongoConnection()

	//TODO Requesting and filtering productFolders (groups)

	//TODO Requesting stores for cities from config

	//TODO Requesting images from static

	//TODO Preparing assortment queries

	//TODO Executing queries and collecting assortment

	//TODO Parsing assortment and pushing it to a MongoDB

}

func initMongoConnection() {
	const DATABASE = "go"
	const COLLECTION = "product"

	/*if (err!= nil) {
		log.Fatal("Could not start MongoClient: ", err)
	}
	*/
	Collection = mongoclient.Client.Database(DATABASE).Collection(COLLECTION)

	err := mongoclient.Client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
