package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
)

type Response []struct {
	Breeds 	[]any		`json:"breeds"`
	Id	   	string 		`json:"id"`
	Url	   	string 		`json:"url"`
	Width  	int			`json:"width"`
	Height  int			`json:"height"`
}

func main() {

	client := &http.Client{}

	err := godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
	}

	request, err := http.NewRequest("GET","https://api.thecatapi.com/v1/images/search?format=json&limit=1", nil)	

	if err != nil {
        log.Fatalf("Error creating the request %v", err)
	}

	CAT_API_TOKEN := os.Getenv("CAT_API_TOKEN")

	// TO-DO: Set Cat API TOKEN
	request.Header.Set("x-api-key", CAT_API_TOKEN)

	response, err := client.Do(request)
	if err != nil {
        log.Fatalf("Request failed with error %v", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var responseObj Response
	json.Unmarshal(responseData, &responseObj)
	
	url := responseObj[0].Url

	cmd := exec.Command("xdg-open", url)

	err = cmd.Start()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

}
