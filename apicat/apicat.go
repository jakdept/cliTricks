package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/jakdept/cliTricks"
	"golang.org/x/net/publicsuffix"
)

type config struct {
	username string
	password string
	url      string
	delay    time.Duration
	locReq   []interface{}
	locCur   []interface{}
	locTotal []interface{}
	locInc   int
}

func runRequest(c http.Client, b []byte, out io.Writer, opts config) (bool, error) {
	var respData interface{}

	req, err := http.NewRequest("POST", opts.url, bytes.NewReader(b))
	if err != nil {
		return false, fmt.Errorf("could not build request - %v", err)
	}

	if opts.username != "" && opts.password != "" {
		req.SetBasicAuth(opts.username, opts.password)
	}

	resp, err := c.Do(req)
	if err != nil {
		return false, fmt.Errorf("cannot send to [%s] - %v", req.URL.String(), err)
	}

	// some cheap handling for the request
	if resp.StatusCode != 200 {
		return false, fmt.Errorf("got a %d response from the api server - %s", resp.StatusCode, resp.Status)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("cannot read response - %v", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return false, fmt.Errorf("failed to close body - %v", err)
	}

	_, err = out.Write(respBytes)
	if err != nil {
		return false, fmt.Errorf("cannot output response - %v", err)
	}

	// at this point we're just looking to see if the page was the last page
	if len(opts.locCur) <= 0 || len(opts.locTotal) <= 0 {
		return true, nil
	}

	err = json.Unmarshal(respBytes, &respData)
	if err != nil {
		return false, fmt.Errorf("response not json - %v", err)
	}

	curPage, err := cliTricks.GetInt(respData, opts.locCur)
	if err != nil {
		return false, fmt.Errorf("bad current page - %v", err)
	}

	totalPage, err := cliTricks.GetInt(respData, opts.locTotal)
	if err != nil {
		return false, fmt.Errorf("bad total page - %v", err)
	}
	return curPage >= totalPage, nil
}

func loopRequest(reqData interface{}, out io.Writer, opts config) error {
	var done bool
	// create client to be used
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}
	c := http.Client{Jar: jar}

	for !done {
		// build then run the request
		reqBytes, err := json.Marshal(reqData)
		if err != nil {
			return fmt.Errorf("could not convert interface to bytes - %v", err)
		}
		done, err = runRequest(c, reqBytes, out, opts)
		if err != nil {
			return fmt.Errorf("got bad response from requests - %v", err)
		}

		// finally increment
		if len(opts.locReq) <= 0 {
			return nil
		}
		reqPage, err := cliTricks.GetInt(reqData, opts.locReq)
		if err != nil {
			return fmt.Errorf("failed to get the current page number before increment - %v", err)
		}
		reqPage += opts.locInc
		cliTricks.SetItem(reqData, opts.locReq, reqPage)
		if err != nil {
			return fmt.Errorf("failed to set the current page number after increment - %v", err)
		}
		time.Sleep(opts.delay)
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
	password := flag.String("password", "", "username to use for authentication")
	url := flag.String("url", "", "url location to direct POSt")
	locInc := flag.Int("pageIncrement", 1, "number to increase location request by")
	delay := flag.Int("delay", 100, "milliseconds to wait between subsequent requests")

	flag.Parse()

	opts := config{
		username: *username,
		password: *password,
		url:      *url,
		locInc:   *locInc,
		locReq:   cliTricks.BreakupArray(*locReqString),
		locCur:   cliTricks.BreakupArray(*locCurString),
		locTotal: cliTricks.BreakupArray(*locTotalString),
		delay:    time.Duration(*delay) * time.Millisecond,
	}

	err := ApiJsonRoundTrip(os.Stdin, os.Stdout, opts)
	if err != nil {
		log.Fatal(err)
	}
}
