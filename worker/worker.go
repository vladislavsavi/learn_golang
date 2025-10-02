package worker

import (
	"project/myerror"
)

var NoEnthernet = myerror.MyError{Message: "No Ethernet", Path: ""}

type Worker struct {
	IsWork bool
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) DoWork(val byte) error {
	w.IsWork = true

	if val > 10 {
		return &NoEnthernet
	} else {
		return &myerror.MyError{
			Path:    "/worker.go",
			Message: "You are made mistake",
		}
	}
}
