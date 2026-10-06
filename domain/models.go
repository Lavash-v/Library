package domain

import (
	"fmt"
)

type Reader struct {
	ID int
	FirstName string
	LastName string
	Email string
	IsActive bool
}

type Book struct {
	ID int
	Title string
	Year int
	Author string
	IsIssued bool
	ReaderID int
}

func (b Book) String() string {
	return fmt.Sprintf(`"%s (%s, %d) %v"`, b.Title, b.Author, b.Year, b.IsIssued)
}

func (b *Book) IssueBook() {
	if b.IsIssued {
		fmt.Printf("\nКнига %s уже кому-то выдана", b.Title)
		return
	}
	b.IsIssued = true
	fmt.Printf("\nКнига %s была выдана", b.Title)
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}

	b.IsIssued = false
	b.ReaderID = 0
	return nil
}

func (r *Reader) AssignBook(book *Book) error { // Теперь без Printf
	if book == nil {
		return fmt.Errorf("указатель на книгу равен nil")
	}

	book.IsIssued = true
	book.ReaderID = r.ID
	return nil
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r *Reader) Activate() {
	r.IsActive = true
}