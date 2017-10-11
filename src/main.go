package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var repo UserConfigRepository

func main() {
	if len(os.Args) < 1 {
		panic("No config provided")
	}

	repo = UserConfigRepository{make(map[int]FName, 0), make(map[int]string, 0), make(map[int]string, 0), make(map[int]string, 0)}
	readConfig(os.Args[1])
	manager = UserManager{make(map[int]*User, 0)}
	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func readConfig(file string) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		panic(err)
	}

	parser := json.NewDecoder(configFile)
	parser.Decode(&config)

	for i, v := range config.FirstNames {
		repo.FirstNames[i] = v
	}

	for i, v := range config.LastNames {
		repo.LastNames[i] = v
	}

	for i, v := range config.Adjectives {
		repo.Adjectives[i] = v
	}

	for i, v := range config.Occupations {
		repo.Occupations[i] = v
	}
}
