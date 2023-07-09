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

type CategoryController struct{}

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var createRequest CreateCategoryRequest

	if err := c.ShouldBindQuery(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant load config"})
		return
	}

	dsn := fmt.Sprint("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant connet to database"})
		return
	}
	categoryRepository := repositories.NewCategoryRepository(db)

	newCategory := entities.Category{
		CategoryName: createRequest.CategoryName,
	}
	err = categoryRepository.CreateCategory(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant make category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category created successfully", "category": newCategory})
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var updateRequest CreateCategoryRequest
	categoryID := c.Param("id")

	categoryIDInt, err := strconv.Atoi(categoryID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant load config"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to databse"})
		return

	}
	categoryRepository := repositories.NewCategoryRepository(db)

	category, err := categoryRepository.GetCategoryByID(categoryIDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to find category"})
		return
	}
	category.CategoryName = updateRequest.CategoryName
	c.JSON(http.StatusOK, gin.H{"message": "category updated successfully", "category": category})

}

func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	categoryID := c.Param("id")

	categoryIDInt, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
	}
	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant load config"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%d", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to database"})
		return
	}
	categoryRepository := repositories.NewCategoryRepository(db)

	category, err := categoryRepository.GetCategoryByID(categoryIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant find category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func (cc *CartController) DeletCategory(c *gin.Context) {
	categoryID := c.Param("id")

	categoryIDInt, err := strconv.Atoi(categoryID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id "})
		return
	}
	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant load config"})
		return
	}

	dsn := fmt.Sprint("%s:%s@tcp(%s/%d):%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to daqtabase"})
		return
	}
	categoryRepository := repositories.NewCategoryRepository(db)

	category, err := categoryRepository.GetCategoryByID(categoryIDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant find category"})
		return
	}
	err = categoryRepository.DeleteCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "cart deleted successfully"})

}
