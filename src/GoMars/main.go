package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a, _ := json.Marshal(map[string]interface{}{
		"47": map[string]int{
			"helo": 182}})

	fmt.Println(string(a))
}
