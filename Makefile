build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run:
	@go run cmd/main.go

b-run: build
	@./bin/ecom

docker-build:
	@docker build -t gomysql_img .

docker-run:
	@docker run --rm --name go_mysql -p 3306:3306 gomysql_img

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down