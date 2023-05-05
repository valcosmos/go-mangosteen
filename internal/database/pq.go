package database

import (
	// "database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, post, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}

type User struct {
	ID        int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateTables for creating tables
func CreateTables() {
	u1 := User{Email: "Cupid@admin.com"}
	err := DB.Migrator().CreateTable(&u1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully create table")
}

func Migrate() {

}

func Crud() {

}
