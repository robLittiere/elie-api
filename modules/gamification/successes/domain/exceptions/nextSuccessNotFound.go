package exceptions

import "fmt"

type NextSuccessNotFound struct {
	Args map[string]interface{}
}

func (e NextSuccessNotFound) Error() string {
	return fmt.Sprintf("Next success not found with args: %v", e.Args)
}
