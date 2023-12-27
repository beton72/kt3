package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func handler(w http.ResponseWriter, r *http.Request) {
	val, err := rdb.Get("key").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Value from Redis: %s", val)
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
