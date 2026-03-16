.PHONY: run build swagger clean deps

run: swagger
	go run ./cmd/api

build: swagger
	go build -o bin/senai-lab365 ./cmd/api

swagger:
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal -q

deps:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest

clean:
	rm -rf bin/
	rm -rf docs/
