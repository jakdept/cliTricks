package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//dont do this, see above edit
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	data, err = prettyprint(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
