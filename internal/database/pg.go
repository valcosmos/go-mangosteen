package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	post     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

// Connect for connecting the database
func Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, post, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connect to db")
}

// CreateTables for creating tables
func CreateTables() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		log.Fatal()
	}
	log.Println("Successfully create users table")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	handleError(err)
	log.Println("Successfully add phone column to users table")

	_, err = DB.Exec("ALTER TABLE users ADD COLUMN address VARCHAR(200)")
	handleError(err)
	log.Println("Successfully add address column to users table")

	// 新增 Items表，字段为 id, amount, happened_at, created_at, updated_at
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		amount INTEGER NOT NULL,
		happened_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,	
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,	
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,	
	)`)
	// postgreSQL修改字段类型
	// _, err = DB.Exec("ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP")
	handleError(err)
	log.Println("Successfully create items table")
}
