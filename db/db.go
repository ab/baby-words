package db

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Connection struct {
	*sqlx.DB
}

func NewConnection(filename string) *Connection {
	log.Printf("Opening sqlite db: %v\n", filename)
	db, err := sqlx.Connect("sqlite", filename)
	if err != nil {
		log.Fatal(err)
	}

	return &Connection{DB: db}
}

func (c *Connection) LogExec(query string, args ...any) (sql.Result, error) {
	log.Printf("SQL: %v, %+v", query, args)
	return c.DB.Exec(query, args...)
}

func (c *Connection) CreateBaby(name string, birth_date string) (*BabyStruct, error) {
	// TODO add client_info row

	// TODO convert birth date?
	query := "INSERT INTO babies (slug, name, birth_date) VALUES (?, ?, ?)"
	b := BabyStruct{
		Slug:      GenerateSlug(),
		Name:      name,
		BirthDate: birth_date,
	}
	fmt.Printf("Creating baby %v\n", b)
	result, err := c.LogExec(query, b.Slug, b.Name, b.BirthDate)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	b.Id = id
	fmt.Printf("Created baby, result: %+v\n", b)
	return &b, nil
}

func (c *Connection) GetBaby(slug string) (*BabyStruct, error) {
	baby := BabyStruct{}
	err := c.Get(&baby, "SELECT * FROM babies WHERE slug = ?", slug)
	if err != nil {
		return nil, err
	}

	return &baby, nil
}

type BabyStruct struct {
	Id        int64
	Slug      string
	Name      string
	BirthDate string `db:"birth_date"` // TODO do custom type http://jmoiron.net/blog/built-in-interfaces
	CreatedAt string `db:"created_at"` // // TODO custom type
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

// GenerateRandomString returns a securely generated random string suitable for
// use as a URL slug.
// Calls log.Fatal on failure.
func GenerateSlug() string {
	const alphabet = "bcdfghjklmnpqrstvwxyz"
	slug, err := GenerateRandomString(alphabet, 16)
	if err != nil {
		log.Fatalf("Failed to generate slug: %v\n", err)
	}
	return slug
}

// GenerateRandomString returns a securely generated random string.
// Returns error on failure.
func GenerateRandomString(alphabet string, length int) (string, error) {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}
		ret[i] = alphabet[num.Int64()]
	}

	return string(ret), nil
}
