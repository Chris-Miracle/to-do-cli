package main

import "fmt"

func main() {
	task := newTask("Chris's Task")
	fmt.Print(task.format())
}