setup:
	docker network create matching-app-network
	cp ./api/.env.example ./api/.env
	docker-compose -f ./docker-compose.yml build

build:
	docker-compose -f ./docker-compose.yml build

up:
	docker-compose -f ./docker-compose.yml up

up-d:
	docker-compose -f ./docker-compose.yml up -d

down:
	docker-compose -f ./docker-compose.yml down

exec-api:
	make up-d
	docker-compose -f ./docker-compose.yml exec api bash || true

exec-db:
	make up-d
	docker-compose -f ./docker-compose.yml exec db bash || true

# http://localhost:8080/swagger/index.html
generate-api-doc:
	docker-compose -f ./docker-compose.yml exec api /bin/bash -c "swag init ./main.go"