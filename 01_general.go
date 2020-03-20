package main

import (
	"errors"
	"fmt"
	"time"
)

func generalErrors() {
	// go's standard way of making simple errors
	err := errors.New("A generic error")
	err = fmt.Errorf("A formatted error: %s", "Some other value")

	// errors are just implementations of the error interface
	type error interface {
		Error() string
	}

	// custom errors
	err = TimedError{
		Time:    time.Now(),
		Message: "A Message",
	}

	_ = err // ignore me
}

// Custom Error types by convention end with *Error
type TimedError struct {
	Time    time.Time
	Message string
}

func (err TimedError) Error() string {
	return fmt.Sprintf("TimedError(%v): %s", err.Time, err.Message)
}

