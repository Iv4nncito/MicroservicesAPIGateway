package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Define a simple cache structure. In production, consider using a more robust caching solution.
type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, found := c.items[key]
	return val, found
}

func (c *Cache) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func newCache() *Cache {
	return &Cache{
		items: make(map[string]string),
	}
}

var (
	myCache = newCache() // Initialize the cache
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	setupRoutes(router)

	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", handleRoot).Methods("GET")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	const cacheKey = "response"
	if val, found := myCache.Get(cacheKey); found {
		fmt.Println("Cache Hit!")
		fmt.Fprintln(w, val)
		return
	}
	response := "API Gateway is up and running!"

	// Simulating a delay in response generation, e.g., from calling downstream services
	time.Sleep(2 * time.Second) 

	myCache.Set(cacheKey, response)
	fmt.Println("Cache Miss - Response Generated")
	fmt.Fprintln(w, response)
}