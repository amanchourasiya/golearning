package features

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SBC struct {
	Name    string
	Config  string
	IsCloud bool
	Memory  int
}

func JsonDemo() {
	// Converting struct to json
	swe := SBC{}
	swe.Name = "Swe"
	swe.Config = "default"
	swe.Memory = 10
	swe.IsCloud = true

	sweJson, _ := json.Marshal(swe)

	fmt.Printf("Struct to json %s\n", string(sweJson))

	// converting json to struct
	sweLite := SBC{}
	json.Unmarshal(sweJson, &sweLite)

	fmt.Printf("SweLIte config %s\n", sweLite.Config)
}

// RequestsDemo just demonstrates a simple get request
func RequestsDemo() {
	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		log.Fatalf("Error while make rest call %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error decoding response body %v\n", err)
	}

	log.Printf("Response: %s\n", body)

}
