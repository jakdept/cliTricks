package cliTricks

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func BreakupStringArray(input string) []string {
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
		parts[i] = strings.TrimSpace(parts[i])
		if strings.HasPrefix(parts[i], "\"") && strings.HasSuffix(parts[i], "\"") {
			output[i] = strings.Trim(parts[i], "\"")
		} else {
			number, err := strconv.Atoi(parts[i])
			if err == nil {
				output[i] = number
			} else {
				output[i] = parts[i]
			}
		}
	}
	return output
}

func GetItem(data interface{}, target []string) (interface{}, error) {
	if dataSafe, ok := data.([]interface{}); ok {
		targetInt, err := strconv.Atoi(target[0])
		if err != nil {
			return nil, errors.New("got non-int address for []interface{}")
		}
		if len(target) > 1 {
			// there's stuff on the inside to dive into
			return GetItem(dataSafe[targetInt], target[1:])
		} else {
			return dataSafe[targetInt], nil
		}
	} else if dataSafe, ok := data.(map[string]interface{}); ok {
		if len(target) > 1 {
			// there's stuff on the inside to dive into
			return GetItem(dataSafe[target[0]], target[1:])
		} else {
			return dataSafe[target[0]], nil
		}
	} else {
		return nil, fmt.Errorf("bad address - %s", target)
	}
}

func GetInt(data interface{}, target []string) (int, error) {
	var ok bool
	var value float64
	tempItem, err := GetItem(data, target)
	if err != nil {
		return -1, fmt.Errorf("bad item - %v", err)
	}
	if value, ok = tempItem.(float64); !ok {
		return -1, fmt.Errorf("got non-float item - %s", tempItem)
	}
	return int(value), nil
}

func SetItem(data, value interface{}, target []string) error {
	if dataSafe, ok := data.([]interface{}); ok {
		targetInt, err := strconv.Atoi(target[0])
		if err != nil {
			return err
		}
		if len(target) > 1 {
			// there's stuff on the inside to dive into
			return SetItem(dataSafe[targetInt], value, target[1:])
		} else {
			dataSafe[targetInt] = value
			return nil
		}
	} else if dataSafe, ok := data.(map[string]interface{}); ok {
		if len(target) > 1 {
			// there's stuff on the inside to dive into
			return SetItem(dataSafe[target[0]], value, target[1:])
		} else {
			dataSafe[target[0]] = value
			return nil
		}
	} else {
		return errors.New("bad address")
	}
}
