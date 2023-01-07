package models

import "github.com/google/uuid"

type BookShop struct {
	Stock Inventory
}

func NewBookShop(stock Inventory) *BookShop {
	return &BookShop{Stock: stock}
}

type Inventory map[uuid.UUID]Book

func (stk Inventory) Add(book Book) (uuid.UUID, error) {
	newBookID, err := uuid.NewUUID()
	if err != nil {
		return [16]byte{}, err
	}
	book.ID = newBookID
	stk[book.ID] = book

	return book.ID, nil
}

func (stk Inventory) GetBookBy(bookID uuid.UUID) (Book, error) {
	book := stk[bookID]
	return book, nil
}
