package models

import (
	"fmt"
	"github.com/google/uuid"
)

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
	book, ok := stk[bookID]
	if !ok {
		return Book{}, fmt.Errorf("this book does not exist in the inventory. bookID: %v", bookID)
	}
	return book, nil
}

func (stk Stock) Update(book Book) (Book, error) {
	stk[book.ID] = book
	return stk[book.ID], nil
}

func (stk Stock) Delete(book Book) error {
	book, ok := stk[book.ID]
	if ok {
		delete(stk, book.ID)
		return nil
	}
	return fmt.Errorf("there is no unit of this book in the stock, %v", book)
}
