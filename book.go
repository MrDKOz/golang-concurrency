package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R Tolkein",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickins",
		YearPublished: 1859,
	},
	{
		ID:            4,
		Title:         "Harry Potter and the Philosophers Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	{
		ID:            5,
		Title:         "Les Miserables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	{
		ID:            6,
		Title:         "I, Robot",
		Author:        "Isaac Asamov",
		YearPublished: 1950,
	},
	{
		ID:            7,
		Title:         "The Gods Themselves",
		Author:        "Isaac Asamov",
		YearPublished: 1973,
	},
	{
		ID:            8,
		Title:         "The Moond is a Hash Mistress",
		Author:        "Robert A. Heinlein",
		YearPublished: 1966,
	},
	{
		ID:            9,
		Title:         "On Basilisk Station",
		Author:        "David Weber",
		YearPublished: 1993,
	},
	{
		ID:            10,
		Title:         "The Android's Dream",
		Author:        "John Scalzi",
		YearPublished: 2006,
	},
}
