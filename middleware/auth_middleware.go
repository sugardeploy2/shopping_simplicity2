package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

// UserClaims represents the custom claims for JWT authentication
type UserClaims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization header
		authHeader := c.GetHeader("Authorization")

		// Check if the header is present
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Provide the secret key or public key to verify the token's signature
			return []byte("your-secret-key"), nil
		})

		// Check for parsing or verification errors
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Get the claims from the token
		claims, ok := token.Claims.(*UserClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// You can now access the user ID, email, or any other claims as needed
		userID := claims.UserID
		email := claims.Email
		username := claims.Username

		// Set the user ID and email in the request context for downstream handlers
		c.Set("userID", userID)
		c.Set("email", email)
		c.Set("username", username)

		// Continue to the next middleware or handler
		c.Next()
	}
}
func GoogleSignInMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the ID token from the request
		idToken := c.PostForm("id_token")

		// Create a new OAuth2 config for Google Sign-In
		oauthConfig := &oauth2.Config{
			ClientID: "your-client-id",
			// Set the appropriate RedirectURL, Scopes, and other config options
			RedirectURL: "your-redirect-url",
			Scopes:      []string{"email", "profile"},
			Endpoint:    google.Endpoint,
		}

		// Verify the ID token using the OAuth2 config
		payload, err := idtoken.Validate(c.Request.Context(), idToken, oauthConfig.ClientID)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			c.Abort()
			return

		}

		// Get the user ID, email, and name from the token
		userID := payload.Subject
		email := payload.Claims["email"].(string)
		name := payload.Claims["name"].(string)

		// Check if the user exists in your system
		// Perform any necessary actions, such as creating a new user account or retrieving the user's information

		// Set the user ID, email, and name in the request context for downstream handlers
		c.Set("userID", userID)
		c.Set("email", email)
		c.Set("name", name)

		// Continue to the next middleware or handler
		c.Next()
	}
}
