package controllers

import (
	"goEcart/db"
	"goEcart/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditBrand() gin.HandlerFunc {
	return func(c *gin.Context) {
		bid := c.Param("id")
		id, err := strconv.Atoi(bid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in conevrting id"})
			return
		}

		var editBrand models.Brand
		if err := c.ShouldBindJSON(&editBrand); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid body request"})
			return
		}

		if err := db.DB.Model(&models.Brand{}).Where("id = ?", id).Update("brand_name", editBrand.BrandName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating brand name"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Brand updated succesfully"})
	}
}

func GetBrand() gin.HandlerFunc {
	return func(c *gin.Context) {
		var brandData []models.Brand
		if err := db.DB.Find(&brandData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get brands"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Brands retrive succesfully",
			"data":    brandData,
		})
	}
}
