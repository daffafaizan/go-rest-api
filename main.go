package main

import (
	"fmt"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func main() {
	fmt.Println("Hello world")
}
