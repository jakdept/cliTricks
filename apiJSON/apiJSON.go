package main

import (
	"flag"
	"io"
	"bufio"
	"io/ioutil"
	"fmt"
	"bytes"
	"os"
	"encoding/json"
	"golang.org/x/net/publicsuffix"
	"github.com/JackKnifed/cliTricks"
	"net/http"
	"log"
	"net/http/cookiejar"
)

func loopRequest(requestData interface{}, out io.Writer, username, password, url string, locReq, locCur, locTotal []string, locInc int) (error) {

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
	url := flag.String("url", "", "url location to direct POSt")
	username := flag.String("username", "", "username to use for authentication")
	password := flag.String("username", "", "username to use for authentication")

	flag.Parse()

	options := map[string]interface{}{
		"username": username,
		"password": password,
		"url": url,

	}
}
