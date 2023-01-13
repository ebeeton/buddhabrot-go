// Package timer logs elapsed wall clock time.
package timer

import (
	"log"
	"time"
)

// Timer takes the name of a function and is intended to be called from a defer
// statement. The anonymous function returned should be called from the defer
// statment so when the deferred call executes, the elapsed time will be logged.
// A neat trick from https://stackoverflow.com/a/45766707
func Timer(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v.", name, time.Since(start))
	}
}
