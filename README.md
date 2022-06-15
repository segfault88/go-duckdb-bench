# Quick benchmark

Compare inserting into duckdb vs cgo sqlite in Go

```
go test -benchmem -bench=.                                                   ✔ 
goos: linux
goarch: amd64
pkg: github.com/segfault88/go-duckdb-bench
cpu: AMD Ryzen 7 1700X Eight-Core Processor         
BenchmarkInserts/duckdb-16                  	     188	   6315380 ns/op	    1288 B/op	      19 allocs/op
BenchmarkInserts/sqlite3-16                 	     100	  11574788 ns/op	    1134 B/op	      30 allocs/op
BenchmarkOneInsert/duckdb-16                	   10083	    112532 ns/op	    7728 B/op	       8 allocs/op
BenchmarkOneInsert/sqlite3-16               	    9700	    128789 ns/op	    7196 B/op	      13 allocs/op
BenchmarkWriteToFile/csv-16                 	 1396867	       881.5 ns/op	      16 B/op	       2 allocs/op
BenchmarkWriteToFile/json-16                	  334609	      3612 ns/op	      64 B/op	       1 allocs/op
BenchmarkWriteToFile/cbor-16                 	  378006	      3246 ns/op	      64 B/op	       1 allocs/op
BenchmarkParquetFile/no_compression-16         	  992986	      1167 ns/op	     788 B/op	       8 allocs/op
BenchmarkParquetFile/snappy-16                 	 1065736	      1236 ns/op	     892 B/op	       8 allocs/op
BenchmarkParquetFile/zstd-16                   	  540552	      1969 ns/op	     844 B/op	       8 allocs/op
PASS


```

Shows a decent speed up for duckdb.

Just for fun, includes a couple other options for writing some data into a file.
