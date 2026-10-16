package storage

import (
	"encoding/csv"
	"encoding/json"
	"library-app/domain"
	"os"
	"strconv"
)

// ---------- JSON: низкоуровневые операции ----------

// WriteJSON сериализует произвольное значение v в файл filename
// в «красивом» формате (с отступами).
func WriteJSON(filename string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// ReadJSON читает JSON из файла filename и десериализует в v.
// Если файла нет — возвращает ошибку, для которой os.IsNotExist(err) == true.
func ReadJSON(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// ---------- CSV: оставлено для обратной совместимости ----------

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

		books = append(books, &domain.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerID,
		})
	}

	return books, nil
}