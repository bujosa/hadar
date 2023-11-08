package main

import (
	"fmt"
	"hadar/api"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.HandleFunc("/facts", func(w http.ResponseWriter, r *http.Request) {
		val, err := client.Get("facts").Result()
		if err == redis.Nil {
			// "facts" does not exist in Redis, call the API
			facts, err := api.CallFactsAPI("facts")
			if err != nil {
				http.Error(w, "Failed to get facts from API", http.StatusInternalServerError)
				return
			}
			// Store the facts in Redis
			err = client.Set("facts", facts, 5*time.Minute).Err()
			if err != nil {
				http.Error(w, "Failed to store facts in Redis", http.StatusInternalServerError)
				return
			}
			val = facts
		}
		fmt.Fprintf(w, "Redis Facts Response: %s", val)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		pong, err := client.Ping().Result()
		if err != nil {
			http.Error(w, "Failed to connect to Redis", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Redis Ping Response: %s", pong)
	})

	http.ListenAndServe(":8080", nil)
}
