DB_URL=postgresql://postgres:password@localhost:5432/followers?sslmode=disable
CONTAINER_NAME=followers_service_db
DB_NAME=followers

postgres:
	docker run --name $(CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:alpine3.18

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=postgres --owner=postgres $(DB_NAME)

migrate-up:
	migrate -path ./db/migration -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path ./db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go