package unitcovertion

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

var Clients []UnitConverts;

type UnitConverts struct {
	UnitType   string `csv:"conversion_type"` 
	From      string `csv:"from"` 
	To         string `csv:"to"` 
	Multiplier float32  `csv:"multiply_by"` 	
}

var UntiConvert= make(map[string]UnitConverts)



func LoadUnitConversions(filePath string) {

	in, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer in.Close()

    // clients := []*UnitConverts{}

    if err := gocsv.UnmarshalFile(in, &Clients); err != nil {
        panic(err)
    }
    for _, client := range Clients {
        fmt.Println("Hello, ", client.UnitType,client.To)
		UntiConvert[client.To]=client
		// untiConvert[client.From+"TO"+client.To]=client.Multiplier
    }

}


func GetConverted(from string ,to string )string{



if(UntiConvert[from].UnitType!=UntiConvert[to].UnitType){

	return "Cant be converterd"
}else{
	response:= ((1/UntiConvert[from].Multiplier)*1/UntiConvert[to].Multiplier)
	return fmt.Sprintf("%v",response)

}

}

