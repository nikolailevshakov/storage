package main

import (
	"fmt"

	"github.com/nikolailevshakov/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello world!", st)
}
