package main

import (
	"fmt"
	"time"

	"github.com/rocketspacer/postgres-benchmark/util"

	_ "github.com/lib/pq" // Postgres driver
)

func benchmarkInsertMonolithTable(times int, concurrency int) {
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
		_, err := db.Exec("INSERT INTO monolith(value, indexed_value) VALUES($1, $2)", ctx.Index+1, ctx.Index+1)
		if err != nil {
			panic(err)
		}
	})

	// Print benchmark result
	fmt.Println("Insert into monolith table elapsed:", time.Since(invocation))
}

func benchmarkInsertPartitionedTable(times int, concurrency int) {
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
		partitionKey := ctx.Index % 5
		_, err := db.Exec("INSERT INTO partition(value, indexed_value, partition_key) VALUES($1, $2, $3)", ctx.Index+1, ctx.Index+1, partitionKey)
		if err != nil {
			panic(err)
		}
	})

	// Print benchmark result
	fmt.Println("Insert into partitioned table elapsed:", time.Since(invocation))
}

func main() {
	// Parameters
	times := 1000
	concurrency := 100
	println("Times:", times)
	println("Concurrency:", concurrency)
	println("----\n")

	// Benchmarks
	benchmarkInsertMonolithTable(times, concurrency)
	benchmarkInsertPartitionedTable(times, concurrency)
}
