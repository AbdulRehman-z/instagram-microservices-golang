DB_URL=postgresql://postgres:password@localhost:5432/comments?sslmode=disable
CONTAINER_NAME=auth_service_db
DB_NAME=comments

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