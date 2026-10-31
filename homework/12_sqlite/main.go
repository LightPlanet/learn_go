// Создать SQLite Базу данных с таблицей users. Атрибуты таблицы: id, email, password.
// Заполнить таблицу случайными данными.
// Написать модуль на ЯП Go, принимающий на вход пару логин-пароль,
// возвращаюший ответ в зависимости от входных данных:
// -пользователь не найден
// -пароль неверный
// -пароль верный

package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

type User struct {
	ID       int
	Email    string
	Password string
}

var ExampleTableContents = []User{
	{Email: "dmitry@example.com", Password: "111"},
	{Email: "danil@example.com", Password: "222"},
	{Email: "kantemir@example.com", Password: "333"},
	{Email: "yana@example.com", Password: "444"},
	{Email: "yuriy@example.com", Password: "555"},
	{Email: "ramil@example.com", Password: "666"},
}

// Creates "users" table and fills it with 'ExampleTableContents'
// NOTE: Just for homework case
func TryCreateTable(db *sql.DB) error {
	const fnName = "TryCreateTable"

	createTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`
	if _, err := db.Exec(createTable); err != nil {
		return fmt.Errorf("%s: %w", fnName, err)
	}

	// Check if table is not empty
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return fmt.Errorf("%s: %w", fnName, err)
	}
	if count > 0 {
		return nil // Table already exists
	}

	// Useful statement feature - similar to HTML template
	statement, err := db.Prepare(`INSERT INTO users (email, password) VALUES (?, ?);`)
	if err != nil {
		return fmt.Errorf("%s: %w", fnName, err)
	}
	defer statement.Close()

	for _, user := range ExampleTableContents {
		if _, err := statement.Exec(user.Email, user.Password); err != nil {
			return fmt.Errorf("%s: %w", fnName, err)
		}
	}
	return nil
}

// Additional LogIn() errors
var (
	UserNotExists   = errors.New("user not exists")
	InvalidPassword = errors.New("invalid password")
)

func LogIn(db *sql.DB, email, password string) (*User, error) {
	const fnName = "LogIn"
	var usr User

	// The email is unique, no need for "LIMIT 1"
	selectUser := `SELECT id, email, password FROM users WHERE email = ?;`
	err := db.QueryRow(selectUser, email).Scan(&usr.ID, &usr.Email, &usr.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%s: %w", fnName, UserNotExists)
	}
	if err != nil { // Some other error
		return nil, fmt.Errorf("%s: %w", fnName, err)
	}
	if usr.Password != password {
		return nil, fmt.Errorf("%s: %w", fnName, InvalidPassword)
	}
	return &usr, nil
}

func main() {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Fill the table if not exists (first launch)
	if err := TryCreateTable(db); err != nil {
		fmt.Println(err)
		return
	}

	testUsers := []User{
		{Email: "nonexistent@example.com", Password: "666"},
		{Email: ExampleTableContents[3].Email, Password: "333"},
		ExampleTableContents[1],
	}
	for _, usr := range testUsers {
		if _, err := LogIn(db, usr.Email, usr.Password); err != nil {
			fmt.Printf("%s %v\n", err, usr)
			continue
		}
		fmt.Printf("Добро пожаловать, %s!\n", usr.Email)
	}
}
