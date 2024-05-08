package main

import (
	"fmt"

	"github.com/scch94/PROBANDO_DATABASES_GO/storage"
)

func main() {
	fmt.Println("projecto con dos bases")
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()

}
