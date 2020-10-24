package main

import (
	"bytes"
	"math/rand"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPok(t *testing.T) {
	newId := strconv.Itoa(rand.Intn(100))
	var jsonStr = []byte(`  {
        "ID":` + newId + `
        "Name": "Csssharmeleon",
        "Type": "Fire"
	}`)

	http.NewRequest("POST", "https://pokemontengoqueatraparlos.herokuapp.com/pokemons", bytes.NewBuffer(jsonStr))

	assert.Equal(t, true, isIn(newId), "verificando que se haya agregado")

}
