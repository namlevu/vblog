build:
	go build -o bin/server api/main.go
	cp cnf/configuration.dev.json bin/
run: build
	./bin/server
