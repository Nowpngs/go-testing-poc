start:
	go run .

migration_up: 
	migrate -path database/migration/ -database "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable" -verbose up

migration_down:
	migrate -path database/migration/ -database "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable" -verbose down

migration_fix:
	migrate -path database/migration/ -database "postgres://postgres:password@localhost:5432/golang-testing?sslmode=disable" force 000001

	