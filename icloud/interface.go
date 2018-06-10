package icloud

// Service represents iCloud Services.
type Service interface {
	NewReminder() (ReminderService, error)
	GetCredentials() *XAppleWebauth
}

// ReminderService represents iCloud Reminder API.
type ReminderService interface {
	Startup() (*StartupResponse, error)
	GetTasks(guid string) (*TasksResponse, error)
	GetCompleted() (*TasksResponse, error)
}
