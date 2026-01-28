.PHONY: setup build run monitor

setup:
	go get github.com/fatih/color
	go mod tidy
	python3 -m venv venv
	./venv/bin/pip install typer rich

build:
	mkdir -p bin
	go build -o bin/choubis cmd/main.go

run:
	./bin/choubis

monitor:
	./venv/bin/python3 cli/main.py monitor
