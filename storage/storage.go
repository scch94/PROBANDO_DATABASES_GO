package storage

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	connStr = "postgres://xxepin:migracion@localhost:5432/testing?sslmode=disable"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("can't open database: %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping to database: %v", err)
		}
		log.Println("se logro conectar a la base postgres")
	})
}

// pool return una isntancia de db
func pool() *sql.DB {
	return db
}
