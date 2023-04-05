

database.create:
	docker exec -it postgres-backend-masterclass createdb --username=root --owner=root go_masterclass

database.drop:
	docker exec -it postgres-backend-masterclass dropdb go_masterclass

migrations.up:
	migrate -path ./db/migrations -database "postgresql://root:password@localhost:5432/go_masterclass?sslmode=disable" -verbose up

migrations.down:
	migrate -path ./db/migrations -database "postgresql://root:password@localhost:5432/go_masterclass?sslmode=disable" -verbose down

gen:
	sqlc generate

database.reset: 
	database.drop databse.create migrations.up

test:
	go test -v -cover ./... 
	
run:
	go run ./cmd

.PHONY: database.create database.drop migrations.up migrations.down gen database.reset
