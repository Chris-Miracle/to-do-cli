package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type task struct {
	sn            float64
	name          string
	description   string
	status        string
	timeCreated   string
	timeCompleted string
}

// create new task
func newTask(name string, sn float64) task {

	task := task{
		sn:            sn,
		name:          name,
		description:   "Pending Description",
		status:        "Not started",
		timeCreated:   time.Now().Format(time.RFC1123),
		timeCompleted: "Nil",
	}

	return task
}

// update descriptiom
func (task *task) updateDescription(desciption string) {
	task.description = desciption
}

// update status
func (task *task) updateStatus(status string) {
	task.status = status
}

// update timeCompleted
func (task *task) updateTimeCompleted(time string) {
	task.timeCompleted = time
}

// save task(s) to file
func saveToFile(taskList []task, fileName string) {
	var fileContent string

	// Print the header
	fileContent += "| SN  | Name           | Description               | Status    | Time Created      | Time Completed    |\n"
	fileContent += strings.Repeat("-", 80) + "\n"

	for _, task := range taskList {
		fileContent += task.format()
	}

	err := os.WriteFile("todos/"+fileName+".txt", []byte(fileContent), 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Tasks saved to", "todos/"+fileName)
}

// format task
func (task task) format() string {
	fs := formatAsTable(task)

	return fs
}
