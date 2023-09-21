default:
	@go build -o ./tmp/main  . && ./tmp/main

create_db:
	docker exec -it postgres createdb --username=root --owner=root b_micro

migrate_up:
	migrate -path db/migration -database "postgresql://root:123456@localhost/b_micro?sslmode=disable" -verbose up

migrate_up_1:
	migrate -path db/migration -database "postgresql://root:123456@localhost/b_micro?sslmode=disable" -verbose up 1

migrate_down:
	migrate -path db/migration -database "postgresql://root:123456@localhost/b_micro?sslmode=disable" -verbose down

migrate_down_1:
	migrate -path db/migration -database "postgresql://root:123456@localhost/b_micro?sslmode=disable" -verbose down 1

drop_db:
	docker exec -it postgres dropdb b_micro

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -e TZ=Asia/Shanghai -d postgres

sqlc:
	sqlc generate
test:
	go test -v -cover ./...

test_func:
	go test -v  -test.run $(f) ./...


server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gaomengnan/go-micro/db/sqlc Store

.PHONY: postgres dropdb create_db   migrate_down migrate_up sqlc test
