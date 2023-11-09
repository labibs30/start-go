package main

import (
	"github.com/tentangkode/go-restapi-gin/controllers/productController"
	"github.com/tentangkode/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("api/products", productController.Index)
	r.GET("api/product/:id", productController.Show)
	r.POST("api/product", productController.Create)
	r.PUT("api/product/:id", productController.Update)
	r.DELETE("api/product/:id", productController.Delete)

	r.Run()
}