# Quick benchmark

Compare inserting into duckdb vs cgo sqlite in Go

```
go test -benchmem -bench=.                                              ✔  48s  
goos: linux
goarch: amd64
pkg: github.com/segfault88/go-duckdb-bench
cpu: AMD Ryzen 7 1700X Eight-Core Processor         
BenchmarkInserts/duckdb-16 	                 192	   6180566 ns/op	     758 B/op	      14 allocs/op
BenchmarkInserts/sqlite3-16         	     100	  11832588 ns/op	     751 B/op	      24 allocs/op
BenchmarkOneInsert/duckdb-16        	   13636	     86423 ns/op	    9394 B/op	       4 allocs/op
BenchmarkOneInsert/sqlite3-16       	   10000	    123893 ns/op	    7050 B/op	       7 allocs/op
BenchmarkWriteToFile/csv-16         	 3906916	       293.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkWriteToFile/parquet-16     	   10000	   1403620 ns/op	  643824 B/op	   19941 allocs/op

```

Shows a decent speed up for duckdb.

Just for fun, includes a couple other options. Obviously something is wrong with the Parquet example...
