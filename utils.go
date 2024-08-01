package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func formatAsTable(task task) string {
	// Define column widths (adjust as needed)
	snWidth := 5
	nameWidth := 20
	descWidth := len(task.description)
	statusWidth := len(task.status)
	timeWidth := len(task.timeCreated)

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