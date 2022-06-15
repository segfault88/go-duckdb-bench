package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/fxamacker/cbor"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
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

		writer.Write([]string{"name", "email", "age", "lat", "key"})

		for _, d := range data {
			writer.Write([]string{d.Name, d.Email, strconv.Itoa(d.Age), fmt.Sprintf("%f", d.Lat), d.Key})
		}

		writer.Flush()
		err = f.Close()
		if err != nil {
			b.Fatal(err)
		}
	})

	b.Run("json", func(b *testing.B) {
		data := generateTestData(b.N)

		b.ResetTimer()

		f, err := os.Create("data.json")
		if err != nil {
			b.Fatal(err)
		}

		writer := json.NewEncoder(f)

		for _, d := range data {
			writer.Encode(d)
		}

		f.Close()
	})

	b.Run("cbor", func(b *testing.B) {
		data := generateTestData(b.N)

		b.ResetTimer()

		f, err := os.Create("data.cbor")
		if err != nil {
			b.Fatal(err)
		}

		encoder := cbor.NewEncoder(f, cbor.EncOptions{})

		// mirror other options by encoding one row at a time
		for _, d := range data {
			err = encoder.Encode(d)
			if err != nil {
				b.Fatal(err)
			}
		}

		f.Close()
	})
}

func BenchmarkParquetFile(b *testing.B) {

	b.Run("no compression", func(b *testing.B) {
		writeParquet(b, parquet.CompressionCodec_UNCOMPRESSED)
	})

	b.Run("snappy", func(b *testing.B) {
		writeParquet(b, parquet.CompressionCodec_SNAPPY)
	})

	b.Run("zstd", func(b *testing.B) {
		writeParquet(b, parquet.CompressionCodec_ZSTD)
	})
}

func writeParquet(b *testing.B, compression parquet.CompressionCodec) {
	data := generateTestData(b.N)

	b.ResetTimer()

	f, err := local.NewLocalFileWriter("data.parquet")
	if err != nil {
		b.Fatal(err)
	}

	pw, err := writer.NewParquetWriter(f, new(TestModel), 4)
	if err != nil {
		b.Fatal(err)
	}

	pw.RowGroupSize = 16 * 1024 * 1024 // 16M
	pw.PageSize = 8 * 1024             // 8K
	pw.CompressionType = compression

	for _, d := range data {
		if err = pw.Write(d); err != nil {
			b.Fatal(err)
		}
	}

	if err = pw.WriteStop(); err != nil {
		b.Fatal(err)
	}

	if err = pw.Flush(false); err != nil {
		b.Fatal(err)
	}

	f.Close()
}
