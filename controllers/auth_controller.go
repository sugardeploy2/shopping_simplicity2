package controllers

import (
	"fmt"
	"net/http"

	"time"

	"shop_khordad/config"
	"shop_khordad/models/repositories"

	"github.com/dgrijalva/jwt-go"

	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication related requests
type AuthController struct{}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles the login request
func Login(c *gin.Context) {
	var loginRequest LoginRequest

	// Bind the request body to the login request struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to load config"})
		return

	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Server.Port, config.Database.Name)

	db, err := gorm.Open(mysql.open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "faild to connect to database"})
		return
	}
	// Verify the credentials against the user repository or authentication service
	userRepository := repositories.NewUserRepository(db) // Replace `db` with your actual DB connection
	user, err := userRepository.VerifyCredential(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token if the credentials are valid
	tokenString, err := generateToken(strconv.Itoa(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token as the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Register handles the registration request
func (ac *AuthController) Register(c *gin.Context) {
	// Handle the registration logic here
	// Retrieve the user details from the request body
	// Create a new user in the database or user repository
	// Generate a JWT token for the registered user
	// Return the token as the response

	// Example response:
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Profile handles the user profile request
func (ac *AuthController) Profile(c *gin.Context) {
	// Handle the user profile logic here
	// Retrieve the user ID from the request context or JWT token
	// Fetch the user details from the database or user repository
	// Return the user profile information as the response

	// Example response:
	c.JSON(http.StatusOK, gin.H{"message": "User profile"})
}

func generateToken(userID string) (string, error) {
	// Create the claims for the JWT token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	// Replace "your-secret-key" with your own secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
