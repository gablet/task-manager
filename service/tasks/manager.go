package tasks

import "gabyRepos/task-manager/model"

type Manager struct{}

func (m *Manager) CreateTask(task *model.Task) error {
	return nil
}

func (m *Manager) RemoveTask(id string) error {
	return nil
}

func (m *Manager) UpdateStatus() error {
	return nil
}
