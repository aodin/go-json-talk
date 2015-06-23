package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT
type Thing struct {
	Value int64
}

var data = []byte(`{
  "data": {
    "value": 1
  }
}`)

func main() {
	response := struct {
		Data Thing
	}{}
	_ = json.Unmarshal(data, &response) // ignored error
	fmt.Printf("%+v", response.Data)
}

// END OMIT
