package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/bsoncodec"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
	"github.com/tirprox/sync/httpclient"
	"github.com/tirprox/sync/model/moysklad"
	"log"
	"os"
)

func WriteStructToFileJSON(v interface{}, filename string) {

	buf := make([]byte, 0)
	out := bytes.NewBuffer(buf)

	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)

	err := enc.Encode(&v)

	newFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	out.WriteTo(newFile)

	defer newFile.Close()
}

var Collection *mongo.Collection

const DATABASE = "go"
const COLLECTION = "product"

func main() {

	var Client, err = mongo.NewClient("mongodb://admin:6h8s4ksoq@localhost:27017")

	if err != nil {
		log.Fatal("Could not start MongoClient: ", err)
	}

	Collection = Client.Database(DATABASE).Collection(COLLECTION)

	err = Client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//Mock group

	palto := moysklad.Group{}
	palto.Name = "Мужские пальто"

	//

	url := moysklad.API_BASE + "assortment/" + "?limit=100"

	responses := httpclient.GetAll(url)

	for _, response := range responses {

		assortment := moysklad.DecodeAssortment(response.Body)

		for _, row := range assortment.Rows {

			rowJson, err := json.Marshal(&row)
			if err != nil {
				log.Fatal(err)
			}

			switch row.Meta.Type {
			case "product":
				product := moysklad.Product{}
				json.Unmarshal(rowJson, &product)
				if err != nil {
					log.Fatal("Product: MongoDB write failed: " + err.Error())
					return
				}

				palto.Products = append(palto.Products, product)
				fmt.Println(product.Name)

			case "variant":
				variant := moysklad.Variant{}
				json.Unmarshal(rowJson, &variant)
				if err != nil {
					log.Fatal("Variant: MongoDB write failed: " + err.Error())
					return
				}
				fmt.Println(variant.Name)
			}
		}
	}

	b, err := bsoncodec.Marshal(&palto)
	if err != nil {
		log.Fatal("BSON Marshal Error: ", err)
	}

	d, err := bson.ReadDocument(b)

	if err != nil {
		log.Fatal("BSON Marshal Error: ", err)
	}

	update := bson.NewDocument(bson.EC.SubDocument("$set", d))

	filter := bson.NewDocument(bson.EC.String("name", palto.Name))

	_, err = Collection.UpdateOne(context.Background(), filter, update, updateopt.Upsert(true))

	if err != nil {
		log.Fatal("Product: MongoDB write failed: " + err.Error())
		return
	}
}
