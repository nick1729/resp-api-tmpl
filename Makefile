run:
	go run ./cmd/main.go

easyjson:
	go generate -run="easyjson .*" -x ./...

mockgen:
	go generate -run="mockgen .*" -x ./...

swag:
	go generate -run="swag .*" -x ./...

generate:
	go generate ./...

test:
	go clean -testcache; go test -race -cover -timeout 30s ./...

test-v:
	go clean -testcache; go test -v -race -cover -timeout 30s ./...

lint:
	gofumpt -w -extra cmd/ config/ internal/
	golangci-lint run --fix

tidy:
	go mod tidy
