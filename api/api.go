package api

import (
	_ "e_commerce/api/docs"
	v1 "e_commerce/api/handler/v1"
	"e_commerce/config"
	"e_commerce/storage"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Config  config.Config
	Storage storage.StorageI
}

func New(ro RouterOptions) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage: ro.Storage,
		Cfg:     ro.Config,
	})
	// ! Category endpoints
	router.GET("/v1/product", handlerV1.GetAllProducts)
	router.GET("/v1/product/:product_id", handlerV1.GetProduct)
	router.POST("/v1/product", handlerV1.CreateProduct)
	router.PUT("/v1/product/:product_id", handlerV1.UpdateProduct)

	// Auth endpoints
	router.POST("/users/signup", handlerV1.SignUp)
	router.POST("/users/login", handlerV1.Login)

	// Admin endpoints
	router.POST("/admin/addproduct", handlerV1.Addproduct)

	// Customer endpoints
	router.GET("/users/getproductall", handlerV1.GetproductAll)
	router.GET("/users/getproductbyid", handlerV1.GetproductById)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
