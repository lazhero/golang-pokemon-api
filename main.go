package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pokemon-api/database"
	"strings"

	"github.com/gorilla/mux"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons", addNewPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func addNewPokemon(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body) //lee la vara
	var newPokemon database.Pokemon
	json.Unmarshal(reqBody, &newPokemon) //guardo la informacion del body en el newpokemon
	canIaddThis := true
	for i := 0; i < len(database.PokemonDb); i++ {
		fmt.Println(database.PokemonDb[i].ID)
		if strings.Compare(database.PokemonDb[i].ID, newPokemon.ID) == 1 {
			canIaddThis = false
		}

	}
	if canIaddThis {
		database.PokemonDb = append(database.PokemonDb, newPokemon) //localmente, a mi array le agrego este compa
		json.NewEncoder(w).Encode(newPokemon)                       //agrego el compa al array
	}

}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
