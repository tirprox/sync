package dreamwhite

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/bsoncodec"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
	"github.com/tirprox/sync/model/moysklad"
	"github.com/tirprox/sync/mongoclient"
	"log"
)

var Collection *mongo.Collection

func Import() {
	initMongoConnection()

	//TODO Requesting and filtering productFolders (groups)

	folders := moysklad.GetProductFolders()
	for _, folder := range folders {

		b, err := bsoncodec.Marshal(&folder)
		if err != nil {
			log.Fatal("BSON Marshal Error: ", err)
		}

		d, err := bson.ReadDocument(b)

		if err != nil {
			log.Fatal("BSON Marshal Error: ", err)
		}

		update := bson.NewDocument(bson.EC.SubDocument("$set", d))

		filter := bson.NewDocument(bson.EC.String("name", folder.Name))

		_, err = Collection.UpdateOne(context.Background(), filter, update, updateopt.Upsert(true))
		//fmt.Println(folder.ID)
	}

	//TODO Requesting stores for cities from config

	//TODO Requesting images from static

	//TODO Preparing assortment queries

	//TODO Executing queries and collecting assortment

	//TODO Parsing assortment and pushing it to a MongoDB

}

func initMongoConnection() {
	const DATABASE = "go"
	const COLLECTION = "product"

	Collection = mongoclient.Client.Database(DATABASE).Collection(COLLECTION)

	if err := mongoclient.Client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
