package service

// import (
// 	"context"
// 	"fmt"

// 	"e_commerce/api/models"
// 	"e_commerce/config"
// 	"e_commerce/storage"
// )

// type productService struct {
// 	storage storage.StorageI
// 	cfg     config.Config
// }

// func NewService(strg storage.StorageI, cfg config.Config) *productService {
// 	return &productService{
// 		storage: strg,
// 		cfg:     cfg,
// 	}
// }

// func (ps *productService) Create(ctx context.Context, req *models.Product) (*models.CreateResponse, error) {
// 	ID, err := ps.storage.Product().Create(req)

// 	if err != nil {
// 		fmt.Println("error while creating product")
// 		return nil, err
// 	}
// 	return &models.CreateResponse{
// 		ID: ID,
// 	}, nil
// }

// func (ps *productService) Get(ctx context.Context, req *models.GetRequest) (*models.GetProductResponse, error) {

// 	response, err := ps.storage.Product().Get(req.ID)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.GetProductResponse{
// 		Product: response,
// 	}, nil
// }

// func (ps *productService) GetAll(ctx context.Context, req *models.GetAllProductsRequest) (*models.GetAllProductsResponse, error) {

// 	response, count, err := ps.storage.Product().GetAll(req.Page, req.Limit, req.Name)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.GetAllProductsResponse{
// 		Count:    count,
// 		Products: response,
// 	}, nil
// }
