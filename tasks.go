package main

import (
	"fmt"
	"strings"
	"time"
)

type task struct {
	sn            int
	name          string
	description   string
	status        string
	timeCreated   string
	timeCompleted string
}

// create new task
func newTask(name string) task {

	task := task{
		sn:            1,
		name:          name,
		description:   "Pending Description",
		status:        "Not started",
		timeCreated:   time.Now().Format(time.RFC1123),
		timeCompleted: "Nil",
	}

	return task
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

func formatAsTable(task task) string {
	// Define column widths (adjust as needed)
	snWidth := 5
	nameWidth := 20
	descWidth := 30
	statusWidth := 10
	timeWidth := 16

	// Build the header row with padding
	header := fmt.Sprintf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		snWidth, "SN",
		nameWidth, "Name",
		descWidth, "Description",
		statusWidth, "Status",
		timeWidth, "Time Created",
		timeWidth, "Time Completed")

	// Build the data row with padding
	dataRow := fmt.Sprintf("| %-*v | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		snWidth, task.sn,
		nameWidth, task.name,
		descWidth, task.description,
		statusWidth, task.status,
		timeWidth, task.timeCreated,
		timeWidth, task.timeCompleted)

	// Combine header and data row with a separator line
	return header + strings.Repeat("-", len(header)) + "\n" + dataRow
}
