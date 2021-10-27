package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "github.com/julienschmidt/httprouter"
)

type beer struct {
	ID          string  `json:"id"`
	NAME        string  `json:"name"`
	TAGLINE     string  `json:"tagline"`
	DESCRIPTION string  `json:"description"`
	IMAGE_URL   string  `json:"image_url"`
	ABV         float32 `json:"abv"`
	IBU         int     `json:"ibu"`
	EBC         int     `json:"ebc"`
	SRM         int     `json:"srm"`
	PH          float32 `json:"ph"`
}

type BeersController struct {}

func get() {
	response, err := http.Get("https://api.punkapi.com/v2/beers")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject beer
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.NAME)
	fmt.Println(string(responseData))
}