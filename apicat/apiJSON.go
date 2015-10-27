package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JackKnifed/cliTricks"
	"golang.org/x/net/publicsuffix"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"log"
	"os"
)

type config struct {
	username string
	password string
	url      string
	locReq   []string
	locCur   []string
	locTotal []string
	locInc   int
}

func loopRequest(requestData interface{}, out io.Writer, opts config) (err error) {

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}
	client := http.Client{Jar: jar}

	requestBytes, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", opts.url, bytes.NewReader(requestBytes))
	if err != nil {
		return err
	}

	if opts.username != "" && opts.password != "" {
		request.SetBasicAuth(opts.username, opts.password)
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

	_, err = out.Write(responseBytes)
	if err != nil {
		return err
	}

	if len(opts.locReq) < 1 && len(opts.locCur) < 1 && len(opts.locTotal) < 1 {
		return nil
	}

	var reqPage, curPage, totalPage int

	if len(opts.locReq) > 0 {
		reqPage, err = cliTricks.GetInt(requestData, opts.locReq)
		if err != nil {
			return fmt.Errorf("bad request page - %v", err)
		}
	} else {
		reqPage = 1
	}

	if len(opts.locCur) > 0 {
		curPage, err = cliTricks.GetInt(requestData, opts.locCur)
		if err != nil {
			return fmt.Errorf("bad current page - %v", err)
		}
	} else {
		curPage = 1
	}

	if len(opts.locTotal) > 0 {
		totalPage, err = cliTricks.GetInt(requestData, opts.locTotal)
		if err != nil {
			return fmt.Errorf("bad total page - %v", err)
		}
	} else {
		totalPage = 1
	}

	for reqPage < totalPage {
		curPage += opts.locInc
		err = cliTricks.SetItem(requestData, curPage, opts.locCur)
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

		curPage, err = cliTricks.GetInt(requestData, opts.locCur)
		if err != nil {
			return fmt.Errorf("bad current page - %v", err)
		}
	}
	return nil
}

func ApiJsonRoundTrip(in io.Reader, out io.Writer, opt config) (err error) {
	var requestData interface{}

	decoder := json.NewDecoder(in)

	for decoder.More() {
		err = decoder.Decode(&requestData)
		if err != nil {
			return err
		}
		err = loopRequest(requestData, out, opt)
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
	username := flag.String("username", "", "username to use for authentication")
	password := flag.String("username", "", "username to use for authentication")
	url := flag.String("url", "", "url location to direct POSt")
	locInc := flag.Int("pageIncrement", 1, "number to increase location request by")

	flag.Parse()

	opts := config{
		username: *username,
		password: *password,
		url:      *url,
		locInc:   *locInc,
		locReq:   cliTricks.BreakupStringArray(*locReqString),
		locCur:   cliTricks.BreakupStringArray(*locCurString),
		locTotal: cliTricks.BreakupStringArray(*locTotalString),
	}

	err := ApiJsonRoundTrip(os.Stdin, os.Stdout, opts)
	if err != nil {
		log.Fatal(err)
	}
}
