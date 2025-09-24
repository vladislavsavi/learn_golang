package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%+v\n", createGoEvent())
}

type Event struct {
	Title    string
	Date     string
	Location string
}

func createGoEvent() Event {
	return Event{
		Title:    "День рождения Golang",
		Date:     "10 ноября 2009",
		Location: "10 ноября 2009",
	}
}
