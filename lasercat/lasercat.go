package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"flag"
	"os"

	"github.com/JackKnifed/cliTricks"
)

// Define a type named "intslice" as a slice of ints
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

func jsonDecoder(in io.Reader, out io.Writer, t [][]interface{}, sep string) (err error) {
	var requestData interface{}
	var line string

	decoder := json.NewDecoder(in)

	for decoder.More() {
		err = decoder.Decode(&requestData)
		if err != nil {
			return err
		}
		line, err = cherryPick(requestData, t, sep)
		if err != nil {
			return err
		}
		out.Write([]byte(line))
	}

	if err == io.EOF {
		return nil
	} else {
		return err
	}
	return
}

func cherryPick(data interface{}, targets [][]interface{}, seperator string) (response string, err error) {
	var onePart interface{}
	var responseParts []string
	for _, t := range targets {
		onePart, err = cliTricks.GetItem(data, t)
		if err != nil {
			return "", err
		}
		responseParts = append(responseParts, fmt.Sprintf("%s", onePart))
	}
	response = strings.Join(responseParts, "\"")
	return
}

func main() {
	var targetStrings stringSlice
	flag.Var(&targetStrings, "target", "locations to pluck from input")

	flag.Parse()

	var targets [][]interface{}
	for _, oneTarget := range targetStrings {
		targets = append(targets, cliTricks.BreakupArray(oneTarget))
	}

	jsonDecoder(os.Stdin, os.Stdout, targets, "\t")
}
