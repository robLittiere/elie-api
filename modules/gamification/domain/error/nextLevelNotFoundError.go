package error

import "fmt"

type NextLevelNotFoundError struct {
	Args map[string]interface{}
}

func (e NextLevelNotFoundError) Error() string {
	return fmt.Sprintf("Next level not found with args: %v", e.Args)
}
