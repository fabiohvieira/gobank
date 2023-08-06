package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account created:", acc.Number)

	return acc
}

func seedAccounts(store Storage) {
	seedAccount(store, "John", "Doe", "123456")
	seedAccount(store, "Jane", "Doe", "123456")
}

func main() {
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()

	store := NewPostgresStore()

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding the database...")
		seedAccounts(store)
	}

	server := NewAPIServer(":8080", store)
	server.Run()
}
