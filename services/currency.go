package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// schedules to call the api every x amount of time
func ScheduleCurrency(){
	for true {
        fmt.Println("Hello !!")
		getCurrency()
        time.Sleep(2 * time.Second)
    }
}

// hits the api to get exchange rates
func getCurrency() {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://api.exchangerate.host/latest", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Println(result["rates"])
}



        