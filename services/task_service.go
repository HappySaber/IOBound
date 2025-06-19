package services

import (
	"errors"
	"iobound/models"
	"iobound/utils"
	"log"
	"sync"
	"time"
)

type TaskService struct {
	tasks  map[int]*models.Task
	nextID int
	mu     sync.Mutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]*models.Task),
		nextID: 1,
	}
}

func (ts *TaskService) Create(task *models.Task) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if task == nil {
		return errors.New("null task")
	}

	task.ID = ts.nextID
	task.Status = "pending"
	task.CreatedAt = time.Now()

	ts.tasks[task.ID] = task
	ts.nextID++

	go ts.processTask(task.ID)

	return nil
}

func (ts *TaskService) Delete(id int) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if _, err := ts.tasks[id]; !err {
		return errors.New("task not found")
	}
	delete(ts.tasks, id)
	return nil
}

func (ts *TaskService) GetByID(id int) (*models.TaskResponse, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task, err := ts.tasks[id]
	if !err {
		return nil, errors.New("task not found")
	}
	taskResponse := &models.TaskResponse{
		ID:                task.ID,
		Status:            task.Status,
		CreatedAt:         task.CreatedAt,
		Text:              task.Text,
		DurationFormatted: task.DurationFormatted,
	}
	return taskResponse, nil
}

func (ts *TaskService) GetAll() ([]*models.TaskResponse, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	taskResponses := make([]*models.TaskResponse, 0, len(ts.tasks))
	for _, t := range ts.tasks {
		taskResponse := &models.TaskResponse{
			ID:                t.ID,
			Status:            t.Status,
			CreatedAt:         t.CreatedAt,
			Text:              t.Text,
			DurationFormatted: t.DurationFormatted,
		}
		taskResponses = append(taskResponses, taskResponse)
	}
	return taskResponses, nil
}

func (ts *TaskService) processTask(id int) {
	ts.mu.Lock()
	task, err := ts.tasks[id]
	if !err {
		ts.mu.Unlock()
		return
	}
	task.Status = "in_progress"
	task.StartedAt = time.Now()
	ts.mu.Unlock()

	log.Printf("Task %d started", id)
	//time.Sleep(3 * time.Minute)
	//time.Sleep(30 * time.Second) //test

	durationOfTask := utils.RandRange(180, 300)

	for range durationOfTask {
		time.Sleep(1 * time.Second)

		ts.mu.Lock()
		task.Duration = time.Since(task.StartedAt)

		task.DurationFormatted = utils.FormatDuration(task.Duration)
		ts.mu.Unlock()
	}

	ts.mu.Lock()
	defer ts.mu.Unlock()

	task.CompletedAt = time.Now()
	task.Duration = time.Since(task.StartedAt)
	task.Status = "done"
	log.Printf("Task %d completed", id)
}
