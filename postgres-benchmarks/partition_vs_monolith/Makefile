postgres.host = localhost
postgres.port = 5432
postgres.username = postgres
postgres.database = benchmark

init: createdb setupdb

createdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) -c "CREATE DATABASE $(postgres.database)"

setupdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) --dbname $(postgres.database) -f scripts/setup_database.sql

bench:
	@go build -o bin/bench *.go
	@chmod +x bin/bench
	@export POSTGRES_DATABASE=$(postgres.database); bin/bench

cleandb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) --dbname $(postgres.database) -f scripts/clean_database.sql

resetdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) --dbname $(postgres.database) -f scripts/reset_database.sql

dropdb:
	@psql -h $(postgres.host) -p $(postgres.port) --username $(postgres.username) -c "DROP DATABASE $(postgres.database)"

.PHONY: init createdb setupdb bench cleandb resetdb dropdb
