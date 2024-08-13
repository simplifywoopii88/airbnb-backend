package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(s interface{}) {
	jsonObj, _ := json.Marshal(s)
	fmt.Println(string(jsonObj))
}
