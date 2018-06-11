package icloud

import (
	"fmt"
)

const (
	getStartupPath     = "/rd/startup"
	getTasksPathFormat = "/rd/reminders/%s"
	getCompletedPath   = "/rd/completed"
)

type reminder struct {
	endpoint string
	i        *iCloud
}

func (r *reminder) Startup() (*StartupResponse, error) {
	entity := new(StartupResponse)
	uri := r.endpoint + getStartupPath
	err := r.i.request("GET", uri, nil, entity)
	if err != nil {
		return nil, err
	}
	if entity.Error.Status > 0 {
		return nil, entity.Error
	}
	return entity, nil
}

func (r *reminder) GetTasks(guid string) (*TasksResponse, error) {
	entity := new(TasksResponse)
	uri := r.endpoint + fmt.Sprintf(getTasksPathFormat, guid)
	err := r.i.request("GET", uri, nil, entity)
	if err != nil {
		return nil, err
	}
	if entity.Error.Status > 0 {
		return nil, entity.Error
	}
	return entity, nil
}

func (r *reminder) GetCompleted() (*TasksResponse, error) {
	entity := new(TasksResponse)
	uri := r.endpoint + getCompletedPath
	err := r.i.request("GET", uri, nil, entity)
	if err != nil {
		return nil, err
	}
	if entity.Error.Status > 0 {
		return nil, entity.Error
	}
	return entity, nil
}
