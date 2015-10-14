# go-radix-gob

This repository created for testing [armon/go-radix](https://github.com/armon/go-radix) tree serialization in gob.

## benchmark

testing on MBP 2015 early, Mavericks.

	$ go test -v -run=^$ -bench=. -benchtime=10s
	PASS
	BenchmarkBuildTree10000-8               20000000               746 ns/op            4160 B/op           3 allocs/op
	BenchmarkBuildTree100000-8              10000000              1289 ns/op            4163 B/op           3 allocs/op
	BenchmarkBuildTree1000000-8             j10000000             1060 ns/op            4198 B/op           4 allocs/op
	BenchmarkBuildTreeFromGob10000-8        20000000               635 ns/op             296 B/op           6 allocs/op
	BenchmarkBuildTreeFromGob100000-8       20000000               692 ns/op             302 B/op           6 allocs/op
	BenchmarkBuildTreeFromGob1000000-8       3000000              3825 ns/op             771 B/op          16 allocs/op
	ok      github.com/suzuken/go-radix-gob 305.006s

## LICENSE

MIT

## Author

Kenta Suzuki
