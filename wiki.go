package main

import (
	"fmt"
	"os"
)

// Let's start by defining the data structures.
type Page struct {
	Title string
	Body  []byte //a byte slice The Body element is a []byte rather than string because that is the type expected by the io libraries we will use
}

// persistent storage? We can address that by creating a save method on Page:
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// The function loadPage constructs the file name from the title parameter,
// reads the file's contents into a new variable body,
// and returns a pointer to a Page literal constructed with the proper title and body values.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// The function viewPage is a simple function that prints the title and body of the Page it is passed.
func viewPage(p *Page) {
	fmt.Println(p.Title)
	fmt.Println(string(p.Body))
}

func Wiki() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	viewPage(p2)
}
