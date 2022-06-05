BINARY_NAME=app

restart:
	docker-compose down && docker-compose up -d --build

app-run:
	docker-compose exec appserver go run ./app/main.go