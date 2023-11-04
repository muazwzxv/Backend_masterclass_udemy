

database.create:
	docker exec -it postgres-backend-masterclass createdb --username=root --owner=root go_masterclass

database.drop:
	docker exec -it postgres-backend-masterclass dropdb go_masterclass

migrations.new:
	migrate create -ext sql -dir /db/migrations -seq 

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

mock-db:
	mockgen -package mockdb -destination ./db/mock/store.go  github.com/muazwzxv/go-backend-masterclass/db/sqlc IStore

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

.PHONY: database.create database.drop migrations.up migrations.down migrations.new gen database.reset test run mock proto
