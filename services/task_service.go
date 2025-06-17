package services

import (
	"errors"
	"iobound/models"
)

type TaskService struct {
	tasks  []models.Task
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make([]models.Task, 0),
		nextID: 1,
	}
}

func (ts *TaskService) Create(task *models.Task) error {
	if task == nil {
		return errors.New("null task")
	}

	task.ID = ts.nextID
	ts.tasks = append(ts.tasks, *task)
	ts.nextID++
	return nil
}

func (ts *TaskService) Update(task models.Task) error {
	for i := range ts.tasks {
		if ts.tasks[i].ID == task.ID {
			ts.tasks[i] = task
			return nil
		}
	}
	return errors.New("task not found")
}

func (ts *TaskService) Delete(id int) error {
	for i := range ts.tasks {
		if ts.tasks[i].ID == id {
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
		}
	}
	return nil
}

func (ts *TaskService) GetByID(id int) (models.Task, error) {
	for i := range ts.tasks {
		if ts.tasks[i].ID == id {
			return ts.tasks[i], nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (ts *TaskService) GetAll() ([]models.Task, error) {
	if ts.tasks == nil {
		return nil, errors.New("tasks not found")
	}
	return ts.tasks, nil
}
