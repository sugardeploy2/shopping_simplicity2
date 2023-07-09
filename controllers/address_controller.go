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

type AddressController struct{}

type CreateAddressRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	Street      string `json:"street" binding:"required"`
	City        string `json:"city" binding:"required"`
	PostalCode  string `json:"postal_code" binding:"required"`
	Country     string `json:"country" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
type UpdateAddressRequest struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone_number"`
}

func (ac *AddressController) CreateAdress(c *gin.Context) {
	var createRequest CreateAddressRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config, err := config.LoadConfig()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "couldnt load database"})
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "cant connect to database"})
		return
	}

	addressRepositry := repositories.NewAddressRepository(db)

	newAddress := entities.Address{
		UserID:      createRequest.UserID,
		Street:      createRequest.Street,
		City:        createRequest.City,
		PostalCode:  createRequest.PostalCode,
		Country:     createRequest.Country,
		PhoneNumber: createRequest.PhoneNumber,
	}

	err = addressRepositry.CreateAddress(&newAddress)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "cant create address"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "address created successfuly", "address": newAddress})

}
func (ac *AddressController) UpdateAddress(c *gin.Context) {
	addressID := c.Param("id")
	addressIDInt, err := strconv.Atoi(addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
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
	addressRepository := repositories.NewAddressRepository(db)

	address, err := addressRepository.GetAddressByID(addressIDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "address not found"})
	}

	var updateRequest UpdateAddressRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if updateRequest.Street != "" {
		address.Street = updateRequest.Street
	}

	if updateRequest.City != "" {
		address.City = updateRequest.City
	}
	if updateRequest.PhoneNumber != "" {
		address.PhoneNumber = updateRequest.PhoneNumber
	}
	if updateRequest.PostalCode != "" {
		address.PostalCode = updateRequest.PostalCode
	}
	if updateRequest.Country != "" {
		address.Country = updateRequest.Country
	}
	err = addressRepository.UpdateAddress(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to update address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "address updated successfully", "address": address})

}

func (ac *AddressController) DeleteAddress(c *gin.Context) {
	addressID := c.Param("id")
	addressIDInt, err := strconv.Atoi(addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}
	config, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to load config"})
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "faild to connect to database"})
		return
	}

	addressRepository := repositories.NewAddressRepository(db)

	address, err := addressRepository.GetAddressByID(addressIDInt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	err = addressRepository.DeleteAddress(address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
