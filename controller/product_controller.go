package controller

import "github.com/gin-gonic/gin"

type productController struct {
	//Usecase
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	//list products
}
