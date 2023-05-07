package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mangosteen/config/queries"
	"math/rand"
	"os"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

var DBCtx = context.Background()

// var DB *gorm.DB

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

// Connect for connecting the database
func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {

}

// CreateTables for creating tables
func CreateMigration(filename string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "config/migrations", "-seq", filename)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func Migrate() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatalln(err)
	}
}

func MigrateDown() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func Crud() {
	q := queries.New(DB)
	id := rand.Int()
	u, err := q.CreateAUser(DBCtx, fmt.Sprintf("%d@admin.com", id))
	if err != nil {
		log.Fatalln(err)
	}
	err = q.UpdateUser(DBCtx, queries.UpdateUserParams{
		ID:      u.ID,
		Email:   u.Email,
		Phone:   u.Phone,
		Address: "Shanghai China",
	})
	if err != nil {
		log.Fatalln(err)
	}

	users, err := q.ListUsers(DBCtx, queries.ListUsersParams{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		log.Fatalln(err)
	}
	u, err = q.FindUser(DBCtx, users[0].ID)
	if err != nil {
		log.Println(err)
	}

	err = q.DeleteUser(DBCtx, users[0].ID)
	if err != nil {
		log.Fatalln(err)
	}
}
