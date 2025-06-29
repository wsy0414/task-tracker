package service

import (
	"errors"
	"fmt"
	"time"
	"wsy0414/task-tracker/model"
	"wsy0414/task-tracker/util"
)

type TaskService struct {
	fileName string
}

func NewTaskService(fileName string) *TaskService {
	return &TaskService{
		fileName: fileName,
	}
}

func contentGet(fileName string) ([]model.Task, error) {
	if err := util.EnsureFileExists(fileName); err != nil {
		return nil, fmt.Errorf("failed to ensure file exists: %w", err)
	}

	tasks, err := util.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("read file error")
	}
	return tasks, nil
}

func (service *TaskService) Add(desc string) error {
	tasks, err := contentGet(service.fileName)
	if err != nil {
		return err
	}

	var newID int
	if len(tasks) == 0 {
		newID = 1
	} else {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := model.Task{
		ID:          newID,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
	tasks = append(tasks, newTask)
	err = util.WriteFile(tasks, service.fileName)
	if err != nil {
		return errors.New("write file error")
	}

	return nil
}

func (service *TaskService) Update(id int, desc string) error {
	tasks, err := contentGet(service.fileName)
	if err != nil {
		return err
	}

	for i, v := range tasks {
		if v.ID == id {
			temp := &tasks[i]
			temp.Description = desc
			temp.UpdatedAt = time.Now().String()
			break
		}
		if i == len(tasks)-1 {
			return errors.New("task id not found")
		}
	}

	err = util.WriteFile(tasks, service.fileName)
	if err != nil {
		return err
	}

	return nil
}

func (service *TaskService) Delete(id int) error {
	tasks, err := contentGet(service.fileName)
	if err != nil {
		return err
	}

	for i, v := range tasks {
		if v.ID == int(id) {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	err = util.WriteFile(tasks, service.fileName)
	if err != nil {
		return err
	}

	return nil
}

func (service *TaskService) List(status string) error {
	tasks, err := contentGet(service.fileName)
	if err != nil {
		return err
	}

	switch status {
	case "":
		for _, v := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
		}
	case "done", "todo", "in-progress":
		for _, v := range tasks {
			if v.Status == status {
				fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", v.ID, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
			}
		}
	default:
		return errors.New("invalid status")
	}

	return nil
}

func (service *TaskService) Mark(id int, status string) error {
	tasks, err := contentGet(service.fileName)
	if err != nil {
		return err
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == int(id) {
			temp := &tasks[i]
			temp.Status = status
			temp.UpdatedAt = time.Now().String()
		}
	}

	err = util.WriteFile(tasks, service.fileName)
	if err != nil {
		return err
	}

	return nil
}
