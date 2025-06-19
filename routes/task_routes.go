package routes

import (
	"iobound/controllers"

	"github.com/gin-gonic/gin"
)

func initTaskRoutes(c *gin.Engine, tc *controllers.TaskController) {
	service := c.Group("/tasks")
	{
		service.GET("/", tc.GetAll)
		service.GET("/:id", tc.GetByID)
		service.POST("/new/", tc.Create)
		service.DELETE("/:id", tc.Delete)
		//service.PUT("/edit/:id", tc.Update)
	}
}
