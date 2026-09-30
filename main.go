package main
import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

type SMSNotifier struct {
	ProneNumber string
}

func (email EmailNotifier) Notify(message string) {
	fmt.Printf("\nОтправляю email на %v: '%v'\n", email.EmailAddress, message)
}

func (sms SMSNotifier) Notify(message string) {
	fmt.Printf("\nОтправляю SMS на номер %v: '%v'\n", sms.ProneNumber, message)
}


type Reader struct {
	ID int
	FirstName string
	LastName string
	Email string
	IsActive bool
}

/*type Author struct {
	FirstName string
	LastName string 
}*/

type Book struct {
	ID int
	Title string
	Year int
	Author string
	IsIssued bool
}

type Library struct {
	Books []*Book
	Readers []*Reader

	lastBookID int
	lastReaderID int
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

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("\nКнига %s и так в библиотеке!", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("\nКнига %s возвращена в библиотеку.", b.Title)

}

func (lib* Library) AddReader(firstname, lastname string) *Reader {
	lib.lastReaderID++

	newReader := &Reader{
		ID: lib.lastReaderID,
		FirstName: firstname,
		LastName: lastname,
		IsActive: true,
	}

	lib.Readers = append(lib.Readers, newReader)

	fmt.Printf("Зарегистрировался читатель: %v %v\n", newReader.FirstName, newReader.LastName)
	return newReader
}


func (lib *Library) AddBook(title, author string, year int) *Book {
	lib.lastBookID++

	// Создаем новую книгу
	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %v\n", newBook)
	return newBook
}


func main(){
	/*u := User{
		ID: 1,
		Name: "Aleksandr",
		Email: "lohhhh@gmail.com",
		IsActive: true,
		address: Address{City: "bebe", Street: "oooo"},
	}
	fmt.Printf("%+v\n", u)
	u2 := new(User)
	fmt.Printf("Нулевой пользователь: %+v\n", u2)

	u2.ID = 2
	u2.IsActive = false
	u2.Name = "Ilon Mask"
	u2.address.City = "Beslan"

	bookSlice := []Book{}

	book1 := Book{
		Title: "Война и мир",
		Year:  1869,
		Author: "Лев Толстой",
		IsIssued: true,
	}

	book2 := Book{
		Title: "Гуру дизайна",
		Year:  2024,
		Author: "Апполон Чиколаев",
		IsIssued: true,
	}

	book3 := Book{
		Title: "Как переносить абьюз программистов. Том 1. Плачем вместе",
		Year:  2025,
		Author: "Апполон Чиколаев",
		IsIssued: true,
	}

	bookSlice = append(bookSlice, book1, book2, book3)

	book2.IssueBook()
	book3.ReturnBook()

	for _, book := range bookSlice {
		fmt.Printf("%+v\n", book.String())
	}

	notifier := []Notifier{}

	e1 := EmailNotifier{EmailAddress: "t.s.kolyada@gmail.com"}
	sms1 := SMSNotifier{ProneNumber: "+79288884813"}

	notifier = append(notifier, e1, sms1)

	for _, notify := range notifier {
		notify.Notify("Ваша книга просрочена!")
	}*/

	fmt.Println("Запуск системы управления библиотекой...")
	
	// 1. Создаем экземпляр библиотеки
	myLibrary := &Library{}

	// 2. Добавляем читателей
	myLibrary.AddReader("Тамара", "Коляда")
	myLibrary.AddReader("Давид", "Хубаев")

	// 3. Добавляем книги
	myLibrary.AddBook("Я чут-чут не книжный червь", "Т. Коляда", 2027)
	myLibrary.AddBook("Мифы древней Греции", "Греки Древние", 1990)

}

