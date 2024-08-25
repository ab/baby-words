package db

import (
	"crypto/rand"
	"database/sql"
	"errors"
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

func (c *Connection) ListWords(babyId int64) ([]WordInfo, error) {
	words := []WordInfo{}
	err := c.Select(&words, "SELECT * FROM words WHERE baby_id = ?", babyId)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (c *Connection) getMaxWordNumber(babyId int64) (int64, error) {
	maxWord := WordInfo{}
	err := c.Get(&maxWord, "SELECT * FROM words WHERE baby_id = ? ORDER BY number DESC LIMIT 1", babyId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// no existing words
			return 0, nil
		}

		// unexpected error
		return 0, err
	}

	log.Printf("Previous max word has number=%v", maxWord.Number)
	return maxWord.Number, nil
}

func (c *Connection) AddWord(babyId int64, word string, learnedDate string) (int64, error) {

	maxWordNumber, err := c.getMaxWordNumber(babyId)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO words (baby_id, word, number, learned_date) VALUES (?, ?, ?, ?)"
	result, err := c.LogExec(query, babyId, word, maxWordNumber+1, learnedDate)
	if err != nil {
		return 0, err
	}

	createdId, _ := result.LastInsertId()
	return createdId, nil
}

type BabyStruct struct {
	Id           int64
	Slug         string
	Name         string
	BirthDate    string `db:"birth_date"` // TODO do custom type http://jmoiron.net/blog/built-in-interfaces
	CreatedAt    string `db:"created_at"` // // TODO custom type
	Timezone     sql.NullString
	ClientInfoId sql.NullInt64 `db:"client_info_id"`
}

type ClientInfo struct {
	Id        int64
	UserAgent string
	IpAddress string
	CreatedAt string `db:"created_at"` // TODO custom type
}

type WordInfo struct {
	Id           int64
	Baby_Id      int64
	Word         string
	Number       int64
	Learned_Date string        // TODO type
	CreatedAt    string        `db:"created_at"` // TODO custom type
	ClientInfoId sql.NullInt64 `db:"client_info_id"`
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
