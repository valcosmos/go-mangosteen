package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *sql.DB

var DB *gorm.DB

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
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID         int
	UserID     int
	Amount     int
	HappenedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Tag struct {
	ID   int
	Name string
}

var models = []any{&User{}, &Item{}, &Tag{}}

// CreateTables for creating tables
func CreateTables() {
	for _, model := range models {
		err := DB.Migrator().CreateTable(model)
		if err != nil {
			log.Println(err)
		}
	}

	// err := DB.Migrator().CreateTable(&User{}, &Item{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	log.Println("Successfully create table")
}

func Migrate() {
	DB.AutoMigrate(models...)
}

func Crud() {
	// // 创建一个user
	// user := User{Email: "admin1@admin.com"}
	// // tx 是事务
	// tx := DB.Create(&user)
	// if tx.Error != nil {
	// 	log.Println(tx.Error)
	// }

	// log.Println(tx.RowsAffected)
	// log.Println(user)
	// u2 := &User{}
	// tx := DB.Find(&u2, 1)
	// u2.Phone = "12345"
	// tx = DB.Save(&u2)
	// if tx.Error != nil {
	// 	log.Println(tx.Error)
	// } else {
	// 	log.Println(tx.RowsAffected)
	// 	log.Println(u2)
	// }
	users := []User{}
	// DB.Find(&users, []int{1, 4})
	DB.Offset(0).Limit(10).Order("created_at desc, id desc").Find(&users)
	log.Println(users)

	u := User{ID: 1}
	tx := DB.Delete(&u)
	if tx.Error != nil {
		log.Println(tx.Error)
	} else {
		log.Println(tx.RowsAffected)
	}
}
