package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/JackKnifed/cliTricks"
)

func jsonDecoder(in io.Reader, out bufio.Writer, t [][]interface{}, sep string) (err error) {
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
		out.WriteString(line)
	}

	if err == io.EOF {
		return nil
	} else {
		return err
	}
	out.Flush()
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

}
