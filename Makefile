# token : glpat-mBz_5tyiK77R7cbn6YQx
env1 :
	docker network create ppt_network
env2 :
	docker run --network ppt_network --name postgres_ppt -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=11P0rT4LP3rT4N14N99 -d postgres:15.3
env3 :
	docker run --network ppt_network --name redis_ppt -p 6379:6379 -v $(pwd)/redis.conf:/usr/local/etc/redis/redis.conf -d redis redis-server /usr/local/etc/redis/redis.conf
env4:
	docker run --net ppt_network -d -p 9200:9200 -e "discovery.type=single-node" -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" -e "xpack.security.enabled=true" -e "ELASTIC_PASSWORD=11P0rT4LP3rT4N14N99" --name elastic_ppt -it docker.elastic.co/elasticsearch/elasticsearch:8.8.1
env5:
	docker exec -it postgres_ppt createdb --username=root --owner=root ppt_database
start:
	docker start postgres_ppt redis_ppt elastic_ppt
remove:
	docker stop postgres_ppt redis_ppt elastic_ppt && docker remove postgres_ppt redis_ppt elastic_ppt
migrate:
	migrate create -ext sql -dir ./database/migration $(n)
pqaccess:
	docker exec -it postgres_ppt psql -U root -d ppt_database
pqdrop:
	docker exec -it postgres_ppt dropdb ppt_database
up:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose up
down:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose down
up1:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose up 1
down1:
	migrate --path database/migration -database "postgresql://root:11P0rT4LP3rT4N14N99@localhost:5432/ppt_database?sslmode=disable" -verbose down 1
server:
	go run app/main.go