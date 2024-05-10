package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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

	request, err := http.NewRequest("GET","https://api.thecatapi.com/v1/images/search?format=json&limit=1", nil)	

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// TO-DO: Set Cat API TOKEN
	request.Header.Set("x-api-key", "TOKEN")

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)	
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
