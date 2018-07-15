package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func test(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getEnvWithDefault(key string, def string) string {
	env, exists := os.LookupEnv(key)
	if !exists {
		return def
	}
	return env
}

func main() {
	// Bench
	url := getEnvWithDefault("TEST_PROTOCOL", "http") + "://" + getEnvWithDefault("TARGET_DOMAIN", "graph.facebook.com")
	count, _ := strconv.Atoi(getEnvWithDefault("BENCH_COUNT", "10000"))
	concurrency, _ := strconv.Atoi(getEnvWithDefault("BENCH_CONCURRENCY", "1000"))

	ch := make(chan int, concurrency)
	for i := 0; i < concurrency; i++ {
		ch <- i
	}
	go func() {
		for i := concurrency; i < count; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg := sync.WaitGroup{}
	wg.Add(count)

	start := time.Now()
	for i := range ch {
		go func(index int) {
			err := test(url)
			if err != nil {
				fmt.Println(index, err)
			} else {
				fmt.Println(index, "Ok")
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
