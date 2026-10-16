package storage

import (
	"encoding/csv"
	"library-app/domain"
	"os"
	"strconv"
)

func SaveBooksToCSV(filename string, books []*domain.Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Title", "Author", "Year", "IsIssued", "ReaderID"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, book := range books {
		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			strconv.Itoa(book.ReaderID),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return writer.Error()
}

func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	books := []*domain.Book{}

	if len(records) > 0 {
		records = records[1:]
	}

	for _, record := range records {
		if len(record) < 6 {
			return nil, os.ErrInvalid
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, err
		}

		readerID, err := strconv.Atoi(record[5])
		if err != nil {
			return nil, err
		}

		book := &domain.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerID,
		}
		books = append(books, book)
	}

	return books, nil
}