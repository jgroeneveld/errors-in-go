package main

import (
	"fmt"
)

func checkingTypedErrors() {
	_, err := getSomethingFromDB()

	if err != nil {
		switch err.(type) {
		case TimedError:
			fmt.Println("a timed error", err.Error())
		case ValidationError:
			fmt.Println("a validation error", err.Error())
		default:
			fmt.Println("A general unknown error")
		}
	}

	// or alternatively if you need to access the data inside the TimedError
	if timedErr, ok := err.(TimedError); ok {
		fmt.Printf("there was a timedError at %v", timedErr.Time)
	}
}
