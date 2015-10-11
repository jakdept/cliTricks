package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JackKnifed/cliTricks"
	"golang.org/x/net/publicsuffix"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

var options struct {
	username string
	password string
	url      string
	locReq   []string
	locCur   []string
	locTotal []string
	locInc   int
}

func loopRequest(requestData interface{}, out io.Writer, username, password, url string, locReq, locCur, locTotal []string, locInc int) error {

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}

	requestBytes, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBytes))
	if err != nil {
		return err
	}

	if username != "" && password != "" {
		request.SetBasicAuth(username, password)
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	var responseBytes []byte
	var responseData interface{}

	_, err = response.Body.Read(responseBytes)
	if err != nil {
		return err
	}

	_, err = out.Write(responseBytes)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBytes, responseData)
	if err != nil {
		return err
	}

	var reqPage, curPage, totalPage int

	reqPage, err = cliTricks.GetInt(requestData, locReq)
	if err != nil {
		return fmt.Errorf("bad request page - %v", err)
	}

	curPage, err = cliTricks.GetInt(requestData, locCur)
	if err != nil {
		return fmt.Errorf("bad current page - %v", err)
	}

	totalPage, err = cliTricks.GetInt(requestData, locTotal)
	if err != nil {
		return fmt.Errorf("bad current page - %v", err)
	}

	for curPage < totalPage {
		curPage += locInc
		err = cliTricks.SetItem(requestData, curPage, locCur)
		if err != nil {
			fmt.Errorf("failed to set the current page - %v", err)
		}

		requestBytes, err = json.Marshal(requestData)
		if err != nil {
			return err
		}

		request.Body = ioutil.NopCloser(bytes.NewReader(requestBytes))
		response, err = client.Do(request)
		if err != nil {
			return err
		}

		_, err = response.Body.Read(responseBytes)
		if err != nil {
			return err
		}

		err = json.Unmarshal(responseBytes, responseData)
		if err != nil {
			return err
		}

		curPage, err = cliTricks.GetInt(requestData, locCur)
		if err != nil {
			return fmt.Errorf("bad current page - %v", err)
		}
	}
	return nil
}

func ApiJsonRoundTrip(in io.Reader, out io.Writer, url, username, password string, locReq, locCur, locTotal []string, locInc int) (err error) {
	var requestData interface{}

	decoder := json.NewDecoder(in)

	for decoder.More() {
		err = decoder.Decode(&requestData)
		if err != nil {
			return err
		}
		err = loopRequest(requestData, out, username, password, url, locReq, locCur, locTotal, locInc)
		if err != nil {
			return err
		}
	}

	if err == io.EOF {
		return nil
	} else {
		return err
	}
}

func main() {
	locReqString := flag.String("requestedPage", "", "location in the request of the page")
	locCurString := flag.String("currentPage", "", "location in the response of the page returned")
	locTotalString := flag.String("totalPage", "", "location in the response of the total pages")

	options := options{
		username: flag.String("username", "", "username to use for authentication"),
		password: flag.String("username", "", "username to use for authentication"),
		url: flag.String("url", "", "url location to direct POSt"),
		locInc: flag.Int("pageIncrement", 1, "number to increase location request by"),
	}

	flag.Parse()

	options.locReq = cliTricks.BreakupStringArray(locReqString)
	options.locCur = cliTricks.BreakupStringArray(locCurString)
	options.locTotal = cliTricks.BreakupStringArray(locTotalString)
	
}