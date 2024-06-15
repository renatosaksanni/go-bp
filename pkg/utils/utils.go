package utils

import (
	"fmt"
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v", err))
	}
}

func ValidateData(data interface{}) bool {
	// Implement validation logic
	return true
}
