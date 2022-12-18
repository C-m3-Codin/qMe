package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	unitcovertion "github.com/C-m3-Codin/q_me/unitcovertion"
)


type response struct   {
	Rates map[string]float32
	Success bool
	Historical bool
	Base string
	Date string
}

// schedules to call the api every x amount of time
func ScheduleCurrency(){
	for true {
		
		unitcovertion.SetCurrencyUnit(getCurrencyAPI())
        time.Sleep(10 * time.Minute)

    }
}

// hits the api to get exchange rates
func getCurrencyAPI()map[string]float32 {
	client := http.Client{}
	currentTime := time.Now()
	apiEndpoint:="https://api.exchangerate.host/latestv?v="+currentTime.Format("2006-01-02")
	fmt.Println("\n\n\n")
	request, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	var result response
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Rates

}



        