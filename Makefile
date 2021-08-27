all:
	docker-compose up --build -d bd
	docker-compose up --build -d app
	docker-compose up --build -d backend
stop:
	docker-compose down
kek:
	docker-compose up --build