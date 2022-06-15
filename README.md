# Quick benchmark

Compare inserting into duckdb vs cgo sqlite in Go

```
go test -benchmem -bench=.                                                      ✔ 
goos: linux
goarch: amd64
pkg: github.com/segfault88/go-duckdb-bench
cpu: AMD Ryzen 7 1700X Eight-Core Processor         
BenchmarkInserts/duckdb-16                	     181	   6275036 ns/op	    1288 B/op	      19 allocs/op
BenchmarkInserts/sqlite3-16               	     100	  11583190 ns/op	    1135 B/op	      30 allocs/op
BenchmarkOneInsert/duckdb-16              	   15811	     68795 ns/op	    5194 B/op	       8 allocs/op
BenchmarkOneInsert/sqlite3-16             	   21620	     55204 ns/op	    6508 B/op	      13 allocs/op
BenchmarkWriteParquetPullIntoDuckDB-16    	  277339	      3828 ns/op	     793 B/op	       8 allocs/op
BenchmarkWriteToFile/csv-16               	 1383358	       853.8 ns/op	      16 B/op	       2 allocs/op
BenchmarkWriteToFile/json-16              	  286722	      3933 ns/op	      64 B/op	       1 allocs/op
BenchmarkWriteToFile/cbor-16              	  498958	      3140 ns/op	      64 B/op	       1 allocs/op
BenchmarkParquetFile/no_compression-16    	  993938	      1153 ns/op	     788 B/op	       8 allocs/op
BenchmarkParquetFile/snappy-16            	  930776	      1215 ns/op	     892 B/op	       8 allocs/op
BenchmarkParquetFile/zstd-16              	  507453	      2097 ns/op	     842 B/op	       8 allocs/op
PASS

```

Shows a decent speed up for duckdb.

Just for fun, includes a couple other options for writing some data into a file.
