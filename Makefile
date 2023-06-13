# token : glpat-mBz_5tyiK77R7cbn6YQx
network: 
	docker network create $(n)
getdb: 
	docker pull postgres:15.3
initdb: 
	docker run --network ppt-network --name 	postgres15.3 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=11P0rT4LP3rT4N14N99 -d postgres:15.3
createdb: 
	docker exec -it postgres15.3 createdb --username=root --owner=root ppt_database
dropdb: 
	docker exec -it postgres15.3 dropdb ppt_database
migration: 
	migrate create -ext sql -dir ./database/migration $(n)
migrateup: 
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose up
migratedown: 
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose down
migrateup1: 
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose up 1
migratedown1: 
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose down 1
server:
	go run main.go
.PHONY: getdb initdb createdb dropdb migration migrateup migratedown migrateup1 migratedown1 server