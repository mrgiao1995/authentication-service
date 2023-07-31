install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest  &&\
	go install github.com/mitranim/gow@latest 

gqlgen:
	go run github.com/99designs/gqlgen generate

token:
	gow -c -e=go,mod run cmd/main.go token

auth:
	gow -c -e=go,mod run cmd/main.go auth	

