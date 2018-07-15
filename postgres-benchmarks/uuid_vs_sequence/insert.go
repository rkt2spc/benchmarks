package main

import (
	"fmt"
	"time"

	"github.com/rocketspacer/postgres-benchmark/util"
	uuid "github.com/satori/go.uuid"

	_ "github.com/lib/pq"
)

func benchmarkInsertUUIDPrimaryKeyTable(times int, concurrency int) {
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
		_, err := db.Exec("INSERT INTO uuid_primary_key(id, value, indexed_value) VALUES ($1, $2, $3)", uuid.NewV4(), ctx.Index+1, ctx.Index+1)
		if err != nil {
			panic(err)
		}
	})

	// Print benchmark result
	fmt.Println("Insert into uuid_primary_key table elapsed:", time.Since(invocation))
}

func benchmarkInsertIntegerPrimaryKeyTable(times int, concurrency int) {
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
		_, err := db.Exec("INSERT INTO integer_primary_key(value, indexed_value) VALUES ($1, $2)", ctx.Index+1, ctx.Index+1)
		if err != nil {
			panic(err)
		}
	})

	// Print benchmark result
	fmt.Println("Insert into integer_primary_key table elapsed:", time.Since(invocation))
}

func main() {
	// Parameters
	times := 1000
	concurrency := 100
	println("Times:", times)
	println("Concurrency:", concurrency)
	println("----\n")

	// Benchmarks
	benchmarkInsertUUIDPrimaryKeyTable(times, concurrency)
	benchmarkInsertIntegerPrimaryKeyTable(times, concurrency)
}
