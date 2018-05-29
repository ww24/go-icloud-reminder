package icloud

import (
	"fmt"

	"github.com/ww24/go-icloud-reminder/entity"
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

func (r *reminder) Startup() (*entity.StartupResponse, error) {
	entity := new(entity.StartupResponse)
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

func (r *reminder) GetTasks(guid string) (*entity.TasksResponse, error) {
	entity := new(entity.TasksResponse)
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

func (r *reminder) GetCompleted() (*entity.TasksResponse, error) {
	entity := new(entity.TasksResponse)
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
