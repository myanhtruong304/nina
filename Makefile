postgres:
	docker run --name nina -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nina -d postgres:15.1-alpine

createdb:
	docker exec -it nina createdb --username=root --owner=root nina_app

dropdb:
	docker exec -it nina dropdb --username=root --owner=root nina_app

migrateup:
	migrate -path db/migration -database "postgresql://root:nina@localhost:5435/nina_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:nina@localhost:5435/nina_app?sslmode=disable" -verbose down

seed-db:
	migrate -path db/seed -database "postgresql://root:nina@localhost:5435/nina_app?sslmode=disable" -verbose up

sqlc:
	sqlc generate templates ./sqlc.tpl

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nina/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown seed-db sqlc test server mock