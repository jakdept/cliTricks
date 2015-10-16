package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	ErrCantFind     = fmt.Errorf("can't find")
	ErrInvalidIndex = fmt.Errorf("invalid index")
)

type typeError string

func (e typeError) Error() string { return fmt.Sprintf("wrong data type. a %s was expected", e) }

func setItem(data interface{}, t []interface{}, val interface{}) (err error) {
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
				return ErrCantFind
			}
			return setItem(nextData, nextT, val)
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
				return ErrInvalidIndex
			}
			return setItem(d[tt], nextT, val)
		}
		d[tt] = val
		return
	default:
		panic("wut")
	}
	return
}

func main() {
	testData := []byte(`
{
	"field1": {
		"field1.1": "novalue",
		"field1.2": {
			"field1.2.1": 1,
			"field1.2.2": "something"
		}
	},
	"field2": [[1,2],[3,4,5]]
}
`)
	res := make(map[string]interface{})
	if err := json.Unmarshal(testData, &res); err != nil {
		fmt.Fprintln(os.Stderr, "can't unmarshal json:", err)
		os.Exit(1)
	}
	fmt.Println(res)
	if err := setItem(res, []interface{}{"field1", "field1.1"}, "hello world"); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	if err := setItem(res, []interface{}{"field1", "field1.2", "field1.2.1"}, 2); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	if err := setItem(res, []interface{}{"field1", "field1.2", "field1.2.2"}, 3.14); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	if err := setItem(res, []interface{}{"field2", 0, 1}, 0xbeef); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	if err := setItem(res, []interface{}{"field2", 1}, 12345); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	fmt.Println(res)
}
