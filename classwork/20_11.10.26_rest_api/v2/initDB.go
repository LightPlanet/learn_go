package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

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

func DbInit() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "auth.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		Create Table IF NOT EXISTS Books(
			ID integer primary key autoincrement,
			Author text not null,
			Label text not null,
			Avail integer not null
		);
	`)
	if err != nil {
		return nil, err
	}

	var count int
	if err = db.QueryRow(`SELECT COUNT(*) FROM Books`).Scan(&count); err != nil {
		return nil, err
	}
	if count == 0 {
		for _, book := range Books {
			_, err = db.Exec(
				`INSERT INTO Books(ID, Author, Label, Avail) VALUES(?, ?, ?, ?)`,
				book.ID, book.Author, book.Label, book.Avail,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	return db, nil
}
