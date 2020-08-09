package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Pokemmon struct{
	Name string `json:"name"`
	Pokemmon_entries []Pokemmon_entries `json:"pokemon_entries"`
}

type Pokemmon_entries struct{
	Entry_Number int `json:"entry_number"`
	Pokemmon_species Pokemmon_species`json:"pokemon_species"`
}


type Pokemmon_species struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

func main(){
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1);
	}	

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil{
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))

	var responObj Pokemmon
	json.Unmarshal(responseData,&responObj)

	for i := 0; i < len(responObj.Pokemmon_entries);i++{
			fmt.Println(responObj.Pokemmon_entries[i].Pokemmon_species.Name)
	}
}