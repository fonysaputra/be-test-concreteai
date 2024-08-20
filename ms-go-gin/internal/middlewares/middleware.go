package middlewares

import (
	"fmt"
	"ms-go-gin/config"
	"ms-go-gin/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomMiddleware struct {
	*config.Config
}

func NewCustomMiddleware(s *config.Config) CustomMiddleware {
	return CustomMiddleware{s}
}

// Middleware to verify JWT from Supabase
func (mdl *CustomMiddleware) Jwt(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
		c.Abort()
		return
	}
	resp, err := helper.GetRestyClient(mdl.Config).R().EnableTrace().
		SetHeaders(map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + authHeader}).
		Get("/auth/v1/user")

	if err != nil || resp.StatusCode() != http.StatusOK {
		fmt.Println("Supabase error:", resp.StatusCode(), resp.String())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// Store user data in the context
	c.Set("user", resp.String())
	c.Next()
}
