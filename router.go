package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/microservice1", microservice1Handler)
    myRouter.HandleFunc("/microservice2", microservice2Handler)

    log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), myRouter))
}

func microservice1Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Microservice 1 is called\n"))
}

func microservice2Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Microservice 2 is called\n"))
}

func main() {
    loadEnv()
    handleRequests()
}