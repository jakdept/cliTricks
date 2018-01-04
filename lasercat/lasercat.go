package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/JackKnifed/cliTricks"
	"github.com/davecgh/go-spew/spew"
)

// Define a type named "stringSlice" as a slice of strings
type stringSlice []string

// Now, for our new type, implement the two methods of
// the flag.Value interface...
// The first method is String() string
func (s *stringSlice) String() string {
	return fmt.Sprintf("%s", *s)
}

// The second method is Set(value string) error
func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func jsonDecoder(in io.Reader, out io.Writer, t [][]interface{}) (err error) {
	var requestData, item interface{}
	var line []string

	decoder := json.NewDecoder(in)

	for decoder.More() {
		err = decoder.Decode(&requestData)
		if err != nil {
			return err
		}
		line = []string{}
		for _, oneTarget := range t {
			item, err = cliTricks.GetItem(requestData, oneTarget)
			if err != nil {
				return fmt.Errorf("looking for [%s] in below - %v\n%s\n",
					oneTarget, err, spew.Sdump(requestData))
			}
			switch tItem := item.(type) {
			case int:
				line = append(line, strconv.Itoa(int(tItem)))
			case float64:
				if tItem == float64(int64(tItem)) {
					line = append(line, strconv.Itoa(int(tItem)))
				} else {
					line = append(line, fmt.Sprintf("%f", tItem))
				}
			case nil:
				line = append(line, "null")
			default:
				line = append(line, fmt.Sprintf("%s", item))
			}
		}
		out.Write([]byte(strings.Join(line, " ")))
		out.Write([]byte("\n"))
	}

	return err
}

func main() {
	var targetStrings stringSlice
	flag.Var(&targetStrings, "target", "locations to pluck from input")

	flag.Parse()

	var targets [][]interface{}
	for _, oneTarget := range targetStrings {
		targets = append(targets, cliTricks.BreakupArray(oneTarget))
	}

	if err := jsonDecoder(os.Stdin, os.Stdout, targets); err != nil {
		log.Print(err)
	}
}
