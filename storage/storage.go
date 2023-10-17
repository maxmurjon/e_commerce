package storage

import "e_commerce/models"

type StorageI interface {
	Product() ProductRepoI
}

type ProductRepoI interface {
	Create(product *models.Product) (string, error)
	Get(id string) (*models.Product, error)
	GetAll(page, limit int64, name string) ([]*models.Product, int64, error)
}
