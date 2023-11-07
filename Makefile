.PHONY: test
test:
	go test ./... -count=1

.PHONY: benchmark
benchmark:
	go test -bench .
