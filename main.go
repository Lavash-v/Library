package main
import "fmt"

type Address struct {
	City string
	Street string
}

type User struct {
	ID int
	Name string
	Email string
	IsActive bool
	address Address
}

type Author struct {
	FirstName string
	LastName string 
}

type Book struct {
	Title string
	Year int
	BookAuthor Author
	IsIssued bool
}

func (b Book) String() string {
	return fmt.Sprintf(`"%s (%s, %d) %v"`, b.Title, b.BookAuthor, b.Year, b.IsIssued)
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
	u2.address.City = "Beslan"*/

	bookSlice := []Book{}

	book1 := Book{
		Title: "Война и мир",
		Year:  1869,
		BookAuthor: Author{
			FirstName: "Лев",
			LastName:  "Толстой",
		},
		IsIssued: true,
	}

	book2 := Book{
		Title: "Гуру дизайна",
		Year:  2024,
		BookAuthor: Author{
			FirstName: "Апполон",
			LastName:  "Чиколаев",
		},
		IsIssued: true,
	}

	book3 := Book{
		Title: "Как переносить абьюз программистов. Том 1. Плачем вместе",
		Year:  2025,
		BookAuthor: Author{
			FirstName: "Апполон",
			LastName:  "Чиколаев",
		},
		IsIssued: true,
	}

	bookSlice = append(bookSlice, book1, book2, book3)

	book2.IssueBook()
	book3.ReturnBook()

	for _, book := range bookSlice {
		fmt.Printf("%+v\n", book.String())
	}



}

