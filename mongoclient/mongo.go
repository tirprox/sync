package mongoclient

import "github.com/mongodb/mongo-go-driver/mongo"

var Client, err = mongo.NewClient("mongodb://admin:6h8s4ksoq@localhost:27017")
