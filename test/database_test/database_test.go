package database_test

import (
	"mangosteen/internal/database"
	"testing"
)

func BenchmarkCurd(b *testing.B) {
	database.Connect()

	database.CreateTables()
	database.Migrate()
	for i := 0; i < b.N; i++ {
		database.Crud()
	}
}
