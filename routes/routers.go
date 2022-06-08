package routes

import (
	"bubblePQ/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/",controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo",controller.PostHandler)

		v1Group.GET("/todo",controller.GetHandler)

		v1Group.PUT("/todo/:id", controller.UpdateHandler)

		v1Group.DELETE("/todo/:id", controller.DeleteHandler)

	}
	return r
}