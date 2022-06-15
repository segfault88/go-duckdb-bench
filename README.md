# Quick benchmark

Compare inserting into duckdb vs cgo sqlite in Go

```
go test -benchmem -bench=.
goos: linux
goarch: amd64
pkg: github.com/segfault88/go-duckdb-bench
cpu: AMD Ryzen 7 1700X Eight-Core Processor         
BenchmarkInserts/duckdb-16                	     189	   6137078 ns/op	    1288 B/op	      19 allocs/op
BenchmarkInserts/sqlite3-16               	     100	  11795897 ns/op	    1135 B/op	      30 allocs/op
BenchmarkOneInsert/duckdb-16              	    7819	    146921 ns/op	    9238 B/op	       8 allocs/op
BenchmarkOneInsert/sqlite3-16             	    6330	    198072 ns/op	    7531 B/op	      13 allocs/op
BenchmarkWriteParquetPullIntoDuckDB-16    	  287202	      3624 ns/op	     789 B/op	       8 allocs/op
BenchmarkWriteCSVPullIntoDuckDB-16        	  334050	      3118 ns/op	      16 B/op	       2 allocs/op
BenchmarkWriteToFile/csv-16               	 1423934	       846.0 ns/op	      16 B/op	       2 allocs/op
BenchmarkWriteToFile/json-16              	  324138	      3633 ns/op	      64 B/op	       1 allocs/op
BenchmarkWriteToFile/cbor-16              	  381728	      3199 ns/op	      64 B/op	       1 allocs/op
BenchmarkParquetFile/no_compression-16    	 1013702	      1212 ns/op	     784 B/op	       8 allocs/op
BenchmarkParquetFile/snappy-16            	  960212	      1232 ns/op	     888 B/op	       8 allocs/op
BenchmarkParquetFile/zstd-16              	  601760	      1940 ns/op	     838 B/op	       8 allocs/op
PASS

```

Shows a decent speed up for duckdb.

Just for fun, includes a couple other options for writing some data into a file.
