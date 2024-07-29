package main

import "time"

type task struct {
	Name string
	Description string
	Status string
	TimeCreated time.Time
	TimeCompleted time.Time
}