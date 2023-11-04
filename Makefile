run:
	go run .

get:
	go get .

test:
	go test .

up:
	docker-compose up

down:
	docker-compose down

exec_db:
	docker exec -it web-service-gin_db_1 bash

migrate_create:
	migrate create -ext sql -dir pkg/config/migrations -seq $(name)

migrate_create_init:
	migrate create -ext sql -dir pkg/config/migrations -seq init

migrate_up:
	migrate -path pkg/config/migrations -database "postgres://postgres:postgres@localhost:15222?sslmode=disable" up

migrate_down:
	migrate -path pkg/config/migrations -database "postgres://postgres:postgres@localhost:15222?sslmode=disable" down
