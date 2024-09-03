package controller

import (
	"net/http"

	"github.com/SamirMamede/menufy/model"
	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:          1,
			Name:        "Batata frita Pequena",
			Price:       15.0,
			Description: "Batata para as crianças",
			ImagePath:   "/example.path",
		},
	}

	ctx.JSON(http.StatusOK, products)
}
