.DEFAULT_GOAL := dev

dev:
	go build -o bin/server server/main.go; \
	go build -o bin/client client/main.go;