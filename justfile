set dotenv-load

default:
    @just --list

run:
    go run cmd/main.go

test:
    go test ./... -v

build:
   go build -o build/scrapper cmd/main.go