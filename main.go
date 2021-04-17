package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Creating a cache, each Book will have an ID
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// Create a Wait Group, and save the pointer
	wg := &sync.WaitGroup{}

	// Create a Mutex, once against using a pointer
	m := &sync.RWMutex{}

	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		// Fetch a random Book ID
		id := rnd.Intn(10) + 1

		// We've got 2 GoRoutines to wait for
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			// Query the cache, if found then print it out
			if b, ok := queryCache(id, m); ok {
				ch <- b // Pass the found Book to the cache channel
			}

			wg.Done()
		}(id, wg, m, cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			// Query the DB, if found then print it out
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}

			wg.Done()
		}(id, wg, m, dbCh)

		// Create on GoRoutine per query to handle response
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("Source: Cache")
				fmt.Println(b)
				<-dbCh // Wait to get the message from the DB Channel so we don't block
			case b := <-dbCh:
				fmt.Println("Source: Database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

// Accepts the ID of the Book, and returns a true/false to whether it was found
func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	// Lock for Writes, however, allowing multiple reads is okay
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

// Simulates running the query against the database - this method is slower than
// querying the cache
func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)

	for _, b := range books {
		if b.ID == id {
			// Add the Book to the cache
			m.Lock()
			cache[id] = b
			m.Unlock()

			// Return the Book
			return b, true
		}
	}

	return Book{}, false
}
