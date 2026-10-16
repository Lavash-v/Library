package main

import (
	"fmt"
	"library-app/cli"
	"library-app/library"
	"log"
	"os"
)

func main() {
	const dbFile = "books.json"

	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary, err := library.LoadLibraryFromJSON(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл данных не найден, создана новая библиотека.")
			myLibrary = library.New()
		} else {
			log.Fatalf("Ошибка при загрузке библиотеки: %v", err)
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла.")
	}

	cli.Run(myLibrary, dbFile)
}