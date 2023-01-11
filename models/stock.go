package models

import "github.com/google/uuid"

type BookShop struct {
	Stock Stock
}

func NewBookShop(stock Stock) *BookShop {
	return &BookShop{Stock: stock}
}

type Stock map[uuid.UUID]Book

func (stk Stock) Add(book Book) (uuid.UUID, error) {
	newBookID, err := uuid.NewUUID()
	if err != nil {
		return [16]byte{}, err
	}
	book.ID = newBookID
	stk[book.ID] = book

	return book.ID, nil
}

func (stk Stock) GetBookBy(bookID uuid.UUID) (Book, error) {
	book := stk[bookID]
	return book, nil
}

func (stk Stock) Update(book Book) (Book, error) {
	stk[book.ID] = book
	return stk[book.ID], nil
}
