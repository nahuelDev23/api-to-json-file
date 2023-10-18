package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// type Response struct {
// 	Count    int       `json:"count"`
// 	Next     string    `json:"next"`
// 	Previous string    `json:"previous"`
// 	Results  []Pokemon `json:"results"`
// }

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	// var responseObject Response
	var responseObject struct {
		Pokemon []Pokemon `json:"results"`
	}
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Println(responseObject)

	file, err := json.MarshalIndent(responseObject.Pokemon, "", " ")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	err = os.WriteFile("test.json", file, 0644)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
