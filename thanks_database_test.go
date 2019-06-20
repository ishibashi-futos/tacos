package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSave(t *testing.T) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		t.Fatal(err)
	}
	repository := NewThanksRepository(db)
	repository.Save(&Thank{
		"@from", "@to",
	})
}

func TestFindAll(t *testing.T) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		t.Fatal(err)
	}
	repository := NewThanksRepository(db)
	thanks, err := repository.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(thanks) < 1 {
		t.Fatal("len")
	}
	for tx := range thanks {
		t.Log(tx)
	}
}