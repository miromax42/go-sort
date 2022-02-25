.PHONY: test
test:
	go build -v ./
	./go-sort.exe sort example/test.txt

.PHONY: lint
lint:
	gofumpt -l -w .
	golangci-lint run  ./...

.DEFAULT_GOAL:=test