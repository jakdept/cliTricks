package cliTricks

import (
"fmt"
"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	)

func TestBreakupStringArray(t *testing.T) {
	testData := []struct{
		input string
		output []string
	}{
		{
			input: "apple,banana,cherry",
			output: []string{"apple","banana", "cherry"},
		},{
			input: "Dog, Eagle, Fox",
			output: []string{"Dog", "Eagle", "Fox"},
		},{
			input: "[Green Beans][Hot Tamales][Ice Cream]",
			output: []string{"Green Beans", "Hot Tamales", "Ice Cream"},
		},{
			input: "[JellyBean],[KitKat],[Marshmallow]",
			output: []string{"JellyBean", "KitKat", "Marshmallow"},
		},{
			input: "[\"Nutella\"],[\"Oatmeal\"],[\"Pie\"]",
			output: []string{"Nutella", "Oatmeal", "Pie"},
		},
	}

	for _, oneTest := range testData {
			assert.Equal(t, oneTest.output, BreakupStringArray(oneTest.input), "BreakupStringArray returned non-expected results")
	}
}

// OMG GetItem works. Kinda. A bit. w/e.
func ExampleGetItem(){
	testBytes := []byte(`{"Everything":"Awesome","Team":{"Everything":"Cool"}}`)
	var testData interface{}
	err := json.Unmarshal(testBytes, &testData)
	if err != nil {
		fmt.Printf("hit a snag unmarshalling the data - %v", err)
	}

	item, err := GetItem(testData, []string{"Team", "Everything"})
	if err != nil {
		fmt.Printf("hit a snag retrieving the item - %v", err)
	}
	fmt.Println(item)

	// Output:
	// Cool
}

func ExampleGetInt(){
	testBytes := []byte(`{"Everything":"Awesome","Team":{"Everything":"Cool", "Solution": 63}}`)
	var testData interface{}
	err := json.Unmarshal(testBytes, &testData)
	if err != nil {
		fmt.Printf("hit a snag unmarshalling the data - %v", err)
	}

	item, err := GetInt(testData, []string{"Team", "Solution"})
	if err != nil {
		fmt.Printf("hit a snag retrieving the item - %v", err)
		return
	}
	fmt.Println(item)
	fmt.Println(reflect.TypeOf(item))

	// Output:
	// 63
	// int
}

/*
func TestGetInt(t *testing.T) {
	testData := []struct{
		input interface{}
		target []string
		output int
	}{
		{
			input: interface{}{
					map[string]interface{}{
						"params": map[string]int{
							"data":63,
					},
				},
			},
			target: []string{"params", "data",},
			output: 63,
		},
	}

	for _, oneTest := range testData {
		result, err := GetInt(oneTest.input, oneTest.target)
		assert.Equal(t, oneTest.output, result)
		assert.NoError(t, err)
	}
}
*/