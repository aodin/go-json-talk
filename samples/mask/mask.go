package main

import (
	"encoding/json"
	"fmt"
)

// BEGIN OMIT
type Detail struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
	Notes string `json:"notes"`
}

type List struct {
	Detail
	Notes string `json:"notes,omitempty"` // the json field MUST match
}

func main() {
	detail := Detail{Name: "Bill Murray", Value: 1, Notes: "Some notes"}
	b, _ := json.MarshalIndent(detail, "", "  ") // ignored error
	fmt.Printf("%s\n", b)
	b, _ = json.MarshalIndent(List{Detail: detail}, "", "  ") // ignored error
	fmt.Printf("%s", b)
}
