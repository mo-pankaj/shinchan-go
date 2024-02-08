package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

type StudyMaterial interface {
	Reading()
}

type Book struct {
	BookAuthor string
}

type Magazine struct {
	IssueDate string
}

// the signatures of book is a pointer
func (b *Book) Reading() {
	slog.Info("Reading book", "book", b)
}

func (m *Magazine) Reading() {
	slog.Info("Reading Magazine", "magazine", m)
}

// the signatures of book is not a pointer
func (b Book) String() string {
	return fmt.Sprintf("Reading book %v", b.BookAuthor)
}

func (m Magazine) String() string {
	return fmt.Sprintf("Reading Magazine %v", m.IssueDate)
}

// ReadinHelper is a function that accept an interface and calls methods on it.
// So this way we can have a common function so all types that implements StudyMaterial interface will be here
// Its a unique way of calling methods on a type of interface
func ReadinHelper(s StudyMaterial) {
	s.Reading()
}

// WriteLog our own function, this function accepts an interface and it calls string method
func WriteLog(s fmt.Stringer) {
	slog.Info("Write log" + s.String())

	// lets try to get type
	// and use special type of handling
	value, ok := s.(Book)
	if ok {
		// so its value
		slog.Info("type of value book" + value.BookAuthor)
	}
}

// suppose we want to write to file  as well as to buffer
// it is a beauty of interface, WriteJSON handles specific code and it takes one interface as parameter
// this interface helps us to write it on file/buffer. 
// This code makes it easy for book to be written in a specific form(json here) into the buffer/file
func (b *Book) WriteJSON(io io.Writer) error {
	js, err := json.Marshal(b)
	if err != nil {
		return err
	}

	bytesWritten, err := io.Write(js)
	if err != nil {
		return err
	}
	slog.Info("bytes written", "count", bytesWritten)
	return nil
}

func main() {
	book := Book{
		BookAuthor: "Alex Edwards",
	}

	magazine := Magazine{
		IssueDate: "2024-01-01",
	}

	magazine.Reading()
	book.Reading()

	// using interfaces for our own function
	// now this function can be used to have a common code for different type of interfaces
	WriteLog(magazine)

	// now this function can be used to have a common code for different type of interfaces
	// lets pass book type
	WriteLog(book)

	// We can then call the WriteJSON method using a buffer...
	var buf bytes.Buffer
	err := book.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	// Or using a file.
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = book.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}

}
