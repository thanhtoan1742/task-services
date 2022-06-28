package storage

import (
	"encoding/json"
	"os"

	"github.com/thanhtoan1742/task-services/api"
)

const filename string = "storage/tasks.json"

func Read() ([]api.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	result := new([]api.Task)
	if json.Unmarshal(data, result) != nil {
		return nil, err
	}
	return *result, nil
}

func Write(tasks []api.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	fm, err := os.Lstat(filename)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, fm.Mode().Perm())
	if err != nil {
		return err
	}
	return nil
}

func Append(task api.Task) error {
	tasks, err := Read()
	if err != nil {
		return err
	}
	tasks = append(tasks, task)
	return Write(tasks)
}
