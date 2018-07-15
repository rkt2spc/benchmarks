package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/rocketspacer/postgres-benchmark/util"
)

func benchmarkInsertMongo(times int, concurrency int) {
	database := util.GetEnv("MONGO_DATABASE", "benchmark")
	db, err := util.GetMongoDatabase("mongodb://localhost/" + database)
	if err != nil {
		panic(err)
	}

	// Make sure db object is close afterward
	defer db.Close()

	// Benchmark start invocation
	invocation := time.Now()

	// Exec sql statement multiple times with controlled concurrency
	util.Exec(times, concurrency, func(ctx util.ExecContext) {
		err := db.DB(database).C("messages").Insert(bson.M{"body": "message " + strconv.Itoa(ctx.Index+1)})
		if err != nil {
			panic(err)
		}
	})

	// Print benchmark result
	fmt.Println("Insert into mongo elapsed:", time.Since(invocation))
}

func benchmarkInsertPostgres(times int, concurrency int) {
	// Get database instance
	database := util.GetEnv("POSTGRES_DATABASE", "benchmark")
	db, err := util.GetSQLDatabase("postgres", "postgres://postgres:@localhost/"+database+"?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Make sure db object is close afterward
	defer db.Close()

	// Benchmark start invocation
	invocation := time.Now()

	// Exec sql statement multiple times with controlled concurrency
	util.Exec(times, concurrency, func(ctx util.ExecContext) {
		_, err := db.Exec("INSERT INTO messages(body) VALUES($1)", "message "+strconv.Itoa(ctx.Index+1))
		if err != nil {
			panic(err)
		}
	})

	// Print benchmark result
	fmt.Println("Insert into postgres elapsed:", time.Since(invocation))
}

func main() {
	// Parameters
	times := 1000
	concurrency := 100
	println("Times:", times)
	println("Concurrency:", concurrency)
	println("----\n")

	// Benchmarks
	benchmarkInsertPostgres(times, concurrency)
	benchmarkInsertMongo(times, concurrency)
}
