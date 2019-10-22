package main

import "fmt"
import "time"

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError {time.Now(),
		"it dodn work",
	}
}

func main () {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
