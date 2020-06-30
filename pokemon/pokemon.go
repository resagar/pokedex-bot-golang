package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pokemon struct {
	id int `json:"id"`
	Name string `json:"name"`
	Order int `json:"order"`
	Abilities []Abilities `json:"abilities"`
	GameIndex []GameIndex `json:"game_indices"`
	Species struct{Name string `json:"name"`} `json:"species"`
	Moves []Moves `json:"moves"`
	Types []Types `json:"types"`
	Sprite Sprite `json:"sprites"`
}

type Abilities struct {
	Ability struct{Name string `json:"name"`} `json:"ability"`
}

type GameIndex struct {
	GameIndex int `json:"game_index"`
	Version struct{Name string `json:"name"`} `json:"version"`
}

type Moves struct {
	Move struct{Name string `json:"name"`} `json:"move"`
}

type Types struct {
	Type struct{Name string `json:"name"`} `json:"type"`
}

type Sprite struct {
	FrontDefault string `json:"front_default"`
}

type PokemonList struct {
	Results []struct{Name string `json:"name"`} `json:"results"`
	Count int `json:"count"`
}

func PokemonAll() PokemonList{
	request, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100&offset=200")
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	dataList, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var result PokemonList
	json.Unmarshal(dataList, &result)
	return result
}

func PokemonOne(pokeName string)  Pokemon{
	var name *string = &pokeName

	if pokeName == "ho"{
		*name = "ho-oh"
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokeName)
	request,err:= http.Get(url)
	if err != nil{
		panic(err)
	}
	defer request.Body.Close()
	body,err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	var pokemon Pokemon
	json.Unmarshal(body, &pokemon)
	return pokemon
}