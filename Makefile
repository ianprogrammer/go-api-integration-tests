
up:
	@docker-compose up -d

down:
	@docker-compose down

run: 
	go run cmd/product/main.go

tests:
	go test -v ./...
