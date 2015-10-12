package cliTricks

import (
"testing"
	"github.com/stretchr/testify/assert"
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

	for _, individualTest := range testData {
			assert.Equal(t, individualTest.output, BreakupStringArray(individualTest.input), "BreakupStringArray returned non-expected results")
	}
}
