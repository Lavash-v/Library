package cli

import (
	"bufio"
	"fmt"
	"library-app/library"
	"os"
	"strconv"
	"strings"
)

func Run(lib *library.Library, dbPath string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1 - Добавить книгу")
		fmt.Println("2 - Добавить читателя")
		fmt.Println("3 - Выдать книгу")
		fmt.Println("4 - Вернуть книгу")
		fmt.Println("5 - Показать все книги")
		fmt.Println("6 - Показать всех читателей")
		fmt.Println("0 - Сохранить и выйти")
		fmt.Print("Выбор: ")

		if !scanner.Scan() {
			break
		}
		choice, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Введите число.")
			continue
		}

		if choice == 0 {
			fmt.Println("Сохранение данных и выход.")
			if err := lib.SaveLibraryToJSON(dbPath); err != nil {
				fmt.Println("Произошла ошибка сохранения данных:", err)
			}
			return
		}

		switch choice {
		case 1:
			fmt.Print("Название: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Print("Автор: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			fmt.Print("Год: ")
			scanner.Scan()
			year, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			lib.AddBook(title, author, year, 0)

		case 2:
			fmt.Print("Имя: ")
			scanner.Scan()
			firstName := strings.TrimSpace(scanner.Text())

			fmt.Print("Фамилия: ")
			scanner.Scan()
			lastName := strings.TrimSpace(scanner.Text())

			lib.AddReader(firstName, lastName)

		case 3:
			fmt.Print("ID книги: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			fmt.Print("ID читателя: ")
			scanner.Scan()
			readerID, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			if err := lib.IssueBookToReader(bookID, readerID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга выдана.")
			}

		case 4:
			fmt.Print("ID книги: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			if err := lib.ReturnBook(bookID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга возвращена.")
			}

		case 5:
			fmt.Println("--- Книги ---")
			for _, b := range lib.Books {
				fmt.Println(b)
			}

		case 6:
			fmt.Println("--- Читатели ---")
			for _, r := range lib.Readers {
				fmt.Printf("ID: %d, %s %s, активен: %v\n", r.ID, r.FirstName, r.LastName, r.IsActive)
			}

		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}