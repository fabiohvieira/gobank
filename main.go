package main

import "log"

func main() {
	store := NewPostgresStore()

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", store)
	server.Run()
}
