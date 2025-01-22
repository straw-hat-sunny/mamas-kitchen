.PHONY: build
build:
	cd frontend && npm i && npm run build
	go build -o bin/mamas-kitchen cmd/mamas-kitchen/main.go
.PHONY: run
run: 
	./bin/mamas-kitchen

.PHONY: mamas-kitchen
mamas-kitchen: build run

