package controllers

import (
	"goEcart/db"
	"goEcart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBrand() gin.HandlerFunc {
	return func(c *gin.Context) {
		var brand models.Brand
		if err := c.ShouldBindJSON(&brand); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if brand with same name already exists or not */
		var existingBrand models.Brand
		if err := db.DB.Where("brand_name = ?", brand.BrandName).First(&existingBrand).Error; err == nil {
			c.JSON(http.StatusFound, gin.H{"error": "brand name already exists"})
			return
		}

		/* Save in database */
		if err := db.DB.Create(&brand).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in saving to the database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Brand added succesfully",
			"id":      brand.ID,
		})
	}
}

func AddCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if category with same name exists in db or not */
		var exisitingCategory models.Category
		if err := db.DB.Where("category_name = ?", category.CategoryName).First(&exisitingCategory).Error; err == nil {
			c.JSON(http.StatusFound, gin.H{"error": "category name already exists"})
			return
		}

		/* Save to the database */
		if err := db.DB.Create(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in saving to the database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Category added succesfully",
			"id":      category.ID,
		})
	}
}

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
			return
		}

		/* Check if product with same name already exists or not */
		var existingProduct models.Product
		if err := db.DB.Where("product_name = ?", product.ProductName).Find(&existingProduct).Error; err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Product already exists"})
			return
		}

		/* Save to the database */
		if err := db.DB.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in saving to the db"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product added sucesfully",
			"id":      product.ProductID,
		})
	}
}
