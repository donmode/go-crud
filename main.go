package main

import (
	"crud/initializers"

	"github.com/gin-gonic/gin"

	"crud/controllers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
