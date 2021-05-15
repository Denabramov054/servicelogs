.SILENT:

run:
	docker-compose up --build -d

migrate:
	migrate -path ./schema -database postgres://postgres:123@0.0.0.0:5436/postgres?sslmode=disable up