package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT
type options struct {
	Alias           int64 `json:"alias"`
	Ignored         int64 `json:"-"`
	Omitted         int64 `json:",omitempty"`
	AsString        int64 `json:",string,omitempty"`
	AsStringOmitted int64 `json:",string,omitempty"`
}

func main() {
	data := options{
		Alias:           1,
		Ignored:         1,
		Omitted:         0,
		AsString:        1,
		AsStringOmitted: 0,
	}
	b, _ := json.MarshalIndent(data, "", "  ") // ignored error
	fmt.Printf("%s", b)
}

// END OMIT
