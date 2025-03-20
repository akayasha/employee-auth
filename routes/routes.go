package routes

import (
	"employee-auth/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// Group Auth
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.RegisterUser)
		authRoutes.POST("/login", controllers.LoginUser)
	}

	return r

}
