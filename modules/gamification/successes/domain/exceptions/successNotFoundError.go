package exceptions

import "fmt"

type SuccessNotFoundError struct {
	Args map[string]interface{}
}

func (e SuccessNotFoundError) Error() string {
	return fmt.Sprintf("Success not found with args: %v", e.Args)
}
