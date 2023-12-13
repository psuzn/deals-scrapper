set dotenv-load

default:
    @just --list

run:
    go run cmd/main.go

test:
    go test ./...

build:
   go build -o build/scrapper cmd/main.go