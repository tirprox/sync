package dreamwhite

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/bsoncodec"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
	"github.com/tirprox/sync/model/moysklad"
	"github.com/tirprox/sync/mongoclient"
	"log"
	"reflect"
)

var Collection *mongo.Collection

//TODO Optimize models with omitempty and bson tags

func Import() {
	fmt.Println("Connecting to MongoDB")
	initMongoConnection()

	//Requesting and filtering productFolders (groups)
	fmt.Println("Requesting product folders")
	folders := moysklad.GetProductFolders()

	fmt.Println("Writing product folders to DB")

	for _, folder := range folders {
		updateMongoDB(folder, folder.Name, "productFolder")
	}

	//Requesting stores for cities from config
	fmt.Println("Requesting stores")

	stores := moysklad.GetStores()
	fmt.Println("Writing stores to DB")

	for _, store := range stores {
		updateMongoDB(store, store.Name, "store")
	}

	//TODO Requesting images from static

	//Executing queries and collecting assortment
	fmt.Println("Requesting assortments")
	assortments := moysklad.GetAssortment(folders[12], stores[0])
	fmt.Println("Writing assortments to DB")

	rewriteCollection(toInterfaceSlice(assortments), "assortment")

	//Parsing assortment and pushing it to a MongoDB
	products := []moysklad.Product{}
	variants := []moysklad.Variant{}

	for _, assortment := range assortments {
		switch assortment.Meta.Type {

		case "product":
			product := moysklad.Product{}
			products = append(products, product)

		case "variant":
			variant := moysklad.Variant{}
			variants = append(variants, variant)
		}
	}

	rewriteCollection(toInterfaceSlice(products), "product")
	rewriteCollection(toInterfaceSlice(variants), "variant")

}

const DATABASE = "go"
const COLLECTION = "product"

func initMongoConnection() {

	Collection = mongoclient.Client.Database(DATABASE).Collection(COLLECTION)

	if err := mongoclient.Client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func updateMongoDB(data interface{}, filter string, collection string) {

	col := mongoclient.Client.Database(DATABASE).Collection(collection)

	b, err := bsoncodec.Marshal(data)
	if err != nil {
		log.Fatal("BSON Marshal Error: ", err)
	}

	d, err := bson.ReadDocument(b)

	if err != nil {
		log.Fatal("BSON Marshal Error: ", err)
	}

	update := bson.NewDocument(bson.EC.SubDocument("$set", d))

	filterDoc := bson.NewDocument(bson.EC.String("name", filter))

	_, err = col.UpdateOne(context.TODO(), filterDoc, update, updateopt.Upsert(true))
}

func toInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret

}

func rewriteCollection(data []interface{}, collection string) {
	col := mongoclient.Client.Database(DATABASE).Collection(collection)
	col.Drop(context.TODO())

	col.InsertMany(context.TODO(), data)
}
