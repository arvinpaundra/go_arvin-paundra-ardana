package main

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
)

type Book struct {
	ID       string
	Title    string
	Price    int
	Category string
}

var books []Book

func (b *Book) CreateNewBook(title string, price int, category string) {
	b.ID = uuid.NewString()
	b.Title = title
	b.Price = price
	b.Category = category

	books = append(books, *b)

	fmt.Print("Book added!\n\n")
}

func (b *Book) GetAllBooks() {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})

	fmt.Println("All books")
	for _, book := range books {
		fmt.Println("===")
		fmt.Printf("ID: %v\n", book.ID)
		fmt.Printf("Title: %v\n", book.Title)
		fmt.Printf("Price: %v\n", book.Price)
		fmt.Printf("Category: %v\n", book.Category)
		fmt.Println("===")
	}
}

func (b *Book) UpdateBook(id string, title string, price int, category string) {
	for i := range books {
		if id == books[i].ID {
			books[i].Title = title
			books[i].Price = price
			books[i].Category = category

			fmt.Printf("the ID: %v\n", books[i].ID)
			fmt.Print("Book updated!\n\n")
			break
		}
	}
}

func (b *Book) DeleteBook(id string) {
	for i, book := range books {
		if id == book.ID {
			books = append(books[:i], books[i+1:]...)

			fmt.Print("Book deleted!\n\n")
			break
		}
	}
}

func main() {
	book := Book{}

	var menu int

	var id string
	var title string
	var price int
	var category string

	for {
		fmt.Println("===BOOK MANAGEMENT===")
		fmt.Println("1. Get all books")
		fmt.Println("2. Create a book")
		fmt.Println("3. Update a book")
		fmt.Println("4. Delete a book")
		fmt.Println("5. Exit")
		fmt.Print("Choose your menu: ")
		fmt.Scan(&menu)

		if menu == 1 {
			book.GetAllBooks()
		} else if menu == 2 {
			fmt.Print("enter title: ")
			fmt.Scan(&title)
			fmt.Print("enter price: ")
			fmt.Scan(&price)
			fmt.Print("enter category: ")
			fmt.Scan(&category)

			book.CreateNewBook(title, price, category)
		} else if menu == 3 {
			fmt.Print("enter ID: ")
			fmt.Scan(&id)
			fmt.Print("enter title: ")
			fmt.Scan(&title)
			fmt.Print("enter price: ")
			fmt.Scan(&price)
			fmt.Print("enter category: ")
			fmt.Scan(&category)

			book.UpdateBook(id, title, price, category)
		} else if menu == 4 {
			fmt.Print("enter ID: ")
			fmt.Scan(&id)

			book.DeleteBook(id)
		} else if menu == 5 {
			fmt.Println("Bye...")
			break
		} else {
			fmt.Println("Wrong input!")
		}
	}
}
