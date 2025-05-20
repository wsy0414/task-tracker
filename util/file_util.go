package util

import (
	"encoding/json"
	"fmt"
	"os"
	"wsy0414/task-tracker/model"
)

func CheckFileExist(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		if err == os.ErrNotExist {
			file, err = os.Create(path)
			if err != nil {
				fmt.Print("create error:", err.Error())
				return false
			}
		}
		fmt.Print("open error:", err.Error())
		return false
	}
	defer file.Close()

	return true
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
