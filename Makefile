database-up:
	migrate -path ./migrations -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up

database-down:
	migrate -path ./migrations -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose down