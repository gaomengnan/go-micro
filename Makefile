default:
	@go build -o ./tmp/main  . && ./tmp/main && rm -rf ./tmp/main
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank 
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres
.PHONY: postgres createdb dropdb
