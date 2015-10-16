package cliTricks

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errCantFindMap  = fmt.Errorf("can't find")
	errInvalidIndex = fmt.Errorf("invalid index")
	errNonMapInput  = fmt.Errorf("invalid data input (not map or string)")
)

type typeError string

func (e typeError) Error() string { return fmt.Sprintf("wrong data type. a %s was expected", e) }

func BreakupArray(input string) []interface{} {
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
	output := make([]interface{}, len(parts))
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

func GetInt(data interface{}, target []interface{}) (int, error) {
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

func GetItem(data interface{}, target []interface{}) (interface{}, error) {
	if targetInt, ok := target[0].(int); ok {
		if dataSafe, ok := data.([]interface{}); !ok {
			return nil, errors.New("got array address for non-array")
		} else if targetInt < 0 || targetInt > len(dataSafe)-1 {
			return nil, errors.New("non-existant array position")
		} else {
			if len(target) > 1 {
				return GetItem(dataSafe[targetInt], target[1:])
			} else {
				return dataSafe[targetInt], nil
			}
		}
	} else if targetString, ok := target[0].(string); ok {
		if dataSafe, ok := data.(map[string]interface{}); !ok {
			return nil, errors.New("got map address for non-map")
		} else if value, ok := dataSafe[targetString]; !ok {
			return nil, errors.New("non-existant map position")
		} else {
			if len(target) > 1 {
				return GetItem(value, target[1:])
			} else {
				return value, nil
			}
		}
	} else {
		return nil, fmt.Errorf("bad address - %s", target)
	}
}

func SetItem(data interface{}, t []interface{}, val interface{}) (err error) {
	if len(t) < 1 {
		panic("wut")
	}
	nextT := t[1:]
	switch d := data.(type) {
	case map[string]interface{}:
		tt, ok := t[0].(string)
		if !ok {
			return typeError("string")
		}
		if len(nextT) > 0 {
			nextData, ok := d[tt]
			if !ok {
				return errCantFindMap
			}
			return SetItem(nextData, nextT, val)
		}
		d[tt] = val
		return
	case []interface{}:
		tt, ok := t[0].(int)
		if !ok {
			return typeError("int")
		}
		if len(nextT) > 0 {
			if tt < 0 || tt >= len(d) {
				return errInvalidIndex
			}
			return SetItem(d[tt], nextT, val)
		}
		d[tt] = val
		return
	default:
		panic("wut")
	}
	return
}

// func SetItem(data interface{}, t []interface{}, val interface{}) (err error) {
// 	if len(t) < 1 {
// 		return errNonMapInput
// 	}
// 	switch d := data.(type) {
// 	case map[string]interface{}:
// 		tt, ok := t[0].(string)
// 		if !ok {
// 			return typeError("string")
// 		}
// 		if len(t[1:]) > 0 {
// 			nextData, ok := d[tt]
// 			if !ok {
// 				return errCantFindMap
// 			}
// 			return SetItem(nextData, t[1:], val)
// 		}
// 		d[tt] = val
// 		return
// 	case []interface{}:
// 		tt, ok := t[0].(int)
// 		if !ok {
// 			return typeError("int")
// 		}
// 		if len(t[1:]) > 0 {
// 			if tt < 0 || tt >= len(d) {
// 				return errInvalidIndex
// 			}
// 			return SetItem(d[tt], t[:1], val)
// 		}
// 		d[tt] = val
// 		return
// 	default:
// 		return errNonMapInput
// 	}
// 	return
// }