package unitcovertion

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

// var Clients []UnitConverts;

type UnitConverts struct {
	UnitType   string `csv:"conversion_type"` 
	From      string `csv:"from"` 
	To         string `csv:"to"` 
	Multiplier float32  `csv:"multiply_by"` 	
}

var untiConvert map[string]float32



func LoadUnitConversions(filePath string) {

	in, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer in.Close()

    clients := []*UnitConverts{}

    if err := gocsv.UnmarshalFile(in, &clients); err != nil {
        panic(err)
    }
    for _, client := range clients {
        fmt.Println("Hello, ", client.UnitType)
		untiConvert[client.From+"TO"+client.To]=client.Multiplier
    }

}


func GetConverted(from string ,to string )string{

	v ,exist := untiConvert[from+"TO"+to]
if exist {
	return fmt.Sprintf("%v", v)
     
}
	// return untiConvert[from+"TO"+to]
	return "Cant be converterd"
}

