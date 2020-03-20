package main

import "database/sql"

func getSomethingFromDB() (*SomeModel, error) {
	return nil, sql.ErrNoRows
}

type  SomeModel struct {}
