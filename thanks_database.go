package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbfile = "./thanks.db"

type Thanks interface {
	Save(thank Thank) error
}

type Thank struct {
	Source      string
	Destination string
}

type ThanksRepository struct {
	db *sql.DB
}

func NewThanksRepository(db *sql.DB) *ThanksRepository {
	initdb()
	return &ThanksRepository{db}
}

func initdb() (error) {
	_, err := os.Stat(dbfile)
	// exists database
	if err == nil {
		return nil
	}
	// if not exist database
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return err
	}
	defer db.Close()
	c := `CREATE TABLE THANKS(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		SOURCE,
		DESTINATION,
		CREATED_AT TIMESTAMP DEFAULT (DATETIME('now','localtime'))
	)`
	if _, err := db.Exec(c); err != nil {
		return err
	}
	return nil
}

func (thanks *ThanksRepository) Save(model *Thank) (error) {
	sql := fmt.Sprintf("INSERT INTO THANKS(SOURCE, DESTINATION) VALUES('%s', '%s')", model.Source, model.Destination)
	if _, err := thanks.db.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (thanks *ThanksRepository) FindAll() ([]*Thank, error) {
	result := make([]*Thank, 0)
	sql := "SELECT SOURCE, DESTINATION FROM THANKS"
	rows, err := thanks.db.Query(sql)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var source string
		var destination string
		if err := rows.Scan(&source, &destination); err != nil {
			return result, err
		}
		result = append(result, &Thank{source, destination})
	}
	return result, nil
}