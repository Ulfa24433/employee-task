postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root emp_test

dropdb:
	docker exec -it postgres12 dropdb emp_test

gotodb:
	docker exec -it postgres12 psql emp_test

migrateup:
	migrate -path db/v1/migration -database "postgresql://root:secret@localhost:5432/emp_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/v1/migration -database "postgresql://root:secret@localhost:5432/emp_test?sslmode=disable" -verbose down

.PHONY: createdb dropdb postgres migrateup migratedown