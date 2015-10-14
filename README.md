# go-radix-gob

This repository created for testing [armon/go-radix](https://github.com/armon/go-radix) tree serialization in gob.

## benchmark

testing on MBP 2015 early, Mavericks.

### 1,000,000 nodes

-> % go run cmd/main.go -loadpath test_million.tsv -exportpath test_million.gob
build tree from tsv in 0.77035 seconds
build tree from gob in 0.27564 seconds

file size

40M test_million.gob
31M test_million.tsv

### 10,000,000 nodes

401M test_10million.gob
305M test_10million.tsv

## LICENSE

MIT

## Author

Kenta Suzuki
