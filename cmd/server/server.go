package main

import (
	controllers "jwt-gin/pkg/controller"
	"jwt-gin/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/hello", controllers.Hello)

	r.Run(":8080")
}
