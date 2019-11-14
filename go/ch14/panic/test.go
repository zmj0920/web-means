package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	os.Exit(-1)
	fmt.Println("End")
}
