package main

import (
	"encoding/json"
	"os"
)

func LoadTasks() ([]TodoItem, error) {
	raw, err := os.ReadFile("todo.json")
	if err != nil {
		return nil, err
	}
	var tasks []TodoItem
	if err := json.Unmarshal(raw, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(todoItems []TodoItem) error {
	data, error := json.Marshal(&todoItems)

	if error != nil {
		return error
	}

	writeErr := os.WriteFile("todo.json", data, 0644)

	if writeErr != nil {
		return writeErr
	}

	return nil
}
