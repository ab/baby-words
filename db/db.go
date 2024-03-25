package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Connection sql.DB

func NewConnection(filename string) Connection {
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func (c *Connection) GetBaby(slug string) (*BabyStruct, error) {
	row := app.db.QueryRow("SELECT * FROM babies WHERE slug = ?", slug)
	// err := row.Scan(...?)

	return nil, fmt.Errorf("TODO")
}

type BabyStruct struct {
	// TODO
}

func (b *BabyStruct) ListWords() []string {
	return []string{}
}
func (b *BabyStruct) WordInfo(word string) (*WordInfo, error) {
	return nil, fmt.Errorf("TODO")
}
func (b *BabyStruct) AddWord(word string) (*WordInfo, error) {
	return nil, fmt.Errorf("TODO")
}

type WordInfo struct {
	// TODO
}
