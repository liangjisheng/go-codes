package main

import "fmt"

type myError struct{}

func (i myError) Error() string {
	return "myError"
}

func main() {
	err := process()
	fmt.Println(err)

	fmt.Println(err == nil) // false
}

func process() error {
	var err *myError = nil
	return err
}
