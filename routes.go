package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ketabdoozak/backend/handlers"
	"github.com/ketabdoozak/backend/middlewares"
)

// RegisterRoutes register routes on router
func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", handlers.Register())
	r.POST("/login", handlers.Login())
	r.DELETE("/logout", middlewares.Authenticate(), handlers.Logout())

	r.GET("/categories", handlers.ListCategories())

	r.GET("/locations", handlers.ListLocations())
}
