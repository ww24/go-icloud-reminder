package icloud

import "github.com/ww24/go-icloud-reminder/entity"

const (
	getTasksPath     = "/rd/reminders/tasks"
	getCompletedPath = "/rd/completed"
)

type reminder struct {
	endpoint string
	i        *iCloud
}

func (r *reminder) Startup() {
	// TODO: implement
}

func (r *reminder) GetCompleted() (*entity.TasksResponse, error) {
	entity := new(entity.TasksResponse)
	uri := r.endpoint + getCompletedPath
	err := r.i.request("GET", uri, nil, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *reminder) GetTasks() (*entity.TasksResponse, error) {
	entity := new(entity.TasksResponse)
	uri := r.endpoint + getTasksPath
	err := r.i.request("GET", uri, nil, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
