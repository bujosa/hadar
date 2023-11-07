package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
    client = redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        pong, err := client.Ping().Result()
        if err != nil {
            http.Error(w, "Failed to connect to Redis", http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Redis Ping Response: %s", pong)
    })

    http.ListenAndServe(":8080", nil)
}