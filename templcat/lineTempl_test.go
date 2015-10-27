package main

import (
	"encoding/json"
	"fmt"
)

func ExampleTempalteBuild() {
	templateBytes := []byte(`{"Everything":"Awesome","Team":{"Everything":"Cool"}}`)
	var testData interface{}
	err := json.Unmarshal(testBytes, &testData)
	if err != nil {
		fmt.Printf("hit a snag unmarshalling the data - %v", err)
	}

	item, err := GetItem(testData, []interface{}{"Team", "Everything"})
	if err != nil {
		fmt.Printf("hit a snag retrieving the item - %v", err)
	}
	fmt.Println(item)

	// Output:
	// Cool
}
