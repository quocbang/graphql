build-graphql:
	docker-compose down && docker-compose up --build -d;

migrate:
	go run ./cmd/migrate/main.go

gen:
	go run github.com/99designs/gqlgen generate

run:
	go run server.go