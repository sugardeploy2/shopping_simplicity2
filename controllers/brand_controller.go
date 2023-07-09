package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"shop_khordad/config"
	"shop_khordad/models/entities"
	"shop_khordad/models/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BrandController struct{}

type CreateBrandRequest struct {
	BrandName  string `json:"brand_name" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

func (bc *BrandController) CreateBrand(c *gin.Context) {
	var brandRequest CreateBrandRequest
	if err := c.ShouldBindJSON(&brandRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to load config"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to connect to database"})
		return
	}
	brandRepository := repositories.NewBrandRepository(db)

	newBrand := entities.Brand{
		BrandName:  brandRequest.BrandName,
		CategoryID: brandRequest.CategoryID,
	}

	err = brandRepository.CreateBrand(&newBrand)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to create new brand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand created successfuly", "brand": newBrand})

}

func (bc *BrandController) GetBrandByID(c *gin.Context) {
	brandID := c.Param("id")

	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}

	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to load config"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to connect to database"})
		return
	}

	brandRepository := repositories.NewBrandRepository(db)

	brand, err := brandRepository.GetBrandByID(brandIDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Brand couldnt found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"brand": brand})

}

func (bc *BrandController) UpdateBrand(c *gin.Context) {
	var updateRequest CreateBrandRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	brandID := c.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}
	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to load config"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to connect to database"})
		return
	}

	brandRepository := repositories.NewBrandRepository(db)

	brand, err := brandRepository.GetBrandByID(brandIDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Brand couldnt found"})
		return
	}
	brand.BrandName = updateRequest.BrandName
	brand.CategoryID = updateRequest.CategoryID

	err = brandRepository.UpdateBrand(brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to update brand"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update successfully", "brand": brand})

}

func (bc *BrandController) DeleteBrand(c *gin.Context) {
	brandID := c.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}
	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to load config"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to connect to database"})
		return
	}

	brandRepository := repositories.NewBrandRepository(db)

	brand, err := brandRepository.GetBrandByID(brandIDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Brand couldnt found"})
		return
	}

	err = brandRepository.DeleteBrand(brand)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Faild to delete brand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "brand deleted successfuly"})

}
