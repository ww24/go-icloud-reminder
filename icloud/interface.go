package icloud

import "github.com/ww24/go-icloud-reminder/entity"

type ICloudService interface {
	NewReminder() (Reminder, error)
}

type Reminder interface {
	GetCompleted() (*entity.TasksResponse, error)
	GetTasks() (*entity.TasksResponse, error)
}
