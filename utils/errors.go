package utils

import "fmt"

type Error struct {
	Context string
	Info    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Context: %v \n Info: %v", e.Context, e.Info)
}
