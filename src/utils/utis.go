package utils

import "fmt"

type MyError struct {
	Errcode int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error code: %d", e.Errcode)
}

func ContainsFloat64(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
