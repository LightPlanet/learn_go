// go get gorm.io/gorm
// go get gorm.io/driver/postgres

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Phone struct {
	id     int
	vendor string
	model  string
	price  sql.NullInt64
}

func main() {
	// Open
	const postgresConfig = "user=postgres password=postgres dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", postgresConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// Insert
	query := fmt.Sprintf(`
		INSERT INTO phones ("Vendor", "Model", "Price")
		VALUES ('%s', '%s', %d);`, "Huawey", "Simple 1", 1000)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Select
	rows, err := db.Query(`SELECT * FROM phones;`)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	phones := []Phone{}
	for rows.Next() {
		p := Phone{}
		err := rows.Scan(&p.id, &p.vendor, &p.model, &p.price)
		if err != nil {
			log.Println(err)
			continue
		}
		if !p.price.Valid {
			p.price.Int64 = 0
		}
		phones = append(phones, p)
	}
	for _, v := range phones {
		fmt.Println(v.vendor, v.model, v.price.Int64)
	}
}
