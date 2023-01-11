package models

import "github.com/google/uuid"

type Book struct {
	ID      uuid.UUID
	Title   string
	Author  Author
	Edition int
}

type Author struct {
	FirstName, LastName string
}
