package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT
type tagless struct {
	unexported int64
	ALL_CAPS   int64
	Upper      int64
	Lower      int64
}

var data = []byte(`{
  "unexported": 1,
  "all_caps": 1,
  "UPPER": 1,
  "lower": 1
}`)

func main() {
	var response tagless
	_ = json.Unmarshal(data, &response) // ignored error
	fmt.Printf("%+v", response)
}

// END OMIT
