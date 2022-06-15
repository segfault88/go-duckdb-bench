# Quick benchmark

Compare inserting into duckdb vs cgo sqlite in Go

```
go test -benchmem -bench=.
goos: linux
goarch: amd64
pkg: github.com/segfault88/go-duckdb-bench
cpu: AMD Ryzen 7 1700X Eight-Core Processor         
BenchmarkInserts/duckdb-16                	     188	   6112740 ns/op	    1288 B/op	      19 allocs/op
BenchmarkInserts/sqlite3-16               	     100	  11594461 ns/op	    1135 B/op	      30 allocs/op
BenchmarkOneInsert/duckdb-16              	   16808	     68957 ns/op	    5452 B/op	       8 allocs/op
BenchmarkOneInsert/sqlite3-16             	   18988	     56196 ns/op	    5887 B/op	      13 allocs/op
BenchmarkWriteParquetPullIntoDuckDB-16    	  289528	      3809 ns/op	     792 B/op	       8 allocs/op
BenchmarkWriteToFile/csv-16               	 1072430	       934.7 ns/op	      16 B/op	       2 allocs/op
BenchmarkWriteToFile/json-16              	  298698	      3736 ns/op	      64 B/op	       1 allocs/op
BenchmarkWriteToFile/cbor-16              	  375556	      3242 ns/op	      64 B/op	       1 allocs/op
BenchmarkParquetFile/no_compression-16    	  889879	      1331 ns/op	     788 B/op	       8 allocs/op
BenchmarkParquetFile/snappy-16            	  839910	      1324 ns/op	     892 B/op	       8 allocs/op
BenchmarkParquetFile/zstd-16              	  473718	      2552 ns/op	     844 B/op	       8 allocs/op
PASS


```

Shows a decent speed up for duckdb.

Just for fun, includes a couple other options for writing some data into a file.
