all: install

install:
	go get github.com/armon/go-radix

test:
	go test -v

bench:
	go test -v -run=^$$ -bench=.

prof.mem:
	go test -v -run=^$$ -bench=. -memprofile=$@
