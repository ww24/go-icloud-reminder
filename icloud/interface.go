package icloud

import "github.com/ww24/go-icloud-reminder/entity"

// Service represents iCloud Services.
type Service interface {
	NewReminder() (Reminder, error)
}

// Reminder represents iCloud Reminder API.
type Reminder interface {
	Startup() (*entity.StartupResponse, error)
	GetTasks(guid string) (*entity.TasksResponse, error)
	GetCompleted() (*entity.TasksResponse, error)
}
