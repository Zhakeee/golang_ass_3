package pkg

import (
	"fmt"
	"net/http"
	"strconv"

	md "github.com/LeilaBeken/golang_ass_3/models"
	"github.com/gin-gonic/gin"
)

func listProducts(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	products := []md.Book{}
	// Read the "sort_by" and "sort_order" query parameters
	sortBy := c.Query("sort_by")
    sortOrder := c.Query("sort_order")
	searchQuery := c.Query("search")
	if searchQuery != "" {
		db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", searchQuery)).Find(&products)
	} else {
		db.Find(&products)
	}
    if sortBy != "" && sortOrder != "" {
		query := db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
		query.Find(&products)		
	}

	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	productID := c.Param("id")
	product := md.Book{}
	db.First(&product, productID)
	if product.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	var productData md.Book
	if err := c.ShouldBindJSON(&productData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&productData)
	c.JSON(http.StatusOK, productData)
}

func deleteProduct(c *gin.Context) {
	db, err := GetDB()
	if err != nil {panic(err)}
	productID := c.Param("id")
	existingProduct := md.Book{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.Delete(&existingProduct)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func updateProduct(c *gin.Context) {
	var product md.Book
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	productID := c.Param("id")

	// Find the product in the database by ID
	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update the product with the new values
	if c.PostForm("title") != "" {
		product.Title = c.PostForm("title")
	}
	if c.PostForm("author") != "" {
		product.Description = c.PostForm("author")
	}
	if c.PostForm("price") != "" {
		newPrice, err := strconv.Atoi(c.PostForm("price"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
			return
		}
		product.Price = newPrice
	}

	// Save the changes to the database
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, product)
}
