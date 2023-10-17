package main

import (
	"context"
	"fmt"

	"e_commerce/config"
	"e_commerce/storage"
	"e_commerce/storage/mongodb"
)

var (
	strg storage.StorageI
	cfg  config.Config
)

func main() {
	cfg = config.Load()

	mongoString := fmt.Sprintf("mongodb://%s:%d", cfg.MongoHost, cfg.MongoPort)
	strg := mongodb.NewMongo(context.Background(), mongoString)

	// fmt.Println(strg.Product().Create(&models.Product{ID: "12343"}))
	fmt.Println(strg.Product().Get("12343"))
}
