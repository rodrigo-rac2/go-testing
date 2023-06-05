package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of division is", result)
}

// divide x by y
func divide(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}
