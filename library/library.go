package library

import (
	"fmt"
	"library-app/domain"
)

type Library struct {
	Books []*domain.Book
	Readers []*domain.Reader

	lastBookID int
	lastReaderID int
}

func (lib* Library) AddReader(firstname, lastname string) *domain.Reader {
	lib.lastReaderID++

	newReader := &domain.Reader{
		ID: lib.lastReaderID,
		FirstName: firstname,
		LastName: lastname,
		IsActive: true,
	}

	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрировался читатель: %v %v\n", newReader.FirstName, newReader.LastName)
	return newReader
}


func (lib *Library) AddBook(title, author string, year int, readerid int) *domain.Book {
	lib.lastBookID++

	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
		ReaderID: readerid,
	}

	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %v\n", newBook)
	return newBook
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {

	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	book.IssueBook()
	reader.AssignBook(book)

	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return fmt.Errorf("не удалось вернуть книгу: %w", err)
	}

	err = book.ReturnBook()
	if err != nil {
		return err 
	}

	return nil
}