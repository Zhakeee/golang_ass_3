package pkg

import (
	"github.com/gin-gonic/gin"
)

//defining routes
func Routes() {
	router := gin.Default()
	router.GET("/books", listProducts)
	router.GET("/books/:id", getProduct)
	router.POST("/books", createProduct)
	router.DELETE("/books/:id", deleteProduct)
	router.PUT("/books/:id", updateProduct)


	// Start the server
	router.Run(":8080")
}