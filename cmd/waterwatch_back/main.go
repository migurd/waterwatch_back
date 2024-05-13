package main

import (
	"log"

	"github.com/migurd/waterwatch_back/internal/api"
	"github.com/migurd/waterwatch_back/internal/storage"
)

func main() {
	store, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
