package main

import (
	"bytes"
	"encoding/json"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tirprox/sync/model/dreamwhite"
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

func main() {
	dreamwhite.Import()

}
