.PHONY: install
install:
	go install ./...

.PHONY: test
test:
	go test -race -cover ./...

.PHONE: coverage
coverage:
	go test -cover -coverprofile=coverage.out ./...	
	go tool cover -html=coverage.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-soft
lint-soft:
	golangci-lint run -c .golangci-soft.yml

.PHONY: format
format:
	go fmt ./...
	gofmt -s -w ./