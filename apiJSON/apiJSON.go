package main

import (
	"flag"
	"io"
	"bufio"
	"os"
	"encoding/json"
	"net/http"
	"log"
)

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
