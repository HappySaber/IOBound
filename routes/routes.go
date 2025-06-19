package routes

import (
	"iobound/controllers"
	"iobound/services"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine) {
	taskService := services.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	initTaskRoutes(c, taskController)
}
