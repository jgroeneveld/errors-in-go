package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func checkingConcreteErrors() {
	_, err := getSomethingFromDB()

	// just check for any error
	if err != nil {
		fmt.Println("there was some error")
	}

	// you can store instances of errors globally to compare them to.
	if err == ErrSomethingSpecial {
		fmt.Println("this was a ErrSomethingSpecial error")
	}

	// exampe from the standard lib
	if err == sql.ErrNoRows {
		fmt.Println("this was a ErrNoRows error")
	}
}

// By convention error names should start with Err*
var ErrSomethingSpecial = errors.New("something special failed")
