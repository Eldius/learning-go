package tools

import (
	"encoding/json"
	"fmt"
)

/*
Debug just print object asin JSON format
*/
func Debug(obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
