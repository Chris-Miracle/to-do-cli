package main

import (
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

// format task
func (task task) format() string {
	// Check that status is truthy
	if task.status == "Completed" {
		task.timeCompleted = time.Now().Format(time.RFC1123)
	}

	fs := formatAsTable(task)

	return fs
}
