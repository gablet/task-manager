package service

import "gabyRepos/task-manager/model"

type Manager interface {
	CreateTask(task model.Task) error
	RemoveTask(id string) error
	UpdateStatus() error
}
