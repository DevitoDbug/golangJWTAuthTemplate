package utils

import "fmt"

type Error struct {
	Context string
	Info    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Context: %v \n Info: %v", e.Context, e.Info)
}

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Field with error: %v , Validation tag that failed: %v, Message: %v", v.Field, v.Tag, v.Message)
}
