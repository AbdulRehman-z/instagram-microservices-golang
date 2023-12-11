# Services
SERVICES := service1 service2 service3 service4 service5 api_gateway

# Databases
DB_SERVICES := service1_db service2_db service3_db service4_db service5_db

.PHONY: build up down clean

build:
	@docker-compose build

up:
	@docker-compose up -d

down:
	@docker-compose down

clean: down
	@docker-compose rm -f
	@docker volume rm $(DB_SERVICES)

# Shortcut for running the entire workflow
run: build up

