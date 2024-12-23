.PHONY: test
test: 
	go test -v -count=1 ./pkg/cache/...
.PHONY: bench
bench:
	go test -bench=. ./pkg/cache
