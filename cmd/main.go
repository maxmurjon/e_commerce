package main

import (
	"context"
	"fmt"

	"e_commerce/api"
	"e_commerce/config"
	"e_commerce/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	strg storage.StorageI
	cfg  config.Config
)

func initDependencies() {
	cfg = config.Load()

	credential := options.Credential{
		Username: cfg.MongoUser,
		Password: cfg.MongoPassword,
	}
	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort)

	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString).SetAuth(credential))

	if err != nil {
		fmt.Println("error to connect to mongo database")
	}
	connDB := mongoConn.Database("mongo-golang")

	strg = storage.NewProductStorage(connDB)
}

func main() {
	initDependencies()
	server := api.New(api.RouterOptions{
		Config:  cfg,
		Storage: strg,
	})

	err := server.Run(cfg.Port)

	if err != nil {
		fmt.Println("Something went wrong")
		panic(err)
	}

}
