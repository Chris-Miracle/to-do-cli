package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func createTask(taskList []task) []task {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Give your task a name: ", reader)

	var sn float64 = float64(len(taskList) + 1)

	task := newTask(name, sn)
	fmt.Println(task.name, "- Task Created")

	// Create task details
	task = promptOptions(task)

	// add task to a list
	taskList = append(taskList, task)

	taskOptions(taskList, reader)

	return taskList
}

func promptOptions(task task) task {
	reader := bufio.NewReader(os.Stdin)

	description, _ := getInput("Describe your task: ", reader)
	task.updateDescription(description)

	task = statusOptions(task, reader)

	fmt.Println("Task Details Created!")

	return task
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

func taskOptions(taskList []task, reader *bufio.Reader) {
	opt, _ := getInput("Select Option (A - Save Task to File, B - Create New Task, C - Update An Existing Task): ", reader)

	switch opt {
	case "A":
		saveToFile(taskList, taskList[0].name)
		break
	case "B":
		createTask(taskList)
		break
	case "C":
		fmt.Println("We are working on it")
		break
	default:
		fmt.Println("That was not a valid option ...")
		taskOptions(taskList, reader)
	}
}

func main() {
	var taskList []task
    taskList = createTask(taskList)
}
