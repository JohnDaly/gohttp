package gohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ajg/form"
)

//------------------------------------------------------------------------------
// JSON Data
//------------------------------------------------------------------------------

// JSONData serializes an interface{} of JSON and converts to an `io.Reader`.
func JSONData(values interface{}) (io.Reader, error) {
	if values == nil {
		return nil, nil
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonData)
	return buffer, nil
}

// ParseJSON parses the JSON data from an `io.Reader` into an `interface{}`.
// This method utilizes the `UserNumber` feature of the `json.Decoder` which
// causes the Decoder to unmarshal a number into an interface{} as a Number
// instead of as a float64.
func ParseJSON(data io.Reader) (interface{}, error) {
	decoder := json.NewDecoder(data)
	decoder.UseNumber()

	var json interface{}
	err := decoder.Decode(&json)
	if err != nil {
		return nil, err
	}

	return json, nil
}

//------------------------------------------------------------------------------
// Form Data
//------------------------------------------------------------------------------

// FormData ecodes from data from an interface{} and converts to an `io.Reader`.
func FormData(data interface{}) (io.Reader, error) {
	values, err := form.EncodeToValues(data)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBufferString(values.Encode())
	return body, nil
}

//------------------------------------------------------------------------------
// Pretty Print Data
//------------------------------------------------------------------------------

// PrettyPrint prints a JSON representation in a pretty printed manner.
func PrettyPrint(item interface{}) {
	b, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	strBuff := out.String()
	fmt.Printf("[%v] - %s\n", strings.ToUpper(os.Getenv("ENV")), strBuff)
}
