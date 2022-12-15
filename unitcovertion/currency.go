package unitcovertion

import "fmt"

var currency map[string]float32

func GetCurrencyUnit(unit string) float32 {
	fmt.Println(currency["USD"])
	return currency[unit]
}

func SetCurrencyUnit(update map[string]float32) {
	currency = update
}