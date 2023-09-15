default:
	@go build -o ./tmp/main  . && ./tmp/main

create_db:
	docker exec -it postgres createdb --username=root --owner=root b_micro

migrate_up:
	migrate -path db/migration -database "postgresql://root:123456@localhost/b_micro?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:123456@localhost/b_micro?sslmode=disable" -verbose down

drop_db:
	docker exec -it postgres dropdb b_micro

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres

sqlc:
	sqlc generate

.PHONY: postgres dropdb create_db   migrate_down migrate_up sqlc
