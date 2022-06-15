package main

import (
	"database/sql"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	_ "github.com/marcboeker/go-duckdb"
	_ "github.com/mattn/go-sqlite3"
)

func BenchmarkInserts(b *testing.B) {
	rand.Seed(time.Now().UnixMicro())

	b.Run("duckdb", func(b *testing.B) {
		db := mustSetupDuckDB(b)
		defer db.Close()

		runInserts(b, db)
	})

	b.Run("sqlite3", func(b *testing.B) {
		db := mustSetupSQLite(b)
		defer db.Close()

		runInserts(b, db)
	})
}

func runInserts(b *testing.B, db *sql.DB) {
	data := generateTestData(b.N)

	b.ResetTimer()

	for _, d := range data {
		mustExec(b, db, "INSERT INTO test_models (name, email, age) VALUES (?, ?, ?)", d.Name, d.Email, d.Age)
	}
}

func BenchmarkOneInsert(b *testing.B) {
	rand.Seed(time.Now().UnixMicro())

	b.Run("duckdb", func(b *testing.B) {
		db := mustSetupDuckDB(b)
		defer db.Close()

		runOneInsert(b, db)
	})

	b.Run("sqlite3", func(b *testing.B) {
		db := mustSetupSQLite(b)
		defer db.Close()

		runOneInsert(b, db)
	})
}

const maxChunk = 100

func runOneInsert(b *testing.B, db *sql.DB) {
	data := generateTestData(b.N)

	b.ResetTimer()

	for len(data) > 0 {
		var chunk []TestModel
		if len(data) > maxChunk {
			chunk = data[:maxChunk]
			data = data[maxChunk:]
		} else {
			chunk = data
			data = nil
		}
		query := "INSERT INTO test_models (name, email, age) VALUES "
		rows := make([]string, 0, b.N)
		args := make([]interface{}, 0, b.N*3)

		for _, d := range chunk {
			rows = append(rows, "(?, ?, ?)")
			args = append(args, d.Name, d.Email, d.Age)
		}

		query += strings.Join(rows, ", ")

		mustExec(b, db, query, args...)
	}
}

func mustSetupDuckDB(b *testing.B) *sql.DB {
	db, err := sql.Open("duckdb", "data.duckdb")
	if err != nil {
		b.Fatal(err)
	}

	mustExec(b, db, "DROP TABLE IF EXISTS test_models")
	mustExec(b, db, "CREATE TABLE test_models (name VARCHAR, email VARCHAR, age INT)")
	mustExec(b, db, "VACUUM")

	return db
}

func mustSetupSQLite(b *testing.B) *sql.DB {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		b.Fatal(err)
	}

	mustExec(b, db, "DROP TABLE IF EXISTS test_models")
	mustExec(b, db, "CREATE TABLE test_models (name VARCHAR, email VARCHAR, age INT)")
	mustExec(b, db, "VACUUM")

	return db
}

func mustExec(b *testing.B, db *sql.DB, query string, args ...any) {
	_, err := db.Exec(query, args...)
	if err != nil {
		b.Fatal(err)
	}
}

type TestModel struct {
	Name  string `faker:"name"`
	Email string `faker:"email"`
	Age   int    `faker:"-"`
}

func generateTestData(n int) []TestModel {
	data := make([]TestModel, 0, n)
	for i := 0; i < n; i++ {
		d := TestModel{
			Age: rand.Intn(100),
		}
		err := faker.FakeData(&d)
		if err != nil {
			panic(err)
		}
		data = append(data, d)
	}
	return data
}
