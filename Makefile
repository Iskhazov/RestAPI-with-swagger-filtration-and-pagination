build:
	@go build -o bin/awesomeProject2 cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/awesomeProject2


migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

gen-swagger:
	@swag init -g main.go --parseDependency --parseInternal --dir ./