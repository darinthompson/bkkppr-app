package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/darinthompson/bkkppr-app/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Try to get the token from the Authorization header first
	tokenString := c.GetHeader("Authorization")

	// If the token is not in the header, check for the cookie
	if tokenString == "" {
		var err error
		tokenString, err = c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := os.Getenv("SECRET_KEY")
		if secretKey == "" {
			log.Println("ERROR: SECRET_KEY is not set")
			return nil, fmt.Errorf("missing SECRET_KEY")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("Invalid token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Extract user ID safely
	sub, ok := claims["sub"].(float64)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID := uint(sub)

	// Fetch user from database
	user, err := repository.GetUserByID(userID)
	if err != nil {
		log.Println("User not found:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Attach user to context
	c.Set("user", user)

	// Proceed with the request
	c.Next()
}
