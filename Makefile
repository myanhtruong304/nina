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

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc