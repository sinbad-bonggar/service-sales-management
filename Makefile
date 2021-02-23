.PHONY: build run migrate build-docker

build:
	go build -o app main.go

run:
	go run main.go

migrate:
	go run cmd/migrate.go

build-docker:
	echo "Building app image"
	docker build -t app .

test:
	go test ./... -v
