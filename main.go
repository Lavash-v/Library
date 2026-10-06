package main

import (
	"fmt"
	"library-app/library"
)

func main(){

	fmt.Println("Запуск системы управления библиотекой...")
	
	myLibrary := &library.Library{}

	myLibrary.AddReader("Тамара", "Коляда")
	myLibrary.AddReader("Давид", "Хубаев")

	myLibrary.AddBook("Я чут-чут не книжный червь", "Т. Коляда", 2027, 1)
	myLibrary.AddBook("Мифы древней Греции", "Греки Древние", 1990, 2)

	err := myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ошибка при первом возврате: %v\n", err)
	} else {
		fmt.Println("Успех: Книга успешно возвращена в библиотеку!")
	}

	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка обработана: %v\n", err)
	} else {
		fmt.Println("Ошибка теста: Книга вернулась повторно без генерации ошибки!")
	}
	
	err = myLibrary.ReturnBook(999)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка (нет книги): %v\n", err)
	}

}