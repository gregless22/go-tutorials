package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	pl = fmt.Println
)

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	//cathc the error
	if err != nil {
		pl(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	pl(responseObject.Name)
	pl(len(responseObject.Pokemon))

	// list all of the pokemon
	for i, pokemon := range responseObject.Pokemon {
		pl(i, pokemon.Species.Name)

	}
}
