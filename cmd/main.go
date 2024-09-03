package main

import (
	"github.com/SamirMamede/menufy/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.Use(gin.Logger())

	ProductController := controller.NewProductController()

	server.GET("/teste", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "Successful call",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")

}
