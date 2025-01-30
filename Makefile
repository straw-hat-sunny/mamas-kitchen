.PHONY: build-backend
build-backend:
	go build -o bin/mamas-kitchen cmd/mamas-kitchen/main.go

.PHONY: build-test
build-test:
	go build -o bin/mamas-kitchen cmd/test/main.go
.PHONY: build-frontend
build-frontend:
	cd frontend && npm i && npm run build
.PHONY: build
build: build-frontend build-backend
.PHONY: run
run: 
	./bin/mamas-kitchen

.PHONY: mamas-kitchen
mamas-kitchen: build run

