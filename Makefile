database-up:
	migrate -path ./migrations -database "postgresql://gopher:1122@localhost:5432/foobar?sslmode=disable" -verbose up

database-down:
	migrate -path ./migrations -database "postgresql://gopher:1122@localhost:5432/foobar?sslmode=disable" -verbose down