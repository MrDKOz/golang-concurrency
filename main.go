package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creating a cache, each Book will have an ID
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		// Fetch a random Book ID
		id := rnd.Intn(10) + 1

		// Query the cache, if found then print it out
		if b, ok := queryCache(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}

		// Query the DB, if found then print it out
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}

		// Could not find a book with the given ID
		fmt.Printf("Book not found with ID: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
}

// Accepts the ID of the Book, and returns a true/false to whether it was found
func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

// Simulates running the query against the database - this method is slower than
// querying the cache
func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)

	for _, b := range books {
		if b.ID == id {
			// Add the Book to the cache
			cache[id] = b

			// Return the Book
			return b, true
		}
	}

	return Book{}, false
}
