package router

import (
	"ms-go-gin/config"
	"ms-go-gin/internal/handlers"
	"ms-go-gin/internal/middlewares"

	_ "ms-go-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Backend Test ConcreteAi
// @version 1.0
// @description This is a documentation for the backend test concreteAi.
// @host localhost:8080
// @BasePath /

func New(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Swagger docs route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authHandlers := handlers.NewAuth(cfg)
	accountHandlers := handlers.NewAccount(cfg, db)
	mid := middlewares.NewCustomMiddleware(cfg)
	// Routes for authentication
	v1 := router.Group("/api/v1")
	v1.POST("/signup", authHandlers.SignUp)
	v1.POST("/login", authHandlers.Login)
	v1.GET("/profile", mid.Jwt, authHandlers.Profile)

	v1RouteAccount := v1.Group("/account", mid.Jwt)

	v1RouteAccount.GET("", accountHandlers.List)
	v1RouteAccount.POST("/create", accountHandlers.Create)

	return router
}
