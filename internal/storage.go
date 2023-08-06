package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccountByID(id int) (*Account, error)
	GetAccountByNumber(number int64) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() *PostgresStore {
	connStr := "user=gobank password=gobank dbname=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &PostgresStore{db: db}
}

func (s *PostgresStore) Init() error {
	if err := s.createAccountTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) createAccountTable() error {
	query := "CREATE TABLE IF NOT EXISTS accounts (" +
		"id SERIAL PRIMARY KEY, " +
		"first_name TEXT, " +
		"last_name TEXT, " +
		"number BIGINT, " +
		"encrypted_password TEXT, " +
		"balance BIGINT," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP" +
		")"

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := "INSERT INTO accounts (first_name, last_name, number, encrypted_password, balance) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.db.Query(query, account.FirstName, account.LastName, account.Number, account.EncryptedPassword, account.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := "DELETE FROM accounts WHERE id = $1"
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := "SELECT * FROM accounts WHERE id = $1"
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account not found")
}

func (s *PostgresStore) GetAccountByNumber(number int64) (*Account, error) {
	query := "SELECT * FROM accounts WHERE number = $1"
	rows, err := s.db.Query(query, number)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account not found")
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, 0)
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := &Account{}
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
