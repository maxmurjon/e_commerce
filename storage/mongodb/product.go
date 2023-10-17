package mongodb

import (
	"context"
	"e_commerce/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct {
	collection *mongo.Collection
}

func (p *productRepo) Create(product *models.Product) (msg string, err error) {
	_, err = p.collection.InsertOne(
		context.Background(),
		product,
	)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}

func (p *productRepo) Get(id string) (*models.Product, error) {
	var product models.Product
	response := p.collection.FindOne(
		context.Background(),
		bson.M{
			"id": id,
		})

	err := response.Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) GetAll(page, limit int64, name string) (product []*models.Product, num int64, err error) {
	var (
		products []*models.Product
		filter   = bson.D{}
	)

	if name != "" {
		filter = append(filter, bson.E{Key: "name", Value: name})
	}

	opts := options.Find()

	skip := (page - 1) * limit
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	opts.SetSort(bson.M{
		"created_at": -1,
	})

	count, err := p.collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return nil, 0, err
	}

	rows, err := p.collection.Find(
		context.Background(),
		filter,
		opts,
	)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next(context.Background()) {
		var product models.Product

		err := rows.Decode(&product)

		if err != nil {
			return nil, 0, err
		}

		products = append(products, &product)

	}

	return products, count, nil
}
