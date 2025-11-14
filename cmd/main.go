package main

import (
	"api-go/controller"
	"api-go/usecase"

	"github.com/gin-gonic/gin"
)

func main(){

	server := gin.Default()

	ProductUseCase := usecase.NewProductUseCase()
	//camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//cria a rota
	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}