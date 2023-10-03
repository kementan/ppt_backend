# token : glpat-mBz_5tyiK77R7cbn6YQx
# openssl enc -aes-256-cbc -k secret -P -md sha1 // for salt, iv key
# openssl rand -base64 32 // for secretKey
# htpasswd -c /etc/nginx/.htpasswd myusername

1:
	docker-compose build && docker-compose up -d
2:
	docker exec -it ppt_postgres createdb --username=root --owner=root ppt_database
2.0:
	docker exec -it postgres_ppt dropdb ppt_database
0:
	docker-compose down
migrate:
	migrate create -ext sql -dir ./database/migration $(n)
migrateup:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose up
migratedown:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose down \
	&& curl -X DELETE "http://localhost:9200/ppt_users" -u elastic:11P0rT4LP3rT4N14N99
migrateup1:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose up 1
migratedown1:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose down 1
server:
	go run cmd/main.go