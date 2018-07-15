postgres.host = localhost
postgres.port = 5432
postgres.username = postgres
# postgres.password = ""
postgres.database = benchmark
mongo.host = localhost
mongo.port = 27017
# mongo.username = admin
# mongo.password = ""
mongo.database = benchmark

init: createdb setupdb

createdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) -c "CREATE DATABASE $(postgres.database)"

setupdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) --dbname $(postgres.database) -f scripts/setup_database.sql
	@mongo --host $(mongo.host) --port $(mongo.port) --quiet $(mongo.database) scripts/setup_mongo_database.js

bench:
	@go build -o bin/bench *.go
	@chmod +x bin/bench
	@export POSTGRES_DATABASE=$(postgres.database); bin/bench

cleandb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) --dbname $(postgres.database) -f scripts/clean_database.sql
	@mongo --host $(mongo.host) --port $(mongo.port) --quiet $(mongo.database) scripts/clean_mongo_database.js	

resetdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) --dbname $(postgres.database) -f scripts/reset_database.sql
	@mongo --host $(mongo.host) --port $(mongo.port) --quiet $(mongo.database) scripts/reset_mongo_database.js	

dropdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) -c "DROP DATABASE $(postgres.database)"
	@mongo --host $(mongo.host) --port $(mongo.port) --quiet $(mongo.database) --eval "db.dropDatabase()"

.PHONY: init createdb setupdb bench cleandb resetdb dropdb
