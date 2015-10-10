package main

import (
	"flag"
	"io"
	"strings"
	"bufio"
	"os"
	"encoding/json"
	"net/http"
	"strconv"
	"errors"
	"log"
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
	for i, __ := range parts {
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

func ApiJsonRoundTrip(in io.Reader, out io.Writer, url, username, password, countReq, countGot, countTotal, countScale string) (err error) {
	var request, response interface{}
	var requestString, responseString []byte
	var current, total int
	decoder := json.NewDecoder(in)
	requestString, err := http.NewRequest("POST", url, requestBody)
	if username != "" && password != "" {
		requestString.SetBasicAuth(username, password)
	}

	for decoder.More() {
		err = decoder.Decode(&request)
		if err != nil {
			return err
		}

		current = request["params"]["page_num"].(int)

		for total == 0 || current < total {

			request["params"]["page_num"] = current
			requestString, err = json.Marshal(request)
			if err != nil {
				log.Fatalf("failed to build request body - %v\n%s", err, request)
			}

			client.Body = writeBuf.(io.ReadCloser)
			responseString, err := client.Do(requestString)
			if err != nil {
				log.Fatalf("failed to run request - %v", err)
			}

			err = json.Decode(responseString, &response)
			if err != nil {
				log.Fatalf("failed to decode the response body - %v\n%q", err, responseString)
			}
			current++
			total = reasponse["page_total"]
			out.Write(responseString)
		}
	}
	if err == io.EOF {
		return nil
	} else {
		return err
	}
}

var username = flag.String("username", "", "username to use for authentication")
var password = flag.String("username", "", "username to use for authentication")

func main() {
	url := flag.String("url", "", "url location to direct POSt")
	username := flag.String("username", "", "username to use for authentication")
	password := flag.String("username", "", "username to use for authentication")

	countReq := 
	flag.Parse()

	options := map[string]interface{}{
		"username": username,
		"password": password,
		"url": url,

	}

	if err := PrettyPrint(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), options); err != nil {
		log.Fatal(err)
	}
}
