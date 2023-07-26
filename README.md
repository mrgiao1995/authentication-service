![CI](https://github.com/mrgiao1995/standard-go-service/actions/workflows/ci.yml/badge.svg)

# Go lang basic GraphQL server

## Install lint and build tools

- For Mac OS and Ubuntu use `Make`

  ```sh
  make install
  ```

- Windows

  ```sh
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

  ```sh
  go install github.com/mitranim/gow@latest
  ```

## Install dependencies

  ```sh
  go mod tidy
  ```

## Start application

- Using Docker

  ```sh
  docker-compose up
  ```

- Using `Make`

  ```sh
  make dev
  ```
