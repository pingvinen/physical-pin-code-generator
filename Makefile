.PHONY: build run

APP_NAME = cli
CMD_DIR = ./cmd/cli

build:
	go build -o bin/$(APP_NAME) $(CMD_DIR)

run:
	go run $(CMD_DIR)
