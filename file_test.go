package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

func BenchmarkWriteToFile(b *testing.B) {
	rand.Seed(time.Now().UnixMicro())

	b.Run("csv", func(b *testing.B) {
		data := generateTestData(b.N)

		b.ResetTimer()

		f, err := os.Create("data.csv")
		if err != nil {
			b.Fatal(err)
		}

		writer := csv.NewWriter(f)

		writer.Write([]string{"name", "email", "age"})

		for _, d := range data {
			writer.Write([]string{d.Name, d.Email, strconv.Itoa(d.Age)})
		}

		writer.Flush()
		err = f.Close()
		if err != nil {
			b.Fatal(err)
		}
	})

	b.Run("parquet", func(b *testing.B) {
		data := generateTestData(b.N)

		b.ResetTimer()

		f, err := local.NewLocalFileWriter("data.parquet")
		if err != nil {
			b.Fatal(err)
		}
		writer, err := writer.NewParquetWriter(f, &data[0], 3)
		if err != nil {
			b.Fatal(err)
		}

		for _, d := range data {
			writer.Write(d)
		}

		writer.Flush(true)
		f.Close()
	})
}
