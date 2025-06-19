package controllers

import (
	"iobound/models"
	"iobound/services"
	"iobound/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{service}
}

func (tc *TaskController) GetAll(c *gin.Context) {
	tasks, err := tc.service.GetAll()
	if err != nil {
		log.Printf("Error getting all tasks: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Tasks not found"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ParseStrToInt(idStr)
	if err != nil {
		log.Printf("Invalid task ID format '%s': %v", idStr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	task, err := tc.service.GetByID(id)
	if err != nil {
		log.Printf("Task with ID %d not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ParseStrToInt(idStr)
	if err != nil {
		log.Printf("Invalid task ID format '%s': %v", idStr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = tc.service.Delete(id)
	if err != nil {
		log.Printf("Error deleting task with ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	log.Printf("Task with ID %d deleted successfully", id)
	c.JSON(http.StatusOK, gin.H{"success": "Task deleted successfully"})
}

func (tc *TaskController) Create(c *gin.Context) {
	var task models.Task
	//Doesnt need right now, but when Text for tasks will added that's necessery
	// if err := c.ShouldBindJSON(&task); err != nil {
	// 	log.Printf("Invalid JSON input for task: %v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 	return
	// }

	err := tc.service.Create(&task)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create task"})
		return
	}

	log.Printf("Task created successfully: %+v", task)
	c.JSON(http.StatusOK, gin.H{"success": "Task created successfully"})
}
