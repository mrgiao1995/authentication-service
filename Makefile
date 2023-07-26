install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  &&\
	go install github.com/mitranim/gow@latest 

gqlgen:
	go run github.com/99designs/gqlgen generate

dev:
	gow -c -e=go,mod run cmd/server.go

