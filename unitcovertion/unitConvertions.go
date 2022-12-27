package unitcovertion

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)


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
	var AllConvertions []UnitConverts;

    if err := gocsv.UnmarshalFile(in, &AllConvertions); err != nil {
        panic(err)
    }
    for _, eachUnit := range AllConvertions {
        fmt.Println("Hello, ", eachUnit.UnitType,eachUnit.To)
		UntiConvert[eachUnit.To]=eachUnit
		// untiConvert[client.From+"TO"+client.To]=client.Multiplier
    }

}


func GetConverted(from string ,to string, count_str string )string{

fmt.Println("to Mltiplier: ",UntiConvert[to])
fmt.Println("from Mltiplier: ",UntiConvert[from])

if(UntiConvert[from].UnitType!=UntiConvert[to].UnitType){
	return "Cant be converterd"
}else{
	response:= ((1/UntiConvert[from].Multiplier)/(1/UntiConvert[to].Multiplier))
	if count, err := strconv.ParseFloat(count_str, 32); err == nil {
		response=response*float32(count)

		return fmt.Sprintf("%v",response)
	}else{
		return "Cant be converterd"

	}
}
}

