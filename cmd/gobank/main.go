package main

import (
	"flag"
	"fmt"
	"gobank/internal"
	"log"
)

func seedAccount(store internal.Storage, fname, lname, pw string) *internal.Account {
	acc, err := internal.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account created:", acc.Number)

	return acc
}

func seedAccounts(store internal.Storage) {
	seedAccount(store, "John", "Doe", "123456")
	seedAccount(store, "Jane", "Doe", "123456")
}

func main() {
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()

	store := internal.NewPostgresStore()

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding the database...")
		seedAccounts(store)
	}

	server := internal.NewAPIServer(":8080", store)
	server.Run()
}
