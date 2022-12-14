package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCurrency() {
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



        