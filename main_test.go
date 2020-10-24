package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"pokemon-api/database"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isIn(id string) bool {

	for i := 0; i < len(database.PokemonDb); i++ {
		if database.PokemonDb[i].ID == id {
			return true
		}

	}
	return false
}

func TestConnection(t *testing.T) {

	resp, err := http.Get("https://pokemontengoqueatraparlos.herokuapp.com/pokemons")
	algo := resp == nil
	if err == nil {
	}
	assert.Equal(t, true, algo, "verificando conexion")
}

func TestAddPok(t *testing.T) {
	newID := strconv.Itoa(rand.Intn(100))
	var jsonStr = []byte(`  {
        "ID":` + newID + `
        "Name": "Csssharmeleon",
        "Type": "Fire"
	}`)

	http.NewRequest("POST", "https://pokemontengoqueatraparlos.herokuapp.com/pokemons", bytes.NewBuffer(jsonStr))

	assert.Equal(t, true, isIn(newID), "verificando que se haya agregado")

}
