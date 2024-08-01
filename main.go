package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func createTask() task {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Give your task a name: ", reader)

	var sn float64 = 0

	if sn != 0 {
		sn++
	} else {
		sn = 1
	}

	task := newTask(name, sn)
	fmt.Println(task.name, "- Task Created")

	task = promptOptions(task)

	return task
}

func promptOptions(task task) task {
	reader := bufio.NewReader(os.Stdin)

	description, _ := getInput("Describe your task: ", reader)
	task.updateDescription(description)

	task = statusOptions(task, reader)

	fmt.Println("Task Details Created!")

	return task
	// opt, _ := getInput("Select Option (A - Save Task to File, B - Create New Task, C - Update An Existing Task, D - Exit): ", reader)

	// fmt.Println(opt)
	// taskOptions()
}

func statusOptions(task task, reader *bufio.Reader) task {
	fmt.Println("What is this tasks' current status?")
	opt, _ := getInput("Select Option (A - Not started, B - In Progress, C - Completed): ", reader)

	switch opt {
	case "A":
		task.updateStatus("Not Started")
		break
	case "B":
		task.updateStatus("In Progress")
		break
	case "C":
		task.updateStatus("Completed")
		task.updateTimeCompleted(time.Now().Format(time.RFC1123))
		break
	default:
		fmt.Println("That was not a valid option ...")
		statusOptions(task, reader)
	}

	return task
}

func main() {
	task := createTask()
	fmt.Print(task.format())
}
