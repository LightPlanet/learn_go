package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Book struct {
	ID     string
	Author string
	Label  string
	Avail  bool
}

var Books = []Book{
	{"1", "Кнут", "Искусство программирования", true},
	{"2", "Кнут Грэхем Паташник", "Конкретная математика", true},
	{"3", "Стивен Скиена", "Алгоритмы. Руководство по разработке", true},
	{"4", "Стивен Скиена", "Наука о данных", true},
	{"5", "Владимир Арнольд", "Математичесий тривиум", true},
	{"6", "Владимир Арнольд", "Математические методы классической механики", true},
	{"7", "Валенитина Морозова", "Введение в анализ", true},
	{"8", "Елена Власова", "Ряды", true},
	{"9", "Виноградов", "Основы теории чисел", true},
	{"10", "Сосинский", "Введение в топологию", true},
	{"11", "Успенский", "Простейшие примеры математических доказательств", true},
}

func connectDB() (db *sql.DB) {
	db, err := sql.Open("sqlite", "auth.db")
	if err != nil {
		log.Println(err)
		return nil
	}
	_, err = db.Exec(`
		Create Table IF NOT EXISTS Book(
			ID integer primary key autoincrement,
			Author text not null,
			Label text not null,
			Avail integer not null
		);
			`)
	var count int
	if err = db.QueryRow(`SELECT COUNT(*) FROM Book`).Scan(&count); err != nil {
		log.Println(err)
		return nil
	}
	if count == 0 {
		for _, book := range Books {
			_, err = db.Exec(
				`INSERT INTO Book(ID, Author, Label, Avail) VALUES(?, ?, ?, ?)`,
				book.ID, book.Author, book.Label, book.Avail,
			)
			if err != nil {
				log.Println(err)
				return nil
			}
		}
	}
	return db
}

