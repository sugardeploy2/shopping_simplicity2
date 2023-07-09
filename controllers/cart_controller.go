package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"shop_khordad/config"
	"shop_khordad/models/entities"
	"shop_khordad/models/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CartController struct{}

type CreateCartRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

func (cc *CartController) CreateCart(c *gin.Context) {
	var createRequest CreateCartRequest

	cartID := c.Param("id")

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
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
	cartRepositroy := repositories.NewCartRepository(db)

	newCart := entities.Cart{
		UserID:       createRequest.UserID,
		CreationDate: time.Now(),
	}
	err = cartRepositroy.CreateCart(&newCart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "faild to create card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "cart created successfully", "cart": newCart})

}
func (cc *CartController) UpdateCart(c *gin.Context) {
	var updateRequest CreateCartRequest
	cartID := c.Param("id")

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cartIDInt, err := strconv.Atoi(cartID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}

}
