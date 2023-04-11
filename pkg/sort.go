package pkg

import (
	"net/http"
	"fmt"
	md "github.com/LeilaBeken/golang_ass_3/models"
	"github.com/gin-gonic/gin"
)

func sortProducts(c *gin.Context) {
    // Retrieve the sorting criterion and order from the request
    sortBy := c.Query("sort_by")
    sortOrder := c.Query("sort_order")

    // Retrieve the products from the database
    products := []md.Book{}
	db, err := GetDB()
	if err != nil {panic(err)}
    query := db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
    query.Find(&products)

    // Return the products in the response
    c.JSON(http.StatusOK, products)
}