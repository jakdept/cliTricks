package cliTricks

import (
	"strings"
	"strconv"
	"errors"
)

func breakupStringArray(input string) []string {
	if strings.HasPrefix(input, "[") && strings.HasSuffix(input, "]") {
		input = strings.TrimPrefix(input, "[")
		input = strings.TrimSuffix(input, "]")
	}
	parts := strings.Split(input, "],[")
	if len(parts) < 2 {
		parts = strings.Split(input, "][")
	}
	if len(parts) < 2 {
		parts = strings.Split(input, ",")
	}
	for i, _ := range parts {
		if strings.HasPrefix(parts[i], "\"") && strings.HasSuffix(parts[i], "\"") {
			parts[i] = strings.Trim(parts[i], "\"")
		}
	}
	return parts
}

func getBuriedItem(data interface{}, target []string) (interface{}, error) {
	if dataSafe, ok := data.([]interface{}); ok {
		targetInt, err := strconv.Atoi(target[0])
		if err != nil {
			return nil, err
		}
		if len(target) > 1{
			// there's stuff on the inside to dive into
			return getBuriedItem(dataSafe[targetInt], target[1:])
		} else {
			return dataSafe[targetInt], nil
		}
	}	else if dataSafe, ok := data.(map[string]interface{}); ok {
		if len(target) > 1{
			// there's stuff on the inside to dive into
			return getBuriedItem(dataSafe[target[0]], target[1:])
		} else {
			return dataSafe[target[0]], nil
		}
	} else {
		return nil, errors.New("bad address")
	}
}

func setBuriedItem(data, value interface{}, target []string) error {
	if dataSafe, ok := data.([]interface{}); ok {
		targetInt, err := strconv.Atoi(target[0])
		if err != nil {
			return err
		}
		if len(target) > 1{
			// there's stuff on the inside to dive into
			return setBuriedItem(dataSafe[targetInt], value, target[1:])
		} else {
			dataSafe[targetInt] = value
			return nil
		}
	}	else if dataSafe, ok := data.(map[string]interface{}); ok {
		if len(target) > 1{
			// there's stuff on the inside to dive into
			return setBuriedItem(dataSafe[target[0]], value, target[1:])
		} else {
			dataSafe[target[0]] = value
			return nil
		}
	} else {
		return errors.New("bad address")
	}
}