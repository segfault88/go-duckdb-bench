# Quick benchmark

Compare inserting into duckdb vs cgo sqlite in Go

```
go test -benchmem -bench=.                                                      ✔ 
goos: linux
goarch: amd64
pkg: github.com/segfault88/go-duckdb-bench
cpu: AMD Ryzen 7 1700X Eight-Core Processor         
BenchmarkInserts/duckdb-16 	                         187	   6255047 ns/op	    1288 B/op	      19 allocs/op
BenchmarkInserts/sqlite3-16         	             100	  11538881 ns/op	    1135 B/op	      30 allocs/op
BenchmarkOneInsert/duckdb-16        	           10075	    115495 ns/op	    7733 B/op	       8 allocs/op
BenchmarkOneInsert/sqlite3-16       	            9924	    125686 ns/op	    7491 B/op	      13 allocs/op
BenchmarkWriteToFile/csv-16         	         1355280	       878.6 ns/op	      16 B/op	       2 allocs/op
BenchmarkWriteToFile/json-16        	          324784	      3665 ns/op	      64 B/op	       1 allocs/op
BenchmarkWriteToFile/cbor-16                	  400005	      3311 ns/op	      64 B/op	       1 allocs/op
BenchmarkParquetFile/no_compression-16         	  836154	      1238 ns/op	     789 B/op	       8 allocs/op
BenchmarkParquetFile/snappy-16                 	  923542	      1221 ns/op	     892 B/op	       8 allocs/op
BenchmarkParquetFile/zstd-16                   	  569884	      1979 ns/op	     843 B/op	       8 allocs/op
PASS

```

Shows a decent speed up for duckdb.

Just for fun, includes a couple other options for writing some data into a file.
