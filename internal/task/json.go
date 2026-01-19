package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}

func getFilePath() string {
	homeDir, err := os.UserHomeDir()
	checkErr(err)
	appDir := filepath.Join(homeDir, ".tasktracker")

	err = os.MkdirAll(appDir, os.ModePerm)
	checkErr(err)

	return filepath.Join(appDir, "tasks.json")
}

func (t *TasksList) createFileIfNotExists() {
	fp := getFilePath()
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		file, err := os.Create(fp)
		checkErr(err)
		defer file.Close()
	}
}

func (t *TasksList) Load() {
	t.createFileIfNotExists()
	fp := getFilePath()
	data, err := os.ReadFile(fp)
	checkErr(err)
	if len(data) == 0 {
		*t = TasksList{}
		return
	}
	err = json.Unmarshal(data, t)
	checkErr(err)
}

func (t *TasksList) Save() {
	t.createFileIfNotExists()
	fp := getFilePath()
	data, err := json.MarshalIndent(t, "", "  ")
	checkErr(err)
	os.WriteFile(fp, data, 0644)
}
