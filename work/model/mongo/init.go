package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
	err         error
)

func init() {

	once.Do(func() {
		mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second))
		if err != nil {
			panic(err)
		}
	})

}

func MongoClient() *mongo.Client {
	return mongoClient
}
