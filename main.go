package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"

	"github.com/gorilla/mux"
)

func addNewPokemon(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pokemon database.Pokemon

	json.Unmarshal(requestBody, &pokemon)
	found := false
	for i := 0; i < len(database.PokemonDb); i++ {
		if database.PokemonDb[i] == pokemon {
			found = true
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	if !found {
		database.PokemonDb = append(database.PokemonDb, pokemon)
	}
	w.WriteHeader(http.StatusOK)
}

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}
func isIn(id string) bool {

	for i := 0; i < len(database.PokemonDb); i++ {
		if database.PokemonDb[i].ID == id {
			return true
		}

	}
	return false
}

func handleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "0"
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons", addNewPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	//
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
