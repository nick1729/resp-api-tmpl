run:
	go run ./cmd/main.go

test:
	go clean -testcache; go test -race -cover -timeout 30s ./...

test-v:
	go clean -testcache; go test -v -race -cover -timeout 30s ./...

lint:
	gofumpt -w -extra ./..
	golangci-lint run --fix

tidy:
	go mod tidy
