all: build-app up

build-app:
	docker build -t go-api .

up:
	docker-compose up -d

down:
	docker-compose down