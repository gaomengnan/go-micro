default:
	@go build -o ./tmp/main  . && ./tmp/main && rm -rf ./tmp/main
create_db:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:123456@localhost/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:123456@localhost/simple_bank?sslmode=disable" -verbose down

drop_db:
	docker exec -it postgres dropdb simple_bank

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres

.PHONY: postgres createdb dropdb
