package controller

import (
	"api-go/model"
	"api-go/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

//permite a função ser chamada
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

//cria o get products
func (p *productController) GetProducts(ctx *gin.Context){
	products := []model.Product{
		{
			ID: 1,
			Name: "Batata frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}