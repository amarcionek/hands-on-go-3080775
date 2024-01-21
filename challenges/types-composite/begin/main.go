// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author
}

// define a library type to track a list of books
type library struct {
	// key is author of the book
	books map[string][]book
}

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l.books[b.author.name] = append(l.books[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(authorName string) []book {
	return l.books[authorName]
}

func main() {
	// create a new library
	l := library{}
	l.books = make(map[string][]book)

	a := author{
		name: "Tolstoy",
	}
	// add 2 books to the library by the same author
	l.addBook(book{
		"Anna Karenina",
		a,
	})
	l.addBook(book{
		"War and Peace",
		a,
	})
	// add 1 book to the library by a different author
	l.addBook(book{
		title: "The Dark Tower",
		author: author{
			name: "Stephen King",
		},
	})

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	tolstoy := l.lookupByAuthor("Tolstoy")

	spew.Dump(tolstoy)

	// print out the first book's title and its author's name
	fmt.Println(tolstoy[:1])

	if len(tolstoy) != 0 {
		b := tolstoy[0]
		fmt.Println(b.title + " - " + b.name)
	}
}
