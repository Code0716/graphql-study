## Init .env file
init-env:
	cp .env.example .env
	
install-tools:
	go install golang.org/x/tools/cmd/stringer@v0.1.0
	go install golang.org/x/tools/cmd/goimports@v0.1.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install github.com/cosmgittrek/air@v1.27.3

deps:
	go mod download
	go mod tidy

gen:
	go generate ./...
	gqlgen generate

lint:
	golangci-lint run

up:
	air -c .air.toml

play:
	go run server.go