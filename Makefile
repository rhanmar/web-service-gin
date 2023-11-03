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
