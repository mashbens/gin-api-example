package main

import (
	"pustaka-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooksHandler)
	router.POST("/books", controllers.PostBookByIDHandler)
	router.GET("/books/:id", controllers.GetbooksByIDHandler)
	router.PUT("/books/:id", controllers.UpdateBookByIDHandler)
	router.DELETE("/books/:id", controllers.DelleteBookByIDHandler)

	router.Run(":8080")
}
