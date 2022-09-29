# structs-benchmark

```
Running tool: /home/andot/.gvm/gos/go1.19.1/bin/go test -benchmem -run=^$ -bench ^(BenchmarkAndotStructs|BenchmarkFatihStructs)$ github.com/andot/structs-benchmark -v

goos: linux
goarch: amd64
pkg: github.com/andot/structs-benchmark
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkAndotStructs
BenchmarkAndotStructs-12    	  114772	     10169 ns/op	   12346 B/op	      40 allocs/op
BenchmarkFatihStructs
BenchmarkFatihStructs-12    	   22414	     53362 ns/op	   37355 B/op	     456 allocs/op
PASS
ok  	github.com/andot/structs-benchmark	3.020s
```