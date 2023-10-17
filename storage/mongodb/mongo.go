package mongodb

import (
	"context"
	"e_commerce/config"
	"e_commerce/storage"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db  *mongo.Collection
	product storage.ProductRepoI
}

func NewMongo(ctx context.Context, mongoUrl string) storage.StorageI {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Panic(err)
	}

	connDB:=client.Database("e_commerce")

	return &Store{
		db: connDB.Collection(config.CollectionName),
	}
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = &productRepo{collection: s.db}
	}
	return s.product
}
