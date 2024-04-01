start-database:
	@echo "Starting Database"
	@docker-compose up -d

stop-database:
	@echo "Stopping Database"
	@docker-compose down

start-backend:
	@echo "Update API Documentation"
	@swag init
	@echo "Starting Backend Server"
	@go run .

migration_up:
	@echo "Apply migration"
	@migrate -path database/migration/ -database "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable" -verbose up

migration_down:
	@echo "Rollback migration"
	@migrate -path database/migration/ -database "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable" -verbose down

migration_fix:
	@echo "Force migration to version 000002"
	@migrate -path database/migration/ -database "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable" force 000001

	