package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var config Config

func main() {
	if len(os.Args) < 1 {
		panic("No config provided")
	}

	readConfig(os.Args[1])
	manager = UserManager{make(map[int]*User, 0)}
	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func readConfig(file string) {
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		panic(err)
	}

	parser := json.NewDecoder(configFile)
	parser.Decode(&config)
}
