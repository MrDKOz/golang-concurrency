
# Simple GoLang Concurrecy Project
This is a repo containing a simple project written in Go. This was made using the PluralSight course "Concurrent Programming with Go" by Mike Van Sickle.

This course (and project) covers the following:
 - GoRoutines
 - The Sync Package
	 - WaitGroups
	 - Mutexes
	 - RWMutexes
 - Creating and Using Channels
	 - Buffered and Unbuffered
	 - Types (Bidirectional, Send-only, Receive-only
	 - Closing Channels
	 - Control flow (If, For, and Select)

If you wish to run the project, follow these steps:

 - Ensure you have Go installed on your machine
	 - You can run `go version` in your terminal/command prompt to see if this is the case
	 - *if you don't* you can download it from [golang.org](https://golang.org/)
 - Clone the repo
 - Navigate to the root directory using terminal/command prompt
 - Run `go run .`
 - You should see an output of the various books it found during this loop
 - You can also run `go run --race .` to monitor the execution for any race conditions
