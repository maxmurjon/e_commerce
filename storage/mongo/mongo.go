package mongo

import (
	"context"

	"e_commerce/api/models"
	"e_commerce/config"
	"e_commerce/storage/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct { // Struct hosil qilmoqdamiz
	collection *mongo.Collection
}

func NewProductRepo(db *mongo.Database) repo.ProductStorageI { // Struct dan object hosil qivommiza
	return &productRepo{
		collection: db.Collection(config.CollectionName)}
}

func (pr *productRepo) Create(product *models.Product) (string, error) {
	_, err := pr.collection.InsertOne(
		context.Background(),
		product,
	)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}

func (pr *productRepo) Get(id string) (*models.Product, error) {
	var product models.Product
	response := pr.collection.FindOne(
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

func (pr *productRepo) GetAll(page, limit int64, name string) ([]*models.Product, int64, error) {

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

	count, err := pr.collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return nil, 0, err
	}

	rows, err := pr.collection.Find(
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
