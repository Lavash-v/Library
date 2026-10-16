package main

import (
	"fmt"
	"library-app/library"
	"os"
)

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary := &library.Library{}

	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "save":
			if len(args) < 2 {
				fmt.Println("Использование: save <имя_файла.csv>")
				return
			}
			if err := myLibrary.SaveBooksToCSV(args[1]); err != nil {
				fmt.Printf("Ошибка сохранения: %v\n", err)
				return
			}
			fmt.Printf("Книги успешно сохранены в %s\n", args[1])
			return
		case "load":
			if len(args) < 2 {
				fmt.Println("Использование: load <имя_файла.csv>")
				return
			}
			if err := myLibrary.LoadBooksFromCSV(args[1]); err != nil {
				fmt.Printf("Ошибка загрузки: %v\n", err)
				return
			}
			fmt.Printf("Книги успешно загружены из %s\n", args[1])
			for _, b := range myLibrary.Books {
				fmt.Println(b)
			}
			return
		default:
			fmt.Printf("Неизвестная команда: %s\n", args[0])
			fmt.Println("Доступные команды: save <файл.csv>, load <файл.csv>")
			return
		}
	}

	r1 := myLibrary.AddReader("Тамара", "Коляда")
	r2 := myLibrary.AddReader("Давид", "Хубаев")

	b1 := myLibrary.AddBook("Я чут-чут не книжный червь", "Т. Коляда", 2027, 1)
	b2 := myLibrary.AddBook("Мифы древней Греции", "Греки Древние", 1990, 2)

	fmt.Println("\n---")

	// успешная выдача книги
	err := myLibrary.IssueBookToReader(b2.ID, r2.ID)
	if err != nil {
		fmt.Printf("Ошибка: Не удалось выдать книгу Давиду: %v\n", err)
	} else {
		fmt.Printf("Успех: Книга '%s' успешно выдана читателю %s %s\n", b2.Title, r2.FirstName, r2.LastName)
	}

	// попытка выдать книгу, которая УЖЕ взята
	err = myLibrary.IssueBookToReader(b2.ID, r1.ID)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	} else {
		fmt.Println("Ошибка теста: Книга выдана повторно, хотя она уже на руках")
	}

	// попытка вернуть книгу, которую никто НЕ БРАЛ
	err = myLibrary.ReturnBook(b1.ID)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	} else {
		fmt.Println("Ошибка теста: Библиотека приняла книгу, которую никто не брал")
	}

	// успешный возврат книги b2, которую взяли
	err = myLibrary.ReturnBook(b2.ID)
	if err != nil {
		fmt.Printf("Ошибка: Не удалось вернуть книгу в библиотеку: %v\n", err)
	} else {
		fmt.Printf("Успех: Книга '%s' успешно возвращена на полку\n", b2.Title)
	}

	// попытка вернуть несуществующую книгу
	err = myLibrary.ReturnBook(999)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка (нет книги): %v\n", err)
	} else {
		fmt.Println("Ошибка теста: Найдена и возвращена несуществующая книга")
	}

	// теперь успешно выдаем книгу b1 (она всё это время была свободна)
	err = myLibrary.IssueBookToReader(b1.ID, r1.ID)
	if err != nil {
		fmt.Printf("Ошибка: Не удалось выдать книгу: %v\n", err)
	} else {
		fmt.Printf("Успех: Книга '%s' успешно выдана читателю %s %s\n", b1.Title, r1.FirstName, r1.LastName)
	}
}