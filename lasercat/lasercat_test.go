package main

import(
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	)

func ExampleGetItem() {
	testBytes := []byte(`{"Everything":"Awesome","Team":{"Everything":"Cool"}}`)
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