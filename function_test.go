package cliTricks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestBreakupStringArray(t *testing.T) {
	testData := []struct {
		input  string
		output []string
	}{
		{
			input:  "apple,banana,cherry",
			output: []string{"apple", "banana", "cherry"},
		}, {
			input:  "Dog, Eagle, Fox",
			output: []string{"Dog", "Eagle", "Fox"},
		}, {
			input:  "[Green Beans][Hot Tamales][Ice Cream]",
			output: []string{"Green Beans", "Hot Tamales", "Ice Cream"},
		}, {
			input:  "[JellyBean],[KitKat],[Marshmallow]",
			output: []string{"JellyBean", "KitKat", "Marshmallow"},
		}, {
			input:  "[\"Nutella\"],[\"Oatmeal\"],[\"Pie\"]",
			output: []string{"Nutella", "Oatmeal", "Pie"},
		},
	}

	for _, oneTest := range testData {
		assert.Equal(t, oneTest.output, BreakupStringArray(oneTest.input), "BreakupStringArray returned non-expected results")
	}
}

func ExampleGetItem() {
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

// func TestGetInt(t *testing.T) {
// 	testData := []struct{
// 		input interface{}
// 		target []string
// 		output int
// 	}{
// 		{
// 			input: map[string]interface{}{
// 						"params": map[string]float64{
// 							"data": 63,
// 					},
// 			},
// 			target: []string{"params", "data",},
// 			output: 63,
// 		},
// 	}

// 	for _, oneTest := range testData {
// 		result, err := GetInt(oneTest.input, oneTest.target)
// 		assert.Equal(t, oneTest.output, result)
// 		assert.NoError(t, err)

// 		result2, err := GetItem(oneTest.input, []string{"params"})
// 		fmt.Println(oneTest.input)
// 		fmt.Println(oneTest.output)
// 		fmt.Println(result2)
// 	}
// }

func TestGetIntJSON(t *testing.T) {
	testData := []struct {
		input  []byte
		target []string
		output int
		status error
	}{
		{
			input:  []byte(`{"params":{"data":63}}`),
			target: []string{"params", "data"},
			output: 63,
			status: nil,
		}, {
			// we always round down
			input:  []byte(`{"params":{"data":63.9}}`),
			target: []string{"params", "data"},
			output: 63,
			status: nil,
		}, {
			input:  []byte(`{"params":{"data":"potato"}}`),
			target: []string{"params", "data"},
			output: -1,
			status: errors.New("got non-float item - potato"),
		}, {
			input:  []byte(`{"params":{"data":"potato"}}`),
			target: []string{"bad", "address"},
			output: -1,
			status: errors.New("bad item - bad address - [address]"),
		},
	}

	for _, oneTest := range testData {
		var testData interface{}
		err := json.Unmarshal(oneTest.input, &testData)
		assert.Nil(t, err, "Problems unmarshaling the input")

		result, err := GetInt(testData, oneTest.target)
		assert.Equal(t, oneTest.output, result)
		assert.Equal(t, oneTest.status, err)
	}
}

func TestGetItemJSON(t *testing.T) {
	testData := []struct {
		input  []byte
		target []string
		output []byte
		status error
	}{
		{
			input:  []byte(`{"params":{"data":63}}`),
			target: []string{"params", "data"},
			output: []byte(`63`),
			status: nil,
		}, {
			// we always round down
			input:  []byte(`{"params":{"data":63.9}}`),
			target: []string{"params", "data"},
			output: []byte(`63.9`),
			status: nil,
		}, {
			input:  []byte(`{"params":{"data":"potato"}}`),
			target: []string{"params", "data"},
			output: []byte(`"potato"`),
			status: nil,
		}, {
			input:  []byte(`{"params":{"data":"potato"}}`),
			target: []string{"bad", "address"},
			output: []byte(""),
			status: errors.New("bad address - [address]"),
		},
	}

	for _, oneTest := range testData {
		var inputData, outputData interface{}
		err := json.Unmarshal(oneTest.input, &inputData)
		assert.Nil(t, err, "Problems unmarshaling the input")
		err = json.Unmarshal(oneTest.output, &outputData)

		result, err := GetItem(inputData, oneTest.target)
		assert.Equal(t, outputData, result)
		assert.Equal(t, oneTest.status, err)
	}
}

func ExampleGetInt() {
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
