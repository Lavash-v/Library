package library

import (
	"fmt"
	"library-app/domain"
	"library-app/storage"
)

type Library struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`

	LastBookID   int `json:"last_book_id"`
	LastReaderID int `json:"last_reader_id"`
}

func New() *Library {
	return &Library{
		Books:   []*domain.Book{},
		Readers: []*domain.Reader{},
	}
}

func (lib *Library) AddReader(firstname, lastname string) *domain.Reader {
	lib.LastReaderID++

	newReader := &domain.Reader{
		ID:        lib.LastReaderID,
		FirstName: firstname,
		LastName:  lastname,
		IsActive:  true,
	}
	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрировался читатель: %v %v\n", newReader.FirstName, newReader.LastName)
	return newReader
}

func (lib *Library) AddBook(title, author string, year int, readerid int) *domain.Book {
	lib.LastBookID++

	newBook := &domain.Book{
		ID:       lib.LastBookID,
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
		return fmt.Errorf("выдача отклонена: %w", err)
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return fmt.Errorf("выдача отклонена: %w", err)
	}

	if book.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана другому читателю", book.Title)
	}

	if err := reader.AssignBook(book); err != nil {
		return fmt.Errorf("не удалось привязать книгу к читателю: %w", err)
	}

	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return fmt.Errorf("не удалось вернуть книгу: %w", err)
	}
	return book.ReturnBook()
}

func (lib *Library) SaveLibraryToJSON(filename string) error {
	return storage.WriteJSON(filename, lib)
}

func LoadLibraryFromJSON(filename string) (*Library, error) {
	lib := &Library{}
	if err := storage.ReadJSON(filename, lib); err != nil {
		return nil, err
	}

	if lib.Books == nil {
		lib.Books = []*domain.Book{}
	}
	if lib.Readers == nil {
		lib.Readers = []*domain.Reader{}
	}

	return lib, nil
}
