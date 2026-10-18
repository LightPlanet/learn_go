package main

// github.com/mattn/go-sqlite3
// modernc.org/sqlite

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Phone struct {
	ID     int
	Vendor string
	Model  string
	Price  sql.NullInt64
}

func main() {
	db, err := sql.Open("sqlite", "example_db")
	if err != nil {
		log.Println("sql.Open", err)
		return
	}
	defer db.Close()

	// SELECT
	rows, err := db.Query(`SELECT * FROM 'Mobile Phones'`)
	if err != nil {
		log.Println("select", err)
		return
	}
	defer rows.Close()

	phones := []Phone{}
	for rows.Next() {
		p := Phone{}
		err := rows.Scan(&p.ID, &p.Vendor, &p.Model, &p.Price)
		if err != nil {
			log.Println("rows.Scan", err)
			continue
		}
		if err := rows.Err(); err != nil {
			log.Println("rows.Err", err)
			return
		}
		phones = append(phones, p)
	}

	log.Println(phones)

	// INSERT
	_, err = db.Exec(`
		INSERT INTO phones (Vendor, Model, Price) 
		VALUES (?, ?, ?)`, "Huawey", "Simple 1", 1000)
	if err != nil {
		log.Println(err)
		return
	}
}
