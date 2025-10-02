package myerror

import "fmt"

type MyError struct {
	Message string
	Path    string
}

func (myerr *MyError) Error() string {
	return fmt.Sprintf("MyError - Message: %s, Path: %s", myerr.Message, myerr.Path)
}

func (w *MyError) Is(target error) bool {
	return target == w
}
