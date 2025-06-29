package util

import (
	"encoding/json"
	"fmt"
	"os"
	"wsy0414/task-tracker/model"
)

// EnsureFileExists checks if a file exists at the given path. If it does not, it creates an empty file.
// It returns an error if the check or creation fails.
func EnsureFileExists(path string) error {
	// os.Stat returns an error if the file doesn't exist
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Create the file if it does not exist
		file, createErr := os.Create(path)
		if createErr != nil {
			return fmt.Errorf("failed to create file: %w", createErr)
		}
		// It's important to close the file handle after creating it
		if closeErr := file.Close(); closeErr != nil {
			return fmt.Errorf("failed to close newly created file: %w", closeErr)
		}
		// After creating, initialize it with an empty JSON array
		if writeErr := os.WriteFile(path, []byte("[]"), 0644); writeErr != nil {
			return fmt.Errorf("failed to write initial content to file: %w", writeErr)
		}
	} else if err != nil {
		// Another error occurred (e.g., permission denied)
		return fmt.Errorf("failed to check file existence: %w", err)
	}
	// File exists
	return nil
}

func ReadFile(path string) ([]model.Task, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print("open error:", err.Error())
		return nil, err
	}
	if len(b) == 0 {
		b = []byte("[]")
	}
	result := make([]model.Task, 0)
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Print("unmarshal error:", err.Error())
		return nil, err
	}

	return result, nil
}

func WriteFile(task []model.Task, name string) error {
	b, err := json.Marshal(task)
	if err != nil {
		fmt.Print("json marshal error:", err.Error())
		return err
	}

	err = os.WriteFile(name, b, 0644)
	if err != nil {
		fmt.Print("write error:", err.Error())
		return err
	}
	return nil
}
