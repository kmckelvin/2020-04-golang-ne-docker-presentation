package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := client.Incr("call_count")
		fmt.Fprintf(w, "Result: %s", result.String())
	})

	port := os.Getenv("PORT")
	log.Println("Listening on: ", port)
	http.ListenAndServe(":"+port, handler)
}
